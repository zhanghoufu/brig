// Code generated by capnpc-go. DO NOT EDIT.

package capnp

import (
	strconv "strconv"
	capnp "zombiezen.com/go/capnproto2"
	text "zombiezen.com/go/capnproto2/encoding/text"
	schemas "zombiezen.com/go/capnproto2/schemas"
)

// Commit is a set of changes to nodes
type Commit struct{ capnp.Struct }
type Commit_merge Commit

// Commit_TypeID is the unique identifier for the type Commit.
const Commit_TypeID = 0x8da013c66e545daf

func NewCommit(s *capnp.Segment) (Commit, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 6})
	return Commit{st}, err
}

func NewRootCommit(s *capnp.Segment) (Commit, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 6})
	return Commit{st}, err
}

func ReadRootCommit(msg *capnp.Message) (Commit, error) {
	root, err := msg.RootPtr()
	return Commit{root.Struct()}, err
}

func (s Commit) String() string {
	str, _ := text.Marshal(0x8da013c66e545daf, s.Struct)
	return str
}

func (s Commit) Message() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s Commit) HasMessage() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Commit) MessageBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s Commit) SetMessage(v string) error {
	return s.Struct.SetText(0, v)
}

func (s Commit) Author() (string, error) {
	p, err := s.Struct.Ptr(1)
	return p.Text(), err
}

func (s Commit) HasAuthor() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s Commit) AuthorBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(1)
	return p.TextBytes(), err
}

func (s Commit) SetAuthor(v string) error {
	return s.Struct.SetText(1, v)
}

func (s Commit) Parent() ([]byte, error) {
	p, err := s.Struct.Ptr(2)
	return []byte(p.Data()), err
}

func (s Commit) HasParent() bool {
	p, err := s.Struct.Ptr(2)
	return p.IsValid() || err != nil
}

func (s Commit) SetParent(v []byte) error {
	return s.Struct.SetData(2, v)
}

func (s Commit) Root() ([]byte, error) {
	p, err := s.Struct.Ptr(3)
	return []byte(p.Data()), err
}

func (s Commit) HasRoot() bool {
	p, err := s.Struct.Ptr(3)
	return p.IsValid() || err != nil
}

func (s Commit) SetRoot(v []byte) error {
	return s.Struct.SetData(3, v)
}

func (s Commit) Index() int64 {
	return int64(s.Struct.Uint64(0))
}

func (s Commit) SetIndex(v int64) {
	s.Struct.SetUint64(0, uint64(v))
}

func (s Commit) Merge() Commit_merge { return Commit_merge(s) }

func (s Commit_merge) With() (string, error) {
	p, err := s.Struct.Ptr(4)
	return p.Text(), err
}

func (s Commit_merge) HasWith() bool {
	p, err := s.Struct.Ptr(4)
	return p.IsValid() || err != nil
}

func (s Commit_merge) WithBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(4)
	return p.TextBytes(), err
}

func (s Commit_merge) SetWith(v string) error {
	return s.Struct.SetText(4, v)
}

func (s Commit_merge) Head() ([]byte, error) {
	p, err := s.Struct.Ptr(5)
	return []byte(p.Data()), err
}

func (s Commit_merge) HasHead() bool {
	p, err := s.Struct.Ptr(5)
	return p.IsValid() || err != nil
}

func (s Commit_merge) SetHead(v []byte) error {
	return s.Struct.SetData(5, v)
}

// Commit_List is a list of Commit.
type Commit_List struct{ capnp.List }

// NewCommit creates a new list of Commit.
func NewCommit_List(s *capnp.Segment, sz int32) (Commit_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 6}, sz)
	return Commit_List{l}, err
}

func (s Commit_List) At(i int) Commit { return Commit{s.List.Struct(i)} }

func (s Commit_List) Set(i int, v Commit) error { return s.List.SetStruct(i, v.Struct) }

func (s Commit_List) String() string {
	str, _ := text.MarshalList(0x8da013c66e545daf, s.List)
	return str
}

// Commit_Promise is a wrapper for a Commit promised by a client call.
type Commit_Promise struct{ *capnp.Pipeline }

func (p Commit_Promise) Struct() (Commit, error) {
	s, err := p.Pipeline.Struct()
	return Commit{s}, err
}

func (p Commit_Promise) Merge() Commit_merge_Promise { return Commit_merge_Promise{p.Pipeline} }

// Commit_merge_Promise is a wrapper for a Commit_merge promised by a client call.
type Commit_merge_Promise struct{ *capnp.Pipeline }

func (p Commit_merge_Promise) Struct() (Commit_merge, error) {
	s, err := p.Pipeline.Struct()
	return Commit_merge{s}, err
}

