// Leap stub file

// The package name is expected by the test program.
package leap

// testVersion should match the targetTestVersion in the test file.
const testVersion = 3

/**
on every year that is evenly divisible by 4
  except every year that is evenly divisible by 100
    unless the year is also evenly divisible by 400
4年一闰，100年不闰，400年再闰
 */
func IsLeapYear(year int) bool {
	// Write some code here to pass the test suite.
	if year%100 == 0 {
		if year%400 == 0 {
			return true
		}
		return false
	} else {
		if year%4 == 0 {
			return true
		}
	}
	return false
}
