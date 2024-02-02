package gorm_repo

import (
	"fmt"
	"strings"

	"github.com/rum-people-preseed/weather-harvester-svc/internal/domain/repository"
	"github.com/rum-people-preseed/weather-harvester-svc/internal/interfaces/generics"
)

type joinAndPreload struct {
	joins   []string
	preload []string
}

func (s joinAndPreload) Joins() []string {
	if s.joins == nil {
		s.joins = []string{}
	}
	return s.joins
}

func (s joinAndPreload) Preload() []string {
	if s.preload == nil {
		s.preload = []string{}
	}
	return s.preload
}

type joinSpecification struct {
	specifications []repository.Specification
	separator      string
}

func (s joinSpecification) Sort() []string {
	return []string{}
}

func (s joinSpecification) Joins() []string {
	var joins []string
	for _, spec := range s.specifications {
		for _, join := range spec.Joins() {
			if !generics.IsInSlice(join, joins) {
				joins = append(joins, join)
			}
		}
	}
	return joins
}

func (s joinSpecification) Preload() []string {
	var preloads []string
	for _, spec := range s.specifications {
		for _, preload := range spec.Preload() {
			if !generics.IsInSlice(preload, preloads) {
				preloads = append(preloads, preload)
			}
		}
	}
	return preloads
}

func (s joinSpecification) GetQuery() string {
	queries := make([]string, 0)
	for _, spec := range s.specifications {
		if spec.GetQuery() == "" {
			continue
		}
		queries = append(queries, spec.GetQuery())
	}
	return strings.Join(queries, fmt.Sprintf(" %s ", s.separator))
}

func (s joinSpecification) GetValues() []any {
	values := make([]any, 0)
	for _, spec := range s.specifications {
		for _, val := range spec.GetValues() {
			if val == "" {
				continue
			}
			values = append(values, val)
		}
	}
	return values
}

func And(specifications ...repository.Specification) repository.Specification {
	return joinSpecification{
		specifications: specifications,
		separator:      "AND",
	}
}

func Or(specifications ...repository.Specification) repository.Specification {
	return joinSpecification{
		specifications: specifications,
		separator:      "OR",
	}
}

type notSpecification struct {
	repository.Specification
}

func (s notSpecification) GetQuery() string {
	return fmt.Sprintf(" NOT (%s)", s.Specification.GetQuery())
}

func Not(specification repository.Specification) repository.Specification {
	return notSpecification{
		specification,
	}
}

type binaryOperatorSpecification[T any] struct {
	joinAndPreload
	field    string
	operator string
	value    T
	sortBy   string
	sortDir  string
}

func (s binaryOperatorSpecification[T]) Sort() []string {
	return []string{s.sortBy, s.sortDir}
}

func (s binaryOperatorSpecification[T]) GetQuery() string {
	if s.field == "" {
		return ""
	}
	return fmt.Sprintf("%s %s ?", s.field, s.operator)
}

func (s binaryOperatorSpecification[T]) GetValues() []any {
	return []any{s.value}
}

func EqualWithJoinEqual[T any](field string, value T, joins []string, preloads []string) repository.Specification {
	return binaryOperatorSpecification[T]{
		joinAndPreload: joinAndPreload{
			joins:   joins,
			preload: preloads,
		},
		field:    field,
		operator: "=",
		value:    value,
	}
}

func Equal[T any](field string, value T) repository.Specification {
	return binaryOperatorSpecification[T]{
		field:    field,
		operator: "=",
		value:    value,
	}
}

type periodSpecification[T any] struct {
	joinAndPreload
	field    string
	operator string
	start    T
	end      T
}

func (p periodSpecification[T]) Sort() []string {
	return nil
}

func (p periodSpecification[T]) GetQuery() string {
	if p.field == "" {
		return ""
	}
	return fmt.Sprintf("%s %s ? AND ? ", p.field, p.operator)
}

func (p periodSpecification[T]) GetValues() []any {
	return []any{p.start, p.end}
}

func Between[T any](field string, start T, end T) repository.Specification {
	return periodSpecification[T]{
		field:    field,
		operator: "BETWEEN",
		start:    start,
		end:      end,
	}
}

func JoinAndEqual(joins []string, preloads []string) repository.Specification {
	return binaryOperatorSpecification[string]{
		joinAndPreload: joinAndPreload{
			joins:   joins,
			preload: preloads,
		},
		field:    "",
		value:    "",
		operator: "",
	}
}

func GreaterThan[T comparable](field string, value T) repository.Specification {
	return binaryOperatorSpecification[T]{
		field:    field,
		operator: ">",
		value:    value,
	}
}

func GreaterOrEqual[T comparable](field string, value T) repository.Specification {
	return binaryOperatorSpecification[T]{
		field:    field,
		operator: ">=",
		value:    value,
	}
}

func LessThan[T comparable](field string, value T) repository.Specification {
	return binaryOperatorSpecification[T]{
		field:    field,
		operator: "<",
		value:    value,
	}
}

func LessOrEqual[T comparable](field string, value T) repository.Specification {
	return binaryOperatorSpecification[T]{
		field:    field,
		operator: ">=",
		value:    value,
	}
}

type stringSpecification string

func (s stringSpecification) Sort() []string {
	return nil
}

func (s stringSpecification) Joins() []string {
	return nil
}

func (s stringSpecification) Preload() []string {
	return nil
}

func (s stringSpecification) GetQuery() string {
	return string(s)
}

func (s stringSpecification) GetValues() []any {
	return nil
}

func IsNull(field string) repository.Specification {
	return stringSpecification(fmt.Sprintf("%s IS NULL", field))
}
