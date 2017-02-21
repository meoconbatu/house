package house

const testVersion = 1

var rhymes = []string{
	" the malt\nthat lay in",
	" the rat\nthat ate",
	" the cat\nthat killed",
	" the dog\nthat worried",
	" the cow with the crumpled horn\nthat tossed",
	" the maiden all forlorn\nthat milked",
	" the man all tattered and torn\nthat kissed",
	" the priest all shaven and shorn\nthat married",
	" the rooster that crowed in the morn\nthat woke",
	" the farmer sowing his corn\nthat kept",
	" the horse and the hound and the horn\nthat belonged to",
}

func Song() string {
	var result string
	for i := 0; i <= len(rhymes); i++ {
		result += Verse(i + 1)
		if i < len(rhymes) {
			result += "\n\n"
		}
	}
	return result
}
func Verse(n int) string {
	return "This is" + Rhyme(n) + " the house that Jack built."
}

func Rhyme(n int) string {
	if n == 1 {
		return ""
	}
	return rhymes[n-2] + Rhyme(n-1)
}
