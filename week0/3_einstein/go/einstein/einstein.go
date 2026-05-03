package einstein

// Speed of light (m/s2)
// C = 300000000;

// Speed of light squared (C * C)
const C2 = 90000000000000000

func Einstein(mass uint) uint64 {
	return uint64(mass) * C2
}