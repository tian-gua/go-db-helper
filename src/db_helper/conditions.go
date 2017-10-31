package dh

type Conditions struct {
	condition string
	args      []interface{}
}

func (c *Conditions) AndEqual(s string, arg interface{}) *Conditions {
	c.condition = c.condition + " AND " + s + " = ?"
	c.args = append(c.args, arg)
	return c
}

func (c *Conditions) AndIsNotNull(s string) *Conditions {
	c.condition = c.condition + " AND " + s + " IS NOT NULL"
	return c
}

func (c *Conditions) AndIsNull(s string) *Conditions {
	c.condition = c.condition + " AND " + s + " IS NULL"
	return c
}

func (c *Conditions) Select(is interface{}) error {
	s, args, err := generateQuerySql(is, c)
	if err != nil {
		return err
	}
	return query(is, s, args)
}

func (c *Conditions) get() (string, []interface{}) {
	return c.condition, c.args
}

type where struct {
	conditions map[uint64]*Conditions
}

func (w *where) has() bool {
	_, ok := w.conditions[getGID()]
	return ok
}

func (w *where) get() *Conditions {
	return w.conditions[getGID()]
}

var whereManager *where
