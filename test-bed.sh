#!/bin/sh

USE_SINGLE=false

while getopts ":s" opt; do
  case $opt in
    s)
      USE_SINGLE=true
      ;;
    \?)
      echo "Invalid option: -$OPTARG" >&2
      exit 1
      ;;
  esac
done

# Kill previous test-bed
pkill -9 brig
rm -rf /tmp/{ali,bob}
rm -f ~/.config/brig/registry.yml

# Have some color in the logs even when piping it somewhere else.
export BRIG_COLOR=always
alias brig-ali='brig --port 6666'

if [ "$USE_SINGLE" = false ]; then
    alias brig-bob='brig --port 6667'
else
    # Fake the invocation of bob's brig.
    alias brig-bob='/bin/true'
fi

# Uncomment the following lines to circumvent logging to syslog.
(brig-ali --repo /tmp/ali daemon launch -s 2>&1 > /tmp/log.ali) &
(brig-bob --repo /tmp/bob daemon launch -s 2>&1 > /tmp/log.bob) &

# Give the daemon to start up a bit.
sleep 0.5

(brig-ali --repo /tmp/ali init ali -x > /dev/null) &
(brig-bob --repo /tmp/bob init bob -x > /dev/null) &

# Wait until the daemon fully booted up.
brig-ali daemon ping -w -c 0
brig-bob daemon ping -w -c 0

# Add them as remotes each
if [ "$USE_SINGLE" = false ]; then
    brig-ali remote add bob $(brig-bob whoami -f)
    brig-bob remote add ali $(brig-ali whoami -f)
fi

brig-ali stage BUGS ali-file
brig-ali commit -m 'Added ali-file'
brig-bob stage LICENSE bob-file
brig-bob commit -m 'Added bob-file'