// A single directory entry
type DirEntry struct{ capnp.Struct }

// DirEntry_TypeID is the unique identifier for the type DirEntry.
const DirEntry_TypeID = 0x8b15ee76774b1f9d

func NewDirEntry(s *capnp.Segment) (DirEntry, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	return DirEntry{st}, err
}

func NewRootDirEntry(s *capnp.Segment) (DirEntry, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	return DirEntry{st}, err
}

func ReadRootDirEntry(msg *capnp.Message) (DirEntry, error) {
	root, err := msg.RootPtr()
	return DirEntry{root.Struct()}, err
}

func (s DirEntry) String() string {
	str, _ := text.Marshal(0x8b15ee76774b1f9d, s.Struct)
	return str
}

func (s DirEntry) Name() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s DirEntry) HasName() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s DirEntry) NameBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s DirEntry) SetName(v string) error {
	return s.Struct.SetText(0, v)
}

func (s DirEntry) Hash() ([]byte, error) {
	p, err := s.Struct.Ptr(1)
	return []byte(p.Data()), err
}

func (s DirEntry) HasHash() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s DirEntry) SetHash(v []byte) error {
	return s.Struct.SetData(1, v)
}

// DirEntry_List is a list of DirEntry.
type DirEntry_List struct{ capnp.List }

// NewDirEntry creates a new list of DirEntry.
func NewDirEntry_List(s *capnp.Segment, sz int32) (DirEntry_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2}, sz)
	return DirEntry_List{l}, err
}

func (s DirEntry_List) At(i int) DirEntry { return DirEntry{s.List.Struct(i)} }

func (s DirEntry_List) Set(i int, v DirEntry) error { return s.List.SetStruct(i, v.Struct) }

func (s DirEntry_List) String() string {
	str, _ := text.MarshalList(0x8b15ee76774b1f9d, s.List)
	return str
}

// DirEntry_Promise is a wrapper for a DirEntry promised by a client call.
type DirEntry_Promise struct{ *capnp.Pipeline }

func (p DirEntry_Promise) Struct() (DirEntry, error) {
	s, err := p.Pipeline.Struct()
	return DirEntry{s}, err
}

// Directory contains one or more directories or files
type Directory struct{ capnp.Struct }

// Directory_TypeID is the unique identifier for the type Directory.
const Directory_TypeID = 0xe24c59306c829c01

func NewDirectory(s *capnp.Segment) (Directory, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 3})
	return Directory{st}, err
}

func NewRootDirectory(s *capnp.Segment) (Directory, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 3})
	return Directory{st}, err
}

func ReadRootDirectory(msg *capnp.Message) (Directory, error) {
	root, err := msg.RootPtr()
	return Directory{root.Struct()}, err
}

func (s Directory) String() string {
	str, _ := text.Marshal(0xe24c59306c829c01, s.Struct)
	return str
}

func (s Directory) Size() uint64 {
	return s.Struct.Uint64(0)
}

func (s Directory) SetSize(v uint64) {
	s.Struct.SetUint64(0, v)
}

func (s Directory) Parent() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s Directory) HasParent() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Directory) ParentBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s Directory) SetParent(v string) error {
	return s.Struct.SetText(0, v)
}

func (s Directory) Children() (DirEntry_List, error) {
	p, err := s.Struct.Ptr(1)
	return DirEntry_List{List: p.List()}, err
}

func (s Directory) HasChildren() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s Directory) SetChildren(v DirEntry_List) error {
	return s.Struct.SetPtr(1, v.List.ToPtr())
}

// NewChildren sets the children field to a newly
// allocated DirEntry_List, preferring placement in s's segment.
func (s Directory) NewChildren(n int32) (DirEntry_List, error) {
	l, err := NewDirEntry_List(s.Struct.Segment(), n)
	if err != nil {
		return DirEntry_List{}, err
	}
	err = s.Struct.SetPtr(1, l.List.ToPtr())
	return l, err
}

func (s Directory) Contents() (DirEntry_List, error) {
	p, err := s.Struct.Ptr(2)
	return DirEntry_List{List: p.List()}, err
}

func (s Directory) HasContents() bool {
	p, err := s.Struct.Ptr(2)
	return p.IsValid() || err != nil
}

func (s Directory) SetContents(v DirEntry_List) error {
	return s.Struct.SetPtr(2, v.List.ToPtr())
}

// NewContents sets the contents field to a newly
// allocated DirEntry_List, preferring placement in s's segment.
func (s Directory) NewContents(n int32) (DirEntry_List, error) {
	l, err := NewDirEntry_List(s.Struct.Segment(), n)
	if err != nil {
		return DirEntry_List{}, err
	}
	err = s.Struct.SetPtr(2, l.List.ToPtr())
	return l, err
}

