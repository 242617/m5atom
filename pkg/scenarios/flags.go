package scenarios

import (
	"time"

	. "github.com/242617/m5atom/pkg/bitmap"
	. "github.com/242617/m5atom/pkg/colors"
	"github.com/242617/m5atom/pkg/motion"
	"github.com/242617/m5atom/pkg/screen"
)

func Flags(screen *screen.Screen) {
	d := 15 * time.Second
	bucket := []struct {
		Bitmap   *Bitmap
		Duration time.Duration
	}{
		{Alfa, d},
		{Bravo, d},
		{Charlie, d},
		{Delta, d},
		{Echo, d},
		{Foxtrot, d},
		{Golf, d},
		{Hotel, d},
		{India, d},
		{Juliett, d},
		{Kilo, d},
		{Lima, d},
		{Mike, d},
		{November, d},
		{Oscar, d},
		{Papa, d},
		{Quebec, d},
		{Romeo, d},
		{Sierra, d},
		{Tango, d},
		{Uniform, d},
		{Victor, d},
		{Whisky, d},
		{Xray, d},
		{Yankee, d},
		{Zulu, d},
	}

	for {
		for _, item := range bucket {
			d5, d7 := item.Duration/5, item.Duration/7
			screen.Add(item.Bitmap)
			w, _ := item.Bitmap.Dimensions()
			motion.New(item.Bitmap, motion.WithEase(motion.EaseInOutSine)).
				Wait(d5+d7).
				To(int(-w+5), 0, d5).
				Wait(d5-d7).
				To(0, 0, d5).
				Wait(d5)
			screen.Remove(item.Bitmap)

			time.Sleep(500 * time.Millisecond)
		}
	}
}

func MustComposeWithSpaces(chars ...*Bitmap) *Bitmap {
	total := make([]*Bitmap, 0, len(chars)+len(chars)-1)
	for i := 0; i < len(chars); i++ {
		total = append(total, chars[i])
		if i != len(chars)-1 {
			total = append(total, Space)
		}
	}
	return MustComposeHorizontal(total...)
}

