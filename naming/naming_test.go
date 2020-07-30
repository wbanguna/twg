package naming

import "testing"

func TestDog(t *testing.T) {

}

func TestDog_Bark_muzzled(t *testing.T) {

}

func TestDog_Bark_unmuzzled(t *testing.T) {

}

func TestSpeak(t *testing.T) {

}

func TestColour(t *testing.T) {
	//hex := Colour("blue")
	arg := "blue"
	want := "#0000FF"
	got := Colour(arg)
	//if hex != "#0000FE" {
	if got != want {
		//t.Errorf("Colour(%s) = %s; want %s", "blue", hex, "#0000FE")
		t.Errorf("Colour(%q) = %s; want %s", arg, got, want)
	}
}
