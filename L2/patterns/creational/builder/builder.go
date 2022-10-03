package main

import (
	"fmt"
	"strings"
)

type SQLQueryBuilderInterface interface {
	Select(table string, fields []string) SQLQueryBuilderInterface
	Where(field string, operator string, value string) SQLQueryBuilderInterface
	GroupBy(fields []string) SQLQueryBuilderInterface
	OrderBy(fields []string) SQLQueryBuilderInterface
	Build() string
}

func NewSqlBuilder() SQLQueryBuilderInterface {
	return &SQLQueryBuilder{}
}

type SQLQueryBuilder struct {
	Query strings.Builder
}

func (b *SQLQueryBuilder) Select(table string, fields []string) SQLQueryBuilderInterface {
	b.Query.WriteString(fmt.Sprintf("SELECT %s FROM %s ", strings.Join(fields, ", "), table))
	return b
}

func (b *SQLQueryBuilder) Where(field string, operator string, value string) SQLQueryBuilderInterface {
	b.Query.WriteString(fmt.Sprintf("WHERE %s %s %s ", field, operator, value))
	return b
}

func (b *SQLQueryBuilder) GroupBy(fields []string) SQLQueryBuilderInterface {
	b.Query.WriteString(fmt.Sprintf("GROUP BY %s ", strings.Join(fields, ", ")))
	return b
}

func (b *SQLQueryBuilder) OrderBy(fields []string) SQLQueryBuilderInterface {
	b.Query.WriteString(fmt.Sprintf("ORDER BY %s ", strings.Join(fields, ", ")))
	return b
}

func (b *SQLQueryBuilder) Build() string {
	return b.Query.String()
}

func main() {
	builder := NewSqlBuilder()
	query := builder.Select("items", []string{"id", "products", "city"}).GroupBy([]string{"city"}).Build()
	fmt.Println(query)
}

/*
Строитель — это порождающий паттерн проектирования, который позволяет создавать объекты пошагово.
+ Подходит для создания объектов, требующих поэтапного создания.
*/
