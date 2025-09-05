package dto

import (
	"fmt"
	"strconv"

	commonv1 "github.com/ose-micro/common/gen/go/ose/micro/common/v1"
)

func BuildRequestDTO(query *commonv1.Request) (*Request, error) {
	if query == nil {
		return nil, fmt.Errorf("query is nil")
	}

	facets := make([]Query, len(query.Facets))

	for i, facet := range query.Facets {

		filters := make([]Filter, len(facet.Filters))

		for i, filter := range facet.Filters {
			filters[i] = Filter{
				Field: filter.Field,
				Op: func() FilterOp {
					switch filter.Op {
					case commonv1.FilterOp_FILTER_OP_EQ:
						return OpEq
					case commonv1.FilterOp_FILTER_OP_GTE:
						return OpGte
					case commonv1.FilterOp_FILTER_OP_GT:
						return OpGt
					case commonv1.FilterOp_FILTER_OP_LT:
						return OpLt
					case commonv1.FilterOp_FILTER_OP_LTE:
						return OpLte
					case commonv1.FilterOp_FILTER_OP_IN:
						return OpIn
					case commonv1.FilterOp_FILTER_OP_NE:
						return OpNe
					case commonv1.FilterOp_FILTER_OP_NIN:
						return OpNin
					default:
						return OpEq
					}
				}(),
				Value: func() interface{} {
					switch filter.Op {
					case commonv1.FilterOp_FILTER_OP_EQ, commonv1.FilterOp_FILTER_OP_IN, commonv1.FilterOp_FILTER_OP_NE,
						commonv1.FilterOp_FILTER_OP_NIN:
						return filter.Value
					case commonv1.FilterOp_FILTER_OP_GTE, commonv1.FilterOp_FILTER_OP_GT, commonv1.FilterOp_FILTER_OP_LT,
						commonv1.FilterOp_FILTER_OP_LTE:
						value, err := strconv.ParseFloat(filter.Value, 32)
						if err != nil {
							return nil
						}
						return float32(value)

					default:
						return OpEq
					}
				}(),
			}
		}

		aggregations := make([]Aggregation, len(facet.Aggregations))
		for i, agg := range facet.Aggregations {
			aggregations[i] = Aggregation{
				Field: agg.Field,
				As:    agg.As,
				Type: func() AggregationType {
					switch agg.Type {
					case commonv1.AggregationType_AGGREGATION_TYPE_SUM:
						return AggSum
					case commonv1.AggregationType_AGGREGATION_TYPE_AVG:
						return AggAvg
					case commonv1.AggregationType_AGGREGATION_TYPE_COUNT:
						return AggCount
					case commonv1.AggregationType_AGGREGATION_TYPE_MAX:
						return AggMax
					case commonv1.AggregationType_AGGREGATION_TYPE_MIN:
						return AggMin
					default:
						return AggCount
					}
				}(),
			}
		}

		sorts := make([]SortOption, len(facet.Sort))
		for i, sort := range facet.Sort {
			sorts[i] = SortOption{
				Field: sort.Field,
				Order: func() SortOrder {
					switch sort.Order {
					case commonv1.SortOrder_SORT_ORDER_ASC:
						return SortAsc
					case commonv1.SortOrder_SORT_ORDER_DESC:
						return SortDesc
					default:
						return SortAsc
					}
				}(),
			}
		}

		computedFields := make([]ComputedField, len(facet.ComputedFields))
		for i, computed := range facet.ComputedFields {
			computedFields[i] = ComputedField{
				Name: computed.Name,
				Operator: func() ComputedOperator {
					switch computed.Operator {
					case commonv1.ComputedOperator_COMPUTED_OPERATOR_ADD:
						return OpAdd
					case commonv1.ComputedOperator_COMPUTED_OPERATOR_SUBTRACT:
						return OpSubtract
					case commonv1.ComputedOperator_COMPUTED_OPERATOR_MULTIPLY:
						return OpMultiply
					case commonv1.ComputedOperator_COMPUTED_OPERATOR_DIVIDE:
						return OpDivide
					case commonv1.ComputedOperator_COMPUTED_OPERATOR_CONCAT:
						return OpConcat
					case commonv1.ComputedOperator_COMPUTED_OPERATOR_DATE_TRUNC:
						return OpDateTrunc
					case commonv1.ComputedOperator_COMPUTED_OPERATOR_DAY_OF_MONTH:
						return OpDateTrunc
					case commonv1.ComputedOperator_COMPUTED_OPERATOR_IF_NULL:
						return OpIfNull
					case commonv1.ComputedOperator_COMPUTED_OPERATOR_MONTH:
						return OpDayOfMonth
					case commonv1.ComputedOperator_COMPUTED_OPERATOR_SUBSTR:
						return OpSubstr
					case commonv1.ComputedOperator_COMPUTED_OPERATOR_TO_LOWER:
						return OpToLower
					case commonv1.ComputedOperator_COMPUTED_OPERATOR_TO_UPPER:
						return OpToUpper
					case commonv1.ComputedOperator_COMPUTED_OPERATOR_YEAR:
						return OpYear
					default:
						return OpAdd
					}
				}(),
				Operands: computed.Operands,
			}
		}

		facets[i] = Query{
			Name:           facet.Name,
			Filters:        filters,
			GroupBy:        facet.GroupBy,
			Aggregations:   aggregations,
			Sort:           sorts,
			ComputedFields: computedFields,
			Skip:           facet.Skip,
			Limit:          facet.Limit,
		}
	}

	return &Request{
		Queries: facets,
	}, nil
}