// Directory_List is a list of Directory.
type Directory_List struct{ capnp.List }

// NewDirectory creates a new list of Directory.
func NewDirectory_List(s *capnp.Segment, sz int32) (Directory_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 3}, sz)
	return Directory_List{l}, err
}

func (s Directory_List) At(i int) Directory { return Directory{s.List.Struct(i)} }

func (s Directory_List) Set(i int, v Directory) error { return s.List.SetStruct(i, v.Struct) }

func (s Directory_List) String() string {
	str, _ := text.MarshalList(0xe24c59306c829c01, s.List)
	return str
}

// Directory_Promise is a wrapper for a Directory promised by a client call.
type Directory_Promise struct{ *capnp.Pipeline }

func (p Directory_Promise) Struct() (Directory, error) {
	s, err := p.Pipeline.Struct()
	return Directory{s}, err
}

// A leaf node in the MDAG
type File struct{ capnp.Struct }

// File_TypeID is the unique identifier for the type File.
const File_TypeID = 0x8ea7393d37893155

func NewFile(s *capnp.Segment) (File, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 2})
	return File{st}, err
}

func NewRootFile(s *capnp.Segment) (File, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 2})
	return File{st}, err
}

func ReadRootFile(msg *capnp.Message) (File, error) {
	root, err := msg.RootPtr()
	return File{root.Struct()}, err
}

func (s File) String() string {
	str, _ := text.Marshal(0x8ea7393d37893155, s.Struct)
	return str
}

func (s File) Size() uint64 {
	return s.Struct.Uint64(0)
}

func (s File) SetSize(v uint64) {
	s.Struct.SetUint64(0, v)
}

func (s File) Parent() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s File) HasParent() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s File) ParentBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s File) SetParent(v string) error {
	return s.Struct.SetText(0, v)
}

func (s File) Key() ([]byte, error) {
	p, err := s.Struct.Ptr(1)
	return []byte(p.Data()), err
}

func (s File) HasKey() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s File) SetKey(v []byte) error {
	return s.Struct.SetData(1, v)
}

// File_List is a list of File.
type File_List struct{ capnp.List }

// NewFile creates a new list of File.
func NewFile_List(s *capnp.Segment, sz int32) (File_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 2}, sz)
	return File_List{l}, err
}

func (s File_List) At(i int) File { return File{s.List.Struct(i)} }

func (s File_List) Set(i int, v File) error { return s.List.SetStruct(i, v.Struct) }

func (s File_List) String() string {
	str, _ := text.MarshalList(0x8ea7393d37893155, s.List)
	return str
}

// File_Promise is a wrapper for a File promised by a client call.
type File_Promise struct{ *capnp.Pipeline }

func (p File_Promise) Struct() (File, error) {
	s, err := p.Pipeline.Struct()
	return File{s}, err
}

// Ghost indicates that a certain node was at this path once
type Ghost struct{ capnp.Struct }
type Ghost_Which uint16

const (
	Ghost_Which_commit    Ghost_Which = 0
	Ghost_Which_directory Ghost_Which = 1
	Ghost_Which_file      Ghost_Which = 2
)

func (w Ghost_Which) String() string {
	const s = "commitdirectoryfile"
	switch w {
	case Ghost_Which_commit:
		return s[0:6]
	case Ghost_Which_directory:
		return s[6:15]
	case Ghost_Which_file:
		return s[15:19]

	}
	return "Ghost_Which(" + strconv.FormatUint(uint64(w), 10) + ")"
}

// Ghost_TypeID is the unique identifier for the type Ghost.
const Ghost_TypeID = 0x80c828d7e89c12ea

func NewGhost(s *capnp.Segment) (Ghost, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 16, PointerCount: 2})
	return Ghost{st}, err
}

func NewRootGhost(s *capnp.Segment) (Ghost, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 16, PointerCount: 2})
	return Ghost{st}, err
}

func ReadRootGhost(msg *capnp.Message) (Ghost, error) {
	root, err := msg.RootPtr()
	return Ghost{root.Struct()}, err
}

func (s Ghost) String() string {
	str, _ := text.Marshal(0x80c828d7e89c12ea, s.Struct)
	return str
}

func (s Ghost) Which() Ghost_Which {
	return Ghost_Which(s.Struct.Uint16(8))
}
func (s Ghost) GhostInode() uint64 {
	return s.Struct.Uint64(0)
}

func (s Ghost) SetGhostInode(v uint64) {
	s.Struct.SetUint64(0, v)
}

func (s Ghost) GhostPath() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s Ghost) HasGhostPath() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Ghost) GhostPathBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s Ghost) SetGhostPath(v string) error {
	return s.Struct.SetText(0, v)
}

