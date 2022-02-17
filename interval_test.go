package interval

import (
	"testing"
	"time"

	avl "github.com/emirpasic/gods/trees/avltree"
	"github.com/stretchr/testify/assert"
)

func TestComparator(t *testing.T) {

	now := time.Now()

	a := Interval{Start: now, End: now.Add(2 * time.Second)}

	b := Interval{Start: now.Add(3 * time.Second), End: now.Add(4 * time.Second)}

	assert.Equal(t, -1, Comparator(a, b))

	assert.Equal(t, 1, Comparator(b, a))

	assert.Equal(t, 0, Comparator(a, a))

	// overlap partially with a
	c := Interval{Start: now.Add(time.Second), End: now.Add(3 * time.Second)}

	assert.Equal(t, 0, Comparator(a, c))

}

func TestAVL(t *testing.T) {

	at := avl.NewWith(Comparator)

	now := time.Now()
	a := Interval{Start: now, End: now.Add(2 * time.Second)}
	b := Interval{Start: now.Add(3 * time.Second), End: now.Add(4 * time.Second)}

	at.Put(a, "x")
	at.Put(b, "y")
	v := at.Values()
	assert.Equal(t, 2, at.Size())
	assert.Equal(t, "x", v[0])
	assert.Equal(t, "y", v[1])

	// overlap partially with a -> should replace a
	// we don't want this behaviour in a booking system, but it is what
	// this implementation does ....
	c := Interval{Start: now.Add(time.Second), End: now.Add(3 * time.Second)}
	at.Put(c, "z")
	assert.Equal(t, 2, at.Size())
	v = at.Values()

	assert.Equal(t, "z", v[0])
	assert.Equal(t, "y", v[1])
}