func BuildGRPCRequest(query *Request) (*commonv1.Request, error) {
	if query == nil {
		return nil, fmt.Errorf("query is nil")
	}

	facets := make([]*commonv1.FacetQuery, len(query.Queries))

	for i, facet := range query.Queries {

		filters := make([]*commonv1.Filter, len(facet.Filters))

		for i, filter := range facet.Filters {
			filters[i] = &commonv1.Filter{
				Field: filter.Field,
				Op: func() commonv1.FilterOp {
					switch filter.Op {
					case OpEq:
						return commonv1.FilterOp_FILTER_OP_EQ
					case OpGte:
						return commonv1.FilterOp_FILTER_OP_GTE
					case OpGt:
						return commonv1.FilterOp_FILTER_OP_GT
					case OpLt:
						return commonv1.FilterOp_FILTER_OP_LT
					case OpLte:
						return commonv1.FilterOp_FILTER_OP_LTE
					case OpIn:
						return commonv1.FilterOp_FILTER_OP_IN
					case OpNe:
						return commonv1.FilterOp_FILTER_OP_NE
					case OpNin:
						return commonv1.FilterOp_FILTER_OP_NIN
					default:
						return commonv1.FilterOp_FILTER_OP_UNSPECIFIED
					}
				}(),
				Value: filter.Value.(string),
			}
		}

		aggregations := make([]*commonv1.Aggregation, len(facet.Aggregations))
		for i, agg := range facet.Aggregations {
			aggregations[i] = &commonv1.Aggregation{
				Field: agg.Field,
				As:    agg.As,
				Type: func() commonv1.AggregationType {
					switch agg.Type {
					case AggSum:
						return commonv1.AggregationType_AGGREGATION_TYPE_SUM
					case AggAvg:
						return commonv1.AggregationType_AGGREGATION_TYPE_AVG
					case AggCount:
						return commonv1.AggregationType_AGGREGATION_TYPE_COUNT
					case AggMax:
						return commonv1.AggregationType_AGGREGATION_TYPE_MAX
					case AggMin:
						return commonv1.AggregationType_AGGREGATION_TYPE_MIN
					default:
						return commonv1.AggregationType_AGGREGATION_TYPE_UNSPECIFIED
					}
				}(),
			}
		}

		sorts := make([]*commonv1.SortOption, len(facet.Sort))
		for i, sort := range facet.Sort {
			sorts[i] = &commonv1.SortOption{
				Field: sort.Field,
				Order: func() commonv1.SortOrder {
					switch sort.Order {
					case SortAsc:
						return commonv1.SortOrder_SORT_ORDER_ASC
					case SortDesc:
						return commonv1.SortOrder_SORT_ORDER_DESC
					default:
						return commonv1.SortOrder_SORT_ORDER_UNSPECIFIED
					}
				}(),
			}
		}

		computedFields := make([]*commonv1.ComputedField, len(facet.ComputedFields))
		for i, computed := range facet.ComputedFields {
			computedFields[i] = &commonv1.ComputedField{
				Name: computed.Name,
				Operator: func() commonv1.ComputedOperator {
					switch computed.Operator {
					case OpAdd:
						return commonv1.ComputedOperator_COMPUTED_OPERATOR_ADD
					case OpSubtract:
						return commonv1.ComputedOperator_COMPUTED_OPERATOR_SUBTRACT
					case OpMultiply:
						return commonv1.ComputedOperator_COMPUTED_OPERATOR_MULTIPLY
					case OpDivide:
						return commonv1.ComputedOperator_COMPUTED_OPERATOR_DIVIDE
					case OpConcat:
						return commonv1.ComputedOperator_COMPUTED_OPERATOR_CONCAT
					case OpDateTrunc:
						return commonv1.ComputedOperator_COMPUTED_OPERATOR_DATE_TRUNC
					case OpDayOfMonth:
						return commonv1.ComputedOperator_COMPUTED_OPERATOR_DAY_OF_MONTH
					case OpIfNull:
						return commonv1.ComputedOperator_COMPUTED_OPERATOR_IF_NULL
					case OpMonth:
						return commonv1.ComputedOperator_COMPUTED_OPERATOR_MONTH
					case OpSubstr:
						return commonv1.ComputedOperator_COMPUTED_OPERATOR_SUBSTR
					case OpToLower:
						return commonv1.ComputedOperator_COMPUTED_OPERATOR_TO_LOWER
					case OpToUpper:
						return commonv1.ComputedOperator_COMPUTED_OPERATOR_TO_UPPER
					case OpYear:
						return commonv1.ComputedOperator_COMPUTED_OPERATOR_YEAR
					default:
						return commonv1.ComputedOperator_COMPUTED_OPERATOR_UNSPECIFIED
					}
				}(),
				Operands: computed.Operands,
			}
		}

		facets[i] = &commonv1.FacetQuery{
			Name:           facet.Name,
			Filters:        filters,
			GroupBy:        facet.GroupBy,
			Aggregations:   aggregations,
			Sort:           sorts,
			ComputedFields: computedFields,
			Skip:           facet.Skip,
			Limit:          facet.Limit,
		}
	}

	return &commonv1.Request{
		Facets: facets,
	}, nil
}