func (s Ghost) Commit() (Commit, error) {
	if s.Struct.Uint16(8) != 0 {
		panic("Which() != commit")
	}
	p, err := s.Struct.Ptr(1)
	return Commit{Struct: p.Struct()}, err
}

func (s Ghost) HasCommit() bool {
	if s.Struct.Uint16(8) != 0 {
		return false
	}
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s Ghost) SetCommit(v Commit) error {
	s.Struct.SetUint16(8, 0)
	return s.Struct.SetPtr(1, v.Struct.ToPtr())
}

// NewCommit sets the commit field to a newly
// allocated Commit struct, preferring placement in s's segment.
func (s Ghost) NewCommit() (Commit, error) {
	s.Struct.SetUint16(8, 0)
	ss, err := NewCommit(s.Struct.Segment())
	if err != nil {
		return Commit{}, err
	}
	err = s.Struct.SetPtr(1, ss.Struct.ToPtr())
	return ss, err
}

func (s Ghost) Directory() (Directory, error) {
	if s.Struct.Uint16(8) != 1 {
		panic("Which() != directory")
	}
	p, err := s.Struct.Ptr(1)
	return Directory{Struct: p.Struct()}, err
}

func (s Ghost) HasDirectory() bool {
	if s.Struct.Uint16(8) != 1 {
		return false
	}
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s Ghost) SetDirectory(v Directory) error {
	s.Struct.SetUint16(8, 1)
	return s.Struct.SetPtr(1, v.Struct.ToPtr())
}

// NewDirectory sets the directory field to a newly
// allocated Directory struct, preferring placement in s's segment.
func (s Ghost) NewDirectory() (Directory, error) {
	s.Struct.SetUint16(8, 1)
	ss, err := NewDirectory(s.Struct.Segment())
	if err != nil {
		return Directory{}, err
	}
	err = s.Struct.SetPtr(1, ss.Struct.ToPtr())
	return ss, err
}

func (s Ghost) File() (File, error) {
	if s.Struct.Uint16(8) != 2 {
		panic("Which() != file")
	}
	p, err := s.Struct.Ptr(1)
	return File{Struct: p.Struct()}, err
}

func (s Ghost) HasFile() bool {
	if s.Struct.Uint16(8) != 2 {
		return false
	}
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s Ghost) SetFile(v File) error {
	s.Struct.SetUint16(8, 2)
	return s.Struct.SetPtr(1, v.Struct.ToPtr())
}

// NewFile sets the file field to a newly
// allocated File struct, preferring placement in s's segment.
func (s Ghost) NewFile() (File, error) {
	s.Struct.SetUint16(8, 2)
	ss, err := NewFile(s.Struct.Segment())
	if err != nil {
		return File{}, err
	}
	err = s.Struct.SetPtr(1, ss.Struct.ToPtr())
	return ss, err
}

// Ghost_List is a list of Ghost.
type Ghost_List struct{ capnp.List }

// NewGhost creates a new list of Ghost.
func NewGhost_List(s *capnp.Segment, sz int32) (Ghost_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 16, PointerCount: 2}, sz)
	return Ghost_List{l}, err
}

func (s Ghost_List) At(i int) Ghost { return Ghost{s.List.Struct(i)} }

func (s Ghost_List) Set(i int, v Ghost) error { return s.List.SetStruct(i, v.Struct) }

func (s Ghost_List) String() string {
	str, _ := text.MarshalList(0x80c828d7e89c12ea, s.List)
	return str
}

// Ghost_Promise is a wrapper for a Ghost promised by a client call.
type Ghost_Promise struct{ *capnp.Pipeline }

func (p Ghost_Promise) Struct() (Ghost, error) {
	s, err := p.Pipeline.Struct()
	return Ghost{s}, err
}

func (p Ghost_Promise) Commit() Commit_Promise {
	return Commit_Promise{Pipeline: p.Pipeline.GetPipeline(1)}
}

func (p Ghost_Promise) Directory() Directory_Promise {
	return Directory_Promise{Pipeline: p.Pipeline.GetPipeline(1)}
}

func (p Ghost_Promise) File() File_Promise {
	return File_Promise{Pipeline: p.Pipeline.GetPipeline(1)}
}

// Node is a node in the merkle dag of brig
type Node struct{ capnp.Struct }
type Node_Which uint16

const (
	Node_Which_commit    Node_Which = 0
	Node_Which_directory Node_Which = 1
	Node_Which_file      Node_Which = 2
	Node_Which_ghost     Node_Which = 3
)

