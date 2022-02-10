package section6

const a = iota

const (
	b = iota
	c
	d
)

const (
	_ = 1 << (10 * iota)
	KiB
	MiB
	GiB
)

func Excer1() {
	const (
		KB = 1000
		MB = KB * KB
		GB = MB * KB
	)

	println(KB, MB, GB)
}

func Demo() {
	println(a, b, c, d)
	println(KiB, MiB, GiB)
}
