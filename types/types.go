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

func (s Int8) Less(other Sortable) bool {
	if o, ok := other.(Int8); ok {
		return s < o
	} else {
		return false
	}
}

func (s UInt8) Less(other Sortable) bool {
	if o, ok := other.(UInt8); ok {
		return s < o
	} else {
		return false
	}
}

func (s Int16) Less(other Sortable) bool {
	if o, ok := other.(Int16); ok {
		return s < o
	} else {
		return false
	}
}

func (s UInt16) Less(other Sortable) bool {
	if o, ok := other.(UInt16); ok {
		return s < o
	} else {
		return false
	}
}

func (s Int32) Less(other Sortable) bool {
	if o, ok := other.(Int32); ok {
		return s < o
	} else {
		return false
	}
}

func (s UInt32) Less(other Sortable) bool {
	if o, ok := other.(UInt32); ok {
		return s < o
	} else {
		return false
	}
}

func (s Int64) Less(other Sortable) bool {
	if o, ok := other.(Int64); ok {
		return s < o
	} else {
		return false
	}
}

func (s UInt64) Less(other Sortable) bool {
	if o, ok := other.(UInt64); ok {
		return s < o
	} else {
		return false
	}
}

func (s Int) Less(other Sortable) bool {
	if o, ok := other.(Int); ok {
		return s < o
	} else {
		return false
	}
}

func (s UInt) Less(other Sortable) bool {
	if o, ok := other.(UInt); ok {
		return s < o
	} else {
		return false
	}
}