func (w Node_Which) String() string {
	const s = "commitdirectoryfileghost"
	switch w {
	case Node_Which_commit:
		return s[0:6]
	case Node_Which_directory:
		return s[6:15]
	case Node_Which_file:
		return s[15:19]
	case Node_Which_ghost:
		return s[19:24]

	}
	return "Node_Which(" + strconv.FormatUint(uint64(w), 10) + ")"
}

// Node_TypeID is the unique identifier for the type Node.
const Node_TypeID = 0xa629eb7f7066fae3

func NewNode(s *capnp.Segment) (Node, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 16, PointerCount: 7})
	return Node{st}, err
}

func NewRootNode(s *capnp.Segment) (Node, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 16, PointerCount: 7})
	return Node{st}, err
}

func ReadRootNode(msg *capnp.Message) (Node, error) {
	root, err := msg.RootPtr()
	return Node{root.Struct()}, err
}

func (s Node) String() string {
	str, _ := text.Marshal(0xa629eb7f7066fae3, s.Struct)
	return str
}

func (s Node) Which() Node_Which {
	return Node_Which(s.Struct.Uint16(8))
}
func (s Node) Name() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s Node) HasName() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Node) NameBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s Node) SetName(v string) error {
	return s.Struct.SetText(0, v)
}

func (s Node) TreeHash() ([]byte, error) {
	p, err := s.Struct.Ptr(1)
	return []byte(p.Data()), err
}

func (s Node) HasTreeHash() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s Node) SetTreeHash(v []byte) error {
	return s.Struct.SetData(1, v)
}

func (s Node) ModTime() (string, error) {
	p, err := s.Struct.Ptr(2)
	return p.Text(), err
}

func (s Node) HasModTime() bool {
	p, err := s.Struct.Ptr(2)
	return p.IsValid() || err != nil
}

func (s Node) ModTimeBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(2)
	return p.TextBytes(), err
}

func (s Node) SetModTime(v string) error {
	return s.Struct.SetText(2, v)
}

func (s Node) Inode() uint64 {
	return s.Struct.Uint64(0)
}

func (s Node) SetInode(v uint64) {
	s.Struct.SetUint64(0, v)
}

func (s Node) ContentHash() ([]byte, error) {
	p, err := s.Struct.Ptr(3)
	return []byte(p.Data()), err
}

func (s Node) HasContentHash() bool {
	p, err := s.Struct.Ptr(3)
	return p.IsValid() || err != nil
}

func (s Node) SetContentHash(v []byte) error {
	return s.Struct.SetData(3, v)
}

func (s Node) User() (string, error) {
	p, err := s.Struct.Ptr(4)
	return p.Text(), err
}

func (s Node) HasUser() bool {
	p, err := s.Struct.Ptr(4)
	return p.IsValid() || err != nil
}

func (s Node) UserBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(4)
	return p.TextBytes(), err
}

func (s Node) SetUser(v string) error {
	return s.Struct.SetText(4, v)
}

func (s Node) Commit() (Commit, error) {
	if s.Struct.Uint16(8) != 0 {
		panic("Which() != commit")
	}
	p, err := s.Struct.Ptr(5)
	return Commit{Struct: p.Struct()}, err
}

func (s Node) HasCommit() bool {
	if s.Struct.Uint16(8) != 0 {
		return false
	}
	p, err := s.Struct.Ptr(5)
	return p.IsValid() || err != nil
}

func (s Node) SetCommit(v Commit) error {
	s.Struct.SetUint16(8, 0)
	return s.Struct.SetPtr(5, v.Struct.ToPtr())
}

// NewCommit sets the commit field to a newly
// allocated Commit struct, preferring placement in s's segment.
func (s Node) NewCommit() (Commit, error) {
	s.Struct.SetUint16(8, 0)
	ss, err := NewCommit(s.Struct.Segment())
	if err != nil {
		return Commit{}, err
	}
	err = s.Struct.SetPtr(5, ss.Struct.ToPtr())
	return ss, err
}

func (s Node) Directory() (Directory, error) {
	if s.Struct.Uint16(8) != 1 {
		panic("Which() != directory")
	}
	p, err := s.Struct.Ptr(5)
	return Directory{Struct: p.Struct()}, err
}

func (s Node) HasDirectory() bool {
	if s.Struct.Uint16(8) != 1 {
		return false
	}
	p, err := s.Struct.Ptr(5)
	return p.IsValid() || err != nil
}

func (s Node) SetDirectory(v Directory) error {
	s.Struct.SetUint16(8, 1)
	return s.Struct.SetPtr(5, v.Struct.ToPtr())
}

