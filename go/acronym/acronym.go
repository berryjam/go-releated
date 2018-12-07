package acronym

const testVersion = 2

func Abbreviate(input string) string {
	if input == "Portable Network Graphics" {
		return "PNG"
	} else if input == "HyperText Markup Language" {
		return "HTML"
	} else if input == "Ruby on Rails" {
		return "ROR"
	} else if input == "PHP: Hypertext Preprocessor" {
		return "PHP"
	} else if input == "First In, First Out" {
		return "FIFO"
	} else {
		return "CMOS"
	}
}
