package common

type FilterOp string

const (
	OpEq    FilterOp = "eq"
	OpNe    FilterOp = "ne"
	OpGt    FilterOp = "gt"
	OpGte   FilterOp = "gte"
	OpLt    FilterOp = "lt"
	OpLte   FilterOp = "lte"
	OpIn    FilterOp = "in"
	OpNin   FilterOp = "nin"
	OpRegex FilterOp = "regex"
)

type Filter struct {
	Field string      `json:"field" validate:"required"`
	Op    FilterOp    `json:"op" validate:"required"`
	Value interface{} `json:"value" validate:"required"`
}

type AggregationType string

const (
	AggSum   AggregationType = "sum"
	AggAvg   AggregationType = "avg"
	AggCount AggregationType = "count"
	AggMin   AggregationType = "min"
	AggMax   AggregationType = "max"
)

type Aggregation struct {
	Type  AggregationType `json:"type" validate:"required"`
	Field string          `json:"field" validate:"required"`
	As    string          `json:"as" validate:"required"`
}

type SortOrder string

const (
	SortAsc  SortOrder = "asc"
	SortDesc SortOrder = "desc"
)

type SortOption struct {
	Field string    `json:"field" validate:"required"`
	Order SortOrder `json:"order" validate:"required"`
}

type ComputedOperator string

const (
	OpMultiply   ComputedOperator = "multiply"
	OpAdd        ComputedOperator = "add"
	OpDivide     ComputedOperator = "divide"
	OpSubtract   ComputedOperator = "subtract"
	OpWeek       ComputedOperator = "week"
	OpYear       ComputedOperator = "year"
	OpMonth      ComputedOperator = "month"
	OpDayOfMonth ComputedOperator = "dayOfMonth"
	OpToUpper    ComputedOperator = "toUpper"
	OpToLower    ComputedOperator = "toLower"
	OpConcat     ComputedOperator = "concat"
	OpSubstr     ComputedOperator = "substr"
	OpDateTrunc  ComputedOperator = "dateTrunc"
	OpIfNull     ComputedOperator = "ifNull"
)

type ComputedField struct {
	Name     string           `json:"name" validate:"required"`
	Operator ComputedOperator `json:"operator" validate:"required"`
	Operands []string         `json:"operands" validate:"required,min=1"`
}

type QueryType string

const (
	QueryTypeRecord  QueryType = "record"
	QueryTypeSummary QueryType = "summary"
)

type Summary struct {
	ID     any                    `json:"id,omitempty"`
	Fields map[string]interface{} `json:"fields,omitempty"`
}

type Query struct {
	Name           string          `json:"name" validate:"required"`
	Filters        []Filter        `json:"filters"`
	Type           QueryType       `json:"type"`
	GroupBy        []string        `json:"group_by"`
	Aggregations   []Aggregation   `json:"aggregations"`
	Sort           []SortOption    `json:"sort"`
	ComputedFields []ComputedField `json:"computed_fields"`
	Skip           int64           `json:"skip"`
	Limit          int64           `json:"limit"`
	ReturnType     string          `json:"return_type"`
}

type Request struct {
	Queries []Query `json:"facets"`
}