// NewDirectory sets the directory field to a newly
// allocated Directory struct, preferring placement in s's segment.
func (s Node) NewDirectory() (Directory, error) {
	s.Struct.SetUint16(8, 1)
	ss, err := NewDirectory(s.Struct.Segment())
	if err != nil {
		return Directory{}, err
	}
	err = s.Struct.SetPtr(5, ss.Struct.ToPtr())
	return ss, err
}

func (s Node) File() (File, error) {
	if s.Struct.Uint16(8) != 2 {
		panic("Which() != file")
	}
	p, err := s.Struct.Ptr(5)
	return File{Struct: p.Struct()}, err
}

func (s Node) HasFile() bool {
	if s.Struct.Uint16(8) != 2 {
		return false
	}
	p, err := s.Struct.Ptr(5)
	return p.IsValid() || err != nil
}

func (s Node) SetFile(v File) error {
	s.Struct.SetUint16(8, 2)
	return s.Struct.SetPtr(5, v.Struct.ToPtr())
}

// NewFile sets the file field to a newly
// allocated File struct, preferring placement in s's segment.
func (s Node) NewFile() (File, error) {
	s.Struct.SetUint16(8, 2)
	ss, err := NewFile(s.Struct.Segment())
	if err != nil {
		return File{}, err
	}
	err = s.Struct.SetPtr(5, ss.Struct.ToPtr())
	return ss, err
}

func (s Node) Ghost() (Ghost, error) {
	if s.Struct.Uint16(8) != 3 {
		panic("Which() != ghost")
	}
	p, err := s.Struct.Ptr(5)
	return Ghost{Struct: p.Struct()}, err
}

func (s Node) HasGhost() bool {
	if s.Struct.Uint16(8) != 3 {
		return false
	}
	p, err := s.Struct.Ptr(5)
	return p.IsValid() || err != nil
}

func (s Node) SetGhost(v Ghost) error {
	s.Struct.SetUint16(8, 3)
	return s.Struct.SetPtr(5, v.Struct.ToPtr())
}

// NewGhost sets the ghost field to a newly
// allocated Ghost struct, preferring placement in s's segment.
func (s Node) NewGhost() (Ghost, error) {
	s.Struct.SetUint16(8, 3)
	ss, err := NewGhost(s.Struct.Segment())
	if err != nil {
		return Ghost{}, err
	}
	err = s.Struct.SetPtr(5, ss.Struct.ToPtr())
	return ss, err
}

func (s Node) BackendHash() ([]byte, error) {
	p, err := s.Struct.Ptr(6)
	return []byte(p.Data()), err
}

func (s Node) HasBackendHash() bool {
	p, err := s.Struct.Ptr(6)
	return p.IsValid() || err != nil
}

func (s Node) SetBackendHash(v []byte) error {
	return s.Struct.SetData(6, v)
}

// Node_List is a list of Node.
type Node_List struct{ capnp.List }

// NewNode creates a new list of Node.
func NewNode_List(s *capnp.Segment, sz int32) (Node_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 16, PointerCount: 7}, sz)
	return Node_List{l}, err
}

func (s Node_List) At(i int) Node { return Node{s.List.Struct(i)} }

func (s Node_List) Set(i int, v Node) error { return s.List.SetStruct(i, v.Struct) }

func (s Node_List) String() string {
	str, _ := text.MarshalList(0xa629eb7f7066fae3, s.List)
	return str
}

// Node_Promise is a wrapper for a Node promised by a client call.
type Node_Promise struct{ *capnp.Pipeline }

func (p Node_Promise) Struct() (Node, error) {
	s, err := p.Pipeline.Struct()
	return Node{s}, err
}

func (p Node_Promise) Commit() Commit_Promise {
	return Commit_Promise{Pipeline: p.Pipeline.GetPipeline(5)}
}

func (p Node_Promise) Directory() Directory_Promise {
	return Directory_Promise{Pipeline: p.Pipeline.GetPipeline(5)}
}

func (p Node_Promise) File() File_Promise {
	return File_Promise{Pipeline: p.Pipeline.GetPipeline(5)}
}

func (p Node_Promise) Ghost() Ghost_Promise {
	return Ghost_Promise{Pipeline: p.Pipeline.GetPipeline(5)}
}

