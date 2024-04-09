package cmd

import "fmt"

func Execute() {
	selectedCountry := Chooser(CountryStore.Text...)
	selectedLanguage := Chooser(LangStore.Text...)
	fmt.Println(selectedCountry)
	fmt.Println(selectedLanguage)

	radioBrowser := "https://de1.api.radio-browser.info/json/stations/search?" + "countrycode=" + CountryStore.Mapper[selectedCountry] + "&language=" + LangStore.Mapper[selectedLanguage] + "&hidebroken=true&order=clickcount&reverse=true"

	fmt.Println(radioBrowser)
}
