package parser

type CreateTableStmt struct {
	Name        string
	PK          []string
	CK          [][]string
	IfNotExists bool
	Ordering    []OrderBy
	Options     map[string]string // make it struct
	Compaction  map[string]string
	Compression map[string]string

	BloomFilterFpChance    float32 //bloom_filter_fp_chance rec: 0.1; def 0.01
	CDC                    bool    // Change Data Capture
	Comments               string
	DLocalReadRepairChance float32 // dclocal_read_repair_chance
	DefaultTimeToLive      int64   // default_time_to_live
	GCGraceSeconds         int64   // gc_grace_seconds // Java
	MemtableFlushPeriodMs  int64   // memtable_flush_period_in_ms
	MinIndexInterval       int64   // min_index_interval
	MaxIndexInterval       int64   // max_index_interval
	ReadRepairChance       float64 // read_repair_chance
	SpeculativeRetry       string  // speculative_retry
}

type OrderBy struct {
	Name string
	Desc bool
}

type TableOptions struct {
}