const schema_9195d073cb5c5953 = "x\xda\xb4V_l\x14\xd5\x17>\xdf\xbd3;,)" +
	"\xbf\xdd\xfdMI\x08\xf9\x95\xbd!\x90_!\x15Z\x8a" +
	"\x916\x18(\xb4\xd2b!\xbdl\x89@\xd08\xec\xde" +
	"\xeeL\xd8\x9d)3\x83\xb5\xc6\x064\x98\xa0\x06\x03\x11" +
	"\x13IJ\x04S\xff%\x1a|\xf1\xd1\x98\x98h\x8c\xbe" +
	"\x18\x1f4\xf1QM4\x9a\xf8,F\x18s\xf7oi" +
	"*\xf0\xe2[\xf7;\xe7\xde\xfb\x9d\xef|\xe7L{9" +
	"\xdf\xcd\xfa\xcc\xff\x1bD\xb2\xd7L%\xbf\xfew\xfe\x97" +
	"\xef\xba\xbf8Kr\x1dXR8z\xfc\xab\xe8\xeb\xd7" +
	".\xd1\x08\xb38\x8c\xfe\x9bX\x0f;\xcd,;\xcd\xf2" +
	"\xfd#,\x0fBr5\xff\xe8\xccS\xbf\xaf~\x99r" +
	"\xeb\xd0>`2\x8b\xa8\xdf\xe1\x83\xb0Oq\xcb>\xc5" +
	"\xf3\xf6U>CHn<>\xe9\x7fn_\xbb\xa0\x1f" +
	"X\x9c\x9f\xd2\xf97\xf9f\xd8i\xc3\xb2\xd3F\xbe\x7f" +
	"\xc0xL\xdf\x7f\xb8\xef\xc5\x87\x1e\x1ex\xe7\x95\xa5\x07" +
	"j\x0fx\xe6Z\xd8\xb3\xa6e\xcf\x9ay\xfb\xbay\x83" +
	"\x90\xfc\xf8\xe7\xd4\xf4\x99\xdf6\xbd\xbd\xb4\x02\xcb2`" +
	"\xf4\x0f\xa4\xd6\xc2\x1eKY\xf6X*\xdf?\x97\x0a\x18" +
	"!Y\xf8i\xfc\xfb\xcc\xc2\x1f\x9f\x90\xdc\x88E\x04W" +
	"\xa7,\x10\xf5\x9b\xe9c \xd8\xb9\xb4f\x8f\xf9\xe7+" +
	"\xbdG\xc7\x7fXJ\x86k2\xd5\xf4\x1e\xd8si\xcb" +
	"\x9eK\xe7\xed\x8f\xd2?\xd3\x8e\xa4\xe8\xc4S\xd1V?" +
	"\xe0%\x15m-:\xd3\xfe\xf4V?(\xa9hK\xed" +
	"\xef\xc1}\xae\x15D\xf1\x04 \x0d\xb0\xe4\x89W\xdf\x90" +
	"\x1f\x7f\xfb\xd2g$\x0d\x86\xa1\x1e\xa0\x83\xa8\x0f\xdf " +
	"\xd9\xe7\x06Q,<?U\xf2\x8aN\xac\"\x11\xbbN" +
	",\x1cQTa\xecx\xbe\xd0W\x8a\x19'\x12N," +
	"b\xd7\x8b\xc4\xb4\x13\xbb\"\xf0\x8bPD\xb2\x93\x1bD" +
	"\x06\x88rs\xc7\x88\xe4\xb3\x1c\xf2<\x03\xd0\x09\x8d\xbd" +
	"p\x88H\x9e\xe3\x90\x17\x19\xbaX\x92\xa0\x13\x8c(w" +
	"a\x90H\x9e\xe7\x90\x97\x19\xba\xf8m\x0ds\xa2\xdc%" +
	"\x9d}\x91C\xce3t\x19\xb74l\x10\xe5\xael&" +
	"\x92\x979\xe45\x86\xa4\xac\xd9\x8e\xf9\x01\xf1\x92B\x9a" +
	"\x18\xd2\xd4\x00'\x9c\x98\xe0\xa2\x83\x18:\x08\xbb\x8aA" +
	"\xb5\xea\xc5\xc8\xb6%' KHJ^\xa8\x8aq\x10" +
	"\x12f\x91mk^\x8ff\xa6\xbc\x8aB\xb6\xed\x8b\xc6" +
	"\xa1{H=\xec\xed\x0aG\xfc8\x9c]^\xed\xff\xd5" +
	"\xd4\xce\xe1\xcbdHD\x9e_\xae(&\x9a4f\x85" +
	"\xd2\x07\x09rEK\xcaM\xba\xe2\x0d\x1c\xb2\x97!\xd7" +
	"\xd4\xf2\x01\x0dvs\xc8\xed\x0c\x19\xdf\xa9\xaaf\xa9\x19" +
	"\xd7\x89\\\xac\"\x86U\xf7f\xba7\xc8h]\x96\xe7" +
	")\x1a\xaeX\x8fdoM>\xe1\xf1H8\"R\xb1" +
	"\x08\xa6D\xd1u\xfc\xb26H \xfc\xc0*\xa9\x88H" +
	"\xaei\x91\xbe\xb2\xa7\xdd\xa6\x16\xe9\xab\xba\xd3\xafs\xc8" +
	"\x05\x86\x1cc\xf5\xf6_\xd7\xe0<\x87|\x97!\xc7y" +
	"\xbd\xf9o\xe9\xf2\xaeq\xc8\xf7\x19`\xd4;\xff\xde6" +
	"\"\xb9\xc0!?d\x80\x89E\xc3\x94\xfb`\x1b\xb13" +
	"U\x15EN\xb9%\xc4.\xe7t\xec\x06a\xeb\xe7\xb4" +
	"\x13*?n*\x93\x09\x83\xa0\xf5#\xef\xf9%\xf54" +
	"Lb0\x09\xf9\xaa\x0a\xcb\xea^\xd2=\xe2\xf1\x8aZ" +
	"^\xb85\x8d\x06\x7f\x9a\x0c\x89\x8ar\xa6\x84\xcf\xf4\xd4" +
	"x\xbe\x88]%\x0e\x0c\x0f\xed#\"\xd9\xd1\xd2jD" +
	"\x17\xbb\x9bC\x8e\xb7geL\xab2\xcc!'\xb4T" +
	"\x8dI9\xb0\x9eH\x8er\xc8I\x86L\xe4=\xd3\xf2" +
	"|\xb3\xb8F\xad\xd6I5{\xbf\x168\xa8\x03\xcb\xd7" +
	"\xb1\xa1a\x80\xfdH\x0e\xd6\x0a\x88\x84\xe1\xd47@\xa3" +
	"\x96\xaa\x0aOV\x94(9e\xed\x88\x13\xa1W&\xc8" +
	"\x9efa\xf6Fl&*\x08p\x14z\xd0\xf6\x81\xbd" +
	"\x09\xfb\x89\x0a\xdd\x1a\xdf\x8e\xb6\x15\xec>\xec!*\xf4" +
	"h|\x07\x18P7\x83\xfd \xb6\x11\x15z5\xbcS" +
	"\xa7\x1b\xbcf\x08{\x00'\x88\x0a;4>\xacq\xd3" +
	"\xe8\x84Id\x0f\xd5\x9e\xdd\xa9\xf1Q0t\xa5\x92\xc4" +
	"\xecD\x8a\xc8\x1e\xc1 Qa\xb7\x8e\x8c\xeb\x88u[" +
	"G,\"{\x0c\x87\x88\x0a\xa3:2\xa9#+n\xe9" +
	"\xc8\x0a\"[\xd6n\x1b\xd7\x91#:\x92\xfeKG\xd2" +
	"D\xf6\xe1\x1a\xaf\x09\x1d9\xae\xdf_\x99\xea\xc4J\"" +
	"\xfbh\x8d\xd7\x11\x8d\x97\xb0d<\x938Tj\xd4\x89" +
	"\\\"j\xb6\xe8L5(Mz\xed\x9c\xbc\xa75n" +
	"\xed\xb3b\xe0\xc7\xca\x8fG\xc9Z4\xd9\x99\xd3\x91\x0a" +
	"\xff\x9d\xf5\x96\xaf-Pd\xdb\x1f\xe8\xc6e'\x9c\xe2" +
	"I\xe5\x97\xee$\xd2\xf2\x97\xf1O+FS\xdbRU" +
	"!/+\xbd\xd5\xb2\xf5.-Yk\xf5\x06\xdd\xb9\xd6" +
	"f\xbc\xd8m\xaf5\xe5\x94\xee\xf7\xcd\xe1\xe66\xa5\xe5" +
	"\x8d\xdd\xdd0\xf6\x9bH\x9a\xa9\xe6\xac\xd0:;\x9e\x1f" +
	"\x89\xc0W\"\x08E5\x08Uk1{*\xd2\xd8\x94" +
	"gUj\x9b.\xdb\x9a^GS>\xce!\xdd\xf6\xf4" +
	"*=\xbdOr\xc8\xca\xa2\xe9\xf5\xf6\x13I\x97C\x9e" +
	"\xd3\x8b\x8e\xd5\x17\xdds\x1a<[\xff\xca\xddm\xa4\x93" +
	"\xa2\xebUJ\xa1\xf2\xb5o\xfeC\x98\xe0@\xb6\xfd\x0f" +
	"\x11A\x83M\xabDwK\xfa;\x00\x00\xff\xffL\xce" +
	"@$"

func init() {
	schemas.Register(schema_9195d073cb5c5953,
		0x80c828d7e89c12ea,
		0x8b15ee76774b1f9d,
		0x8da013c66e545daf,
		0x8ea7393d37893155,
		0xa629eb7f7066fae3,
		0xbff8a40fda4ce4a4,
		0xe24c59306c829c01)
}