var (
	AlfaFlag = MustNew5x5(
		White, White, Blue, Blue, Blue,
		White, White, Blue, Blue, Black,
		White, White, Blue, Black, Black,
		White, White, Blue, Blue, Black,
		White, White, Blue, Blue, Blue,
	)
	Alfa      = MustComposeWithSpaces(AlfaFlag, A, L, F, A)
	BravoFlag = MustNew5x5(
		Red, Red, Red, Red, Red,
		Red, Red, Red, Red, Black,
		Red, Red, Red, Black, Black,
		Red, Red, Red, Red, Black,
		Red, Red, Red, Red, Red,
	)
	Bravo       = MustComposeWithSpaces(BravoFlag, B, R, A, V, O)
	CharlieFlag = MustNew5x5(
		Blue, Blue, Blue, Blue, Blue,
		White, White, White, White, White,
		Red, Red, Red, Red, Red,
		White, White, White, White, White,
		Blue, Blue, Blue, Blue, Blue,
	)
	Charlie   = MustComposeWithSpaces(CharlieFlag, C, H, A, R, L, I, E)
	DeltaFlag = MustNew5x5(
		Yellow, Yellow, Yellow, Yellow, Yellow,
		Blue, Blue, Blue, Blue, Blue,
		Blue, Blue, Blue, Blue, Blue,
		Blue, Blue, Blue, Blue, Blue,
		Yellow, Yellow, Yellow, Yellow, Yellow,
	)
	Delta    = MustComposeWithSpaces(DeltaFlag, D, E, L, T, A)
	EchoFlag = MustNew5x5(
		Blue, Blue, Blue, Blue, Blue,
		Blue, Blue, Blue, Blue, Blue,
		Blue, Blue, Blue, Blue, Blue,
		Red, Red, Red, Red, Red,
		Red, Red, Red, Red, Red,
	)
	Echo        = MustComposeWithSpaces(EchoFlag, E, C, H, O)
	FoxtrotFlag = MustNew5x5(
		White, White, Red, White, White,
		White, Red, Red, Red, White,
		Red, Red, Red, Red, Red,
		White, Red, Red, Red, White,
		White, White, Red, White, White,
	)
	Foxtrot  = MustComposeWithSpaces(FoxtrotFlag, F, O, X, T, R, O, T)
	GolfFlag = MustNew5x5(
		Yellow, Blue, Yellow, Blue, Yellow,
		Yellow, Blue, Yellow, Blue, Yellow,
		Yellow, Blue, Yellow, Blue, Yellow,
		Yellow, Blue, Yellow, Blue, Yellow,
		Yellow, Blue, Yellow, Blue, Yellow,
	)
	Golf      = MustComposeWithSpaces(GolfFlag, G, O, L, F)
	HotelFlag = MustNew5x5(
		White, White, Red, Red, Red,
		White, White, Red, Red, Red,
		White, White, Red, Red, Red,
		White, White, Red, Red, Red,
		White, White, Red, Red, Red,
	)
	Hotel     = MustComposeWithSpaces(HotelFlag, H, O, T, E, L)
	IndiaFlag = MustNew5x5(
		Yellow, Yellow, Yellow, Yellow, Yellow,
		Yellow, Yellow, Black, Yellow, Yellow,
		Yellow, Black, Black, Black, Yellow,
		Yellow, Yellow, Black, Yellow, Yellow,
		Yellow, Yellow, Yellow, Yellow, Yellow,
	)
	India       = MustComposeWithSpaces(IndiaFlag, I, N, D, I, A)
	JuliettFlag = MustNew5x5(
		Blue, Blue, Blue, Blue, Blue,
		White, White, White, White, White,
		White, White, White, White, White,
		White, White, White, White, White,
		Blue, Blue, Blue, Blue, Blue,
	)
	Juliett  = MustComposeWithSpaces(JuliettFlag, J, U, L, I, E, T, T)
	KiloFlag = MustNew5x5(
		Yellow, Yellow, Blue, Blue, Blue,
		Yellow, Yellow, Blue, Blue, Blue,
		Yellow, Yellow, Blue, Blue, Blue,
		Yellow, Yellow, Blue, Blue, Blue,
		Yellow, Yellow, Blue, Blue, Blue,
	)
	Kilo     = MustComposeWithSpaces(KiloFlag, K, I, L, O)
	LimaFlag = MustNew5x5(
		Yellow, Yellow, Yellow, Black, Black,
		Yellow, Yellow, Yellow, Black, Black,
		Yellow, Yellow, Yellow, Yellow, Yellow,
		Black, Black, Yellow, Yellow, Yellow,
		Black, Black, Yellow, Yellow, Yellow,
	)
	Lima     = MustComposeWithSpaces(LimaFlag, L, I, M, A)
	MikeFlag = MustNew5x5(
		White, Blue, Blue, Blue, White,
		Blue, White, Blue, White, Blue,
		Blue, Blue, White, Blue, Blue,
		Blue, White, Blue, White, Blue,
		White, Blue, Blue, Blue, White,
	)
	Mike         = MustComposeWithSpaces(MikeFlag, M, I, K, E)
	NovemberFlag = MustNew5x5(
		Blue, White, Blue, White, Black,
		White, Blue, White, Blue, Black,
		Blue, White, Blue, White, Black,
		White, Blue, White, Blue, Black,
		Black, Black, Black, Black, Black,
	)
	November  = MustComposeWithSpaces(NovemberFlag, N, O, V, E, M, B, E, R)
	OscarFlag = MustNew5x5(
		Red, Red, Red, Red, Red,
		Yellow, Red, Red, Red, Red,
		Yellow, Yellow, Red, Red, Red,
		Yellow, Yellow, Yellow, Red, Red,
		Yellow, Yellow, Yellow, Yellow, Red,
	)
	Oscar    = MustComposeWithSpaces(OscarFlag, O, S, C, A, R)
	PapaFlag = MustNew5x5(
		Blue, Blue, Blue, Blue, Blue,
		Blue, White, White, White, Blue,
		Blue, White, White, White, Blue,
		Blue, White, White, White, Blue,
		Blue, Blue, Blue, Blue, Blue,
	)
	Papa       = MustComposeWithSpaces(PapaFlag, P, A, P, A)
	QuebecFlag = MustNew5x5(
		Yellow, Yellow, Yellow, Yellow, Yellow,
		Yellow, Yellow, Yellow, Yellow, Yellow,
		Yellow, Yellow, Yellow, Yellow, Yellow,
		Yellow, Yellow, Yellow, Yellow, Yellow,
		Yellow, Yellow, Yellow, Yellow, Yellow,
	)
	Quebec    = MustComposeWithSpaces(QuebecFlag, Q, U, E, B, E, C)
	RomeoFlag = MustNew5x5(
		Red, Red, Yellow, Red, Red,
		Red, Red, Yellow, Red, Red,
		Yellow, Yellow, Yellow, Yellow, Yellow,
		Red, Red, Yellow, Red, Red,
		Red, Red, Yellow, Red, Red,
	)
	Romeo      = MustComposeWithSpaces(RomeoFlag, R, O, M, E, O)
	SierraFlag = MustNew5x5(
		White, White, White, White, White,
		White, Blue, Blue, Blue, White,
		White, Blue, Blue, Blue, White,
		White, Blue, Blue, Blue, White,
		White, White, White, White, White,
	)
	Sierra    = MustComposeWithSpaces(SierraFlag, S, I, E, R, R, A)
	TangoFlag = MustNew5x5(
		Red, White, White, Blue, Blue,
		Red, White, White, Blue, Blue,
		Red, White, White, Blue, Blue,
		Red, White, White, Blue, Blue,
		Red, White, White, Blue, Blue,
	)
	Tango       = MustComposeWithSpaces(TangoFlag, T, A, N, G, O)
	UniformFlag = MustNew5x5(
		Red, Red, Red, White, White,
		Red, Red, Red, White, White,
		Red, Red, Red, Red, Red,
		White, White, Red, Red, Red,
		White, White, Red, Red, Red,
	)
	Uniform    = MustComposeWithSpaces(UniformFlag, U, N, I, F, O, R, M)
	VictorFlag = MustNew5x5(
		Red, White, White, White, Red,
		White, Red, White, Red, White,
		White, White, Red, White, White,
		White, Red, White, Red, White,
		Red, White, White, White, Red,
	)
	Victor     = MustComposeWithSpaces(VictorFlag, V, I, C, T, O, R)
	WhiskyFlag = MustNew5x5(
		Blue, Blue, Blue, Blue, Blue,
		Blue, White, White, White, Blue,
		Blue, White, Red, White, Blue,
		Blue, White, White, White, Blue,
		Blue, Blue, Blue, Blue, Blue,
	)
	Whisky   = MustComposeWithSpaces(WhiskyFlag, W, H, I, S, K, Y)
	XrayFlag = MustNew5x5(
		White, White, Blue, White, White,
		White, White, Blue, White, White,
		Blue, Blue, Blue, Blue, Blue,
		White, White, Blue, White, White,
		White, White, Blue, White, White,
	)
	Xray       = MustComposeWithSpaces(XrayFlag, X, R, A, Y)
	YankeeFlag = MustNew5x5(
		Yellow, Red, Red, Yellow, Yellow,
		Red, Red, Yellow, Yellow, Red,
		Red, Yellow, Yellow, Red, Red,
		Yellow, Yellow, Red, Red, Yellow,
		Yellow, Red, Red, Yellow, Yellow,
	)
	Yankee   = MustComposeWithSpaces(YankeeFlag, Y, A, N, K, E, E)
	ZuluFlag = MustNew5x5(
		Yellow, Yellow, Yellow, Yellow, Blue,
		Black, Yellow, Yellow, Blue, Blue,
		Black, Black, Black, Blue, Blue,
		Black, Black, Red, Red, Blue,
		Black, Red, Red, Red, Red,
	)
	Zulu = MustComposeWithSpaces(ZuluFlag, Z, U, L, U)
)
