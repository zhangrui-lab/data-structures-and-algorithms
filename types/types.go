package types

type Int8 int8
type UInt8 uint8
type Int16 int16
type UInt16 uint16
type Int32 int32
type UInt32 uint32
type Int64 int64
type UInt64 uint64
type Int int
type UInt uint

func (self Int8) Less(other Sortable) bool {
	if o, ok := other.(Int8); ok {
		return self < o
	} else {
		return false
	}
}

func (self UInt8) Less(other Sortable) bool {
	if o, ok := other.(UInt8); ok {
		return self < o
	} else {
		return false
	}
}

func (self Int16) Less(other Sortable) bool {
	if o, ok := other.(Int16); ok {
		return self < o
	} else {
		return false
	}
}

func (self UInt16) Less(other Sortable) bool {
	if o, ok := other.(UInt16); ok {
		return self < o
	} else {
		return false
	}
}

func (self Int32) Less(other Sortable) bool {
	if o, ok := other.(Int32); ok {
		return self < o
	} else {
		return false
	}
}

func (self UInt32) Less(other Sortable) bool {
	if o, ok := other.(UInt32); ok {
		return self < o
	} else {
		return false
	}
}

func (self Int64) Less(other Sortable) bool {
	if o, ok := other.(Int64); ok {
		return self < o
	} else {
		return false
	}
}

func (self UInt64) Less(other Sortable) bool {
	if o, ok := other.(UInt64); ok {
		return self < o
	} else {
		return false
	}
}

func (self Int) Less(other Sortable) bool {
	if o, ok := other.(Int); ok {
		return self < o
	} else {
		return false
	}
}

func (self UInt) Less(other Sortable) bool {
	if o, ok := other.(UInt); ok {
		return self < o
	} else {
		return false
	}
}
