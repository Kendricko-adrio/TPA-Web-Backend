package migration

import (
	"github.com/kendricko-adrio/gqlgen-todos/database"
	"github.com/kendricko-adrio/gqlgen-todos/graph/model"
)

func init(){
	db, err := database.Connect()
	if err != nil{
		panic(err)
	}

	db.Exec("DROP TABLE users_games")
	db.Migrator().DropTable(&model.User{}, &model.Country{}, &model.Game{}, &model.Provider{}, &model.Status{})
	db.AutoMigrate(&model.User{}, &model.Country{}, &model.Game{}, &model.Provider{}, &model.Status{})
}

func SeedAll(){
	SeedStatus()
	SeedCountry()
	SeedGames()
	SeedProvider()
	SeedUser()
}

func SeedUser(){
	db, err := database.Connect()
	if err != nil{
		panic(err)
	}

	pass, _ := model.HashPassword("dummy")
	db.Create(&model.User{
		UserName:   "dummy",
		Password:   pass,
		FirstName:  "dum",
		LastName:   "my",
		AuthToken:  "asdf",
		Email:      "asdf",
		IDToken:    "asdf",
		PhotoURL:   "https://cdn.idntimes.com/content-images/community/2020/07/2f6d6ba28aee33bd220c6d419fd5faee-a238c2af6f14dc0f0653294961526bca_600x400.jpg",
		Provider:   nil,
		ProviderID: 2,
		CountryID:  1,
		StatusID:   2,
		Games: []*model.Game{
			{
				ID:          1,

			},
			{
				ID:          2,
			},
		},
		Otp:        "ABCDE",
	})
}

func SeedStatus(){
	db, err := database.Connect()
	if err != nil{
		panic(err)
	}

	db.Create(&model.Status{
		StatusID: 1,
		StatusName: "Not Auth",
	})
	db.Create(&model.Status{
		StatusID: 2,
		StatusName: "Auth",
	})
}

func SeedProvider(){
	db, err := database.Connect()
	if err != nil{
		panic(err)
	}

	db.Create(&model.Provider{
		Name: "GOOGLE",
	})

	db.Create(&model.Provider{
		Name: "STAEM",
	})
}

func SeedGames(){
	db, err := database.Connect()
	if err != nil{
		panic(err)
	}


	db.Create(&model.Game{
		Name:        "Cyber Punk",
		Description: "Ini game cyber",
		Price:       100000,
		Rating:      2,
		Image:       "https://6.viki.io/image/595814709c7646178ec7352870810506.jpeg?s=900x600&e=t",

	})
	db.Create(&model.Game{
		Name:        "Cyber Punk2",
		Description: "Ini game cyber",
		Price:       100000,
		Rating:      2,
		Image:       "https://6.viki.io/image/595814709c7646178ec7352870810506.jpeg?s=900x600&e=t",

	})

}

func SeedCountry()  {
	db, err := database.Connect()
	if err != nil{
		panic(err)
	}


	db.Create(&model.Country{
		Name:     "Indonesia",
		Code:     "ID",
	})
	db.Create(&model.Country{ Name: "Afghanistan", Code: "AF" })
	db.Create(&model.Country{ Name: "Åland Islands", Code: "AX" })
	db.Create(&model.Country{ Name: "Albania", Code: "AL" })
	db.Create(&model.Country{ Name: "Algeria", Code: "DZ" })
	db.Create(&model.Country{ Name: "American Samoa", Code: "AS" })
	db.Create(&model.Country{ Name: "Andorra", Code: "AD" })
	db.Create(&model.Country{ Name: "Angola", Code: "AO" })
	db.Create(&model.Country{ Name: "Anguilla", Code: "AI" })
	db.Create(&model.Country{ Name: "Antarctica", Code: "AQ" })
	db.Create(&model.Country{ Name: "Antigua and Barbuda", Code: "AG" })
	db.Create(&model.Country{ Name: "Argentina", Code: "AR" })
	db.Create(&model.Country{ Name: "Armenia", Code: "AM" })
	db.Create(&model.Country{ Name: "Aruba", Code: "AW" })
	db.Create(&model.Country{ Name: "Australia", Code: "AU" })
	db.Create(&model.Country{ Name: "Austria", Code: "AT" })
	db.Create(&model.Country{ Name: "Azerbaijan", Code: "AZ" })
	db.Create(&model.Country{ Name: "Bahamas", Code: "BS" })
	db.Create(&model.Country{ Name: "Bahrain", Code: "BH" })
	db.Create(&model.Country{ Name: "Bangladesh", Code: "BD" })
	db.Create(&model.Country{ Name: "Barbados", Code: "BB" })
	db.Create(&model.Country{ Name: "Belarus", Code: "BY" })
	db.Create(&model.Country{ Name: "Belgium", Code: "BE" })
	db.Create(&model.Country{ Name: "Belize", Code: "BZ" })
	db.Create(&model.Country{ Name: "Benin", Code: "BJ" })
	db.Create(&model.Country{ Name: "Bermuda", Code: "BM" })
	db.Create(&model.Country{ Name: "Bhutan", Code: "BT" })
	db.Create(&model.Country{ Name: "Bolivia", Code: "BO" })
	db.Create(&model.Country{ Name: "Bosnia and Herzegovina", Code: "BA" })
	db.Create(&model.Country{ Name: "Botswana", Code: "BW" })
	db.Create(&model.Country{ Name: "Bouvet Island", Code: "BV" })
	db.Create(&model.Country{ Name: "Brazil", Code: "BR" })
	db.Create(&model.Country{ Name: "British Indian Ocean Territory", Code: "IO" })
	db.Create(&model.Country{ Name: "Brunei Darussalam", Code: "BN" })
	db.Create(&model.Country{ Name: "Bulgaria", Code: "BG" })
	db.Create(&model.Country{ Name: "Burkina Faso", Code: "BF" })
	db.Create(&model.Country{ Name: "Burundi", Code: "BI" })
	db.Create(&model.Country{ Name: "Cambodia", Code: "KH" })
	db.Create(&model.Country{ Name: "Cameroon", Code: "CM" })
	db.Create(&model.Country{ Name: "Canada", Code: "CA" })
	db.Create(&model.Country{ Name: "Cape Verde", Code: "CV" })
	db.Create(&model.Country{ Name: "Cayman Islands", Code: "KY" })
	db.Create(&model.Country{ Name: "Central African Republic", Code: "CF" })
	db.Create(&model.Country{ Name: "Chad", Code: "TD" })
	db.Create(&model.Country{ Name: "Chile", Code: "CL" })
	db.Create(&model.Country{ Name: "China", Code: "CN" })
	db.Create(&model.Country{ Name: "Christmas Island", Code: "CX" })
	db.Create(&model.Country{ Name: "Cocos (Keeling) Islands", Code: "CC" })
	db.Create(&model.Country{ Name: "Colombia", Code: "CO" })
	db.Create(&model.Country{ Name: "Comoros", Code: "KM" })
	db.Create(&model.Country{ Name: "Congo", Code: "CG" })
	db.Create(&model.Country{ Name: "Congo, The Democratic Republic of the", Code: "CD" })
	db.Create(&model.Country{ Name: "Cook Islands", Code: "CK" })
	db.Create(&model.Country{ Name: "Costa Rica", Code: "CR" })
	db.Create(&model.Country{ Name: "Cote D'Ivoire", Code: "CI" })
	db.Create(&model.Country{ Name: "Croatia", Code: "HR" })
	db.Create(&model.Country{ Name: "Cuba", Code: "CU" })
	db.Create(&model.Country{ Name: "Cyprus", Code: "CY" })
	db.Create(&model.Country{ Name: "Czech Republic", Code: "CZ" })
	db.Create(&model.Country{ Name: "Denmark", Code: "DK" })
	db.Create(&model.Country{ Name: "Djibouti", Code: "DJ" })
	db.Create(&model.Country{ Name: "Dominica", Code: "DM" })
	db.Create(&model.Country{ Name: "Dominican Republic", Code: "DO" })
	db.Create(&model.Country{ Name: "Ecuador", Code: "EC" })
	db.Create(&model.Country{ Name: "Egypt", Code: "EG" })
	db.Create(&model.Country{ Name: "El Salvador", Code: "SV" })
	db.Create(&model.Country{ Name: "Equatorial Guinea", Code: "GQ" })
	db.Create(&model.Country{ Name: "Eritrea", Code: "ER" })
	db.Create(&model.Country{ Name: "Estonia", Code: "EE" })
	db.Create(&model.Country{ Name: "Ethiopia", Code: "ET" })
	db.Create(&model.Country{ Name: "Falkland Islands (Malvinas)", Code: "FK" })
	db.Create(&model.Country{ Name: "Faroe Islands", Code: "FO" })
	db.Create(&model.Country{ Name: "Fiji", Code: "FJ" })
	db.Create(&model.Country{ Name: "Finland", Code: "FI" })
	db.Create(&model.Country{ Name: "France", Code: "FR" })
	db.Create(&model.Country{ Name: "French Guiana", Code: "GF" })
	db.Create(&model.Country{ Name: "French Polynesia", Code: "PF" })
	db.Create(&model.Country{ Name: "French Southern Territories", Code: "TF" })
	db.Create(&model.Country{ Name: "Gabon", Code: "GA" })
	db.Create(&model.Country{ Name: "Gambia", Code: "GM" })
	db.Create(&model.Country{ Name: "Georgia", Code: "GE" })
	db.Create(&model.Country{ Name: "Germany", Code: "DE" })
	db.Create(&model.Country{ Name: "Ghana", Code: "GH" })
	db.Create(&model.Country{ Name: "Gibraltar", Code: "GI" })
	db.Create(&model.Country{ Name: "Greece", Code: "GR" })
	db.Create(&model.Country{ Name: "Greenland", Code: "GL" })
	db.Create(&model.Country{ Name: "Grenada", Code: "GD" })
	db.Create(&model.Country{ Name: "Guadeloupe", Code: "GP" })
	db.Create(&model.Country{ Name: "Guam", Code: "GU" })
	db.Create(&model.Country{ Name: "Guatemala", Code: "GT" })
	db.Create(&model.Country{ Name: "Guernsey", Code: "GG" })
	db.Create(&model.Country{ Name: "Guinea", Code: "GN" })
	db.Create(&model.Country{ Name: "Guinea-Bissau", Code: "GW" })
	db.Create(&model.Country{ Name: "Guyana", Code: "GY" })
	db.Create(&model.Country{ Name: "Haiti", Code: "HT" })
	db.Create(&model.Country{ Name: "Heard Island and Mcdonald Islands", Code: "HM" })
	db.Create(&model.Country{ Name: "Holy See (Vatican City State)", Code: "VA" })
	db.Create(&model.Country{ Name: "Honduras", Code: "HN" })
	db.Create(&model.Country{ Name: "Hong Kong", Code: "HK" })
	db.Create(&model.Country{ Name: "Hungary", Code: "HU" })
	db.Create(&model.Country{ Name: "Iceland", Code: "IS" })
	db.Create(&model.Country{ Name: "India", Code: "IN" })
	db.Create(&model.Country{ Name: "Indonesia", Code: "ID" })
	db.Create(&model.Country{ Name: "Iran, Islamic Republic Of", Code: "IR" })
	db.Create(&model.Country{ Name: "Iraq", Code: "IQ" })
	db.Create(&model.Country{ Name: "Ireland", Code: "IE" })
	db.Create(&model.Country{ Name: "Isle of Man", Code: "IM" })
	db.Create(&model.Country{ Name: "Israel", Code: "IL" })
	db.Create(&model.Country{ Name: "Italy", Code: "IT" })
	db.Create(&model.Country{ Name: "Jamaica", Code: "JM" })
	db.Create(&model.Country{ Name: "Japan", Code: "JP" })
	db.Create(&model.Country{ Name: "Jersey", Code: "JE" })
	db.Create(&model.Country{ Name: "Jordan", Code: "JO" })
	db.Create(&model.Country{ Name: "Kazakhstan", Code: "KZ" })
	db.Create(&model.Country{ Name: "Kenya", Code: "KE" })
	db.Create(&model.Country{ Name: "Kiribati", Code: "KI" })
	db.Create(&model.Country{ Name: "Democratic People's Republic of Korea", Code: "KP" })
	db.Create(&model.Country{ Name: "Korea, Republic of", Code: "KR" })
	db.Create(&model.Country{ Name: "Kosovo", Code: "XK" })
	db.Create(&model.Country{ Name: "Kuwait", Code: "KW" })
	db.Create(&model.Country{ Name: "Kyrgyzstan", Code: "KG" })
	db.Create(&model.Country{ Name: "Lao People's Democratic Republic", Code: "LA" })
	db.Create(&model.Country{ Name: "Latvia", Code: "LV" })
	db.Create(&model.Country{ Name: "Lebanon", Code: "LB" })
	db.Create(&model.Country{ Name: "Lesotho", Code: "LS" })
	db.Create(&model.Country{ Name: "Liberia", Code: "LR" })
	db.Create(&model.Country{ Name: "Libyan Arab Jamahiriya", Code: "LY" })
	db.Create(&model.Country{ Name: "Liechtenstein", Code: "LI" })
	db.Create(&model.Country{ Name: "Lithuania", Code: "LT" })
	db.Create(&model.Country{ Name: "Luxembourg", Code: "LU" })
	db.Create(&model.Country{ Name: "Macao", Code: "MO" })
	db.Create(&model.Country{ Name: "Macedonia, The Former Yugoslav Republic of", Code: "MK" })
	db.Create(&model.Country{ Name: "Madagascar", Code: "MG" })
	db.Create(&model.Country{ Name: "Malawi", Code: "MW" })
	db.Create(&model.Country{ Name: "Malaysia", Code: "MY" })
	db.Create(&model.Country{ Name: "Maldives", Code: "MV" })
	db.Create(&model.Country{ Name: "Mali", Code: "ML" })
	db.Create(&model.Country{ Name: "Malta", Code: "MT" })
	db.Create(&model.Country{ Name: "Marshall Islands", Code: "MH" })
	db.Create(&model.Country{ Name: "Martinique", Code: "MQ" })
	db.Create(&model.Country{ Name: "Mauritania", Code: "MR" })
	db.Create(&model.Country{ Name: "Mauritius", Code: "MU" })
	db.Create(&model.Country{ Name: "Mayotte", Code: "YT" })
	db.Create(&model.Country{ Name: "Mexico", Code: "MX" })
	db.Create(&model.Country{ Name: "Micronesia, Federated States of", Code: "FM" })
	db.Create(&model.Country{ Name: "Moldova, Republic of", Code: "MD" })
	db.Create(&model.Country{ Name: "Monaco", Code: "MC" })
	db.Create(&model.Country{ Name: "Mongolia", Code: "MN" })
	db.Create(&model.Country{ Name: "Montenegro", Code: "ME" })
	db.Create(&model.Country{ Name: "Montserrat", Code: "MS" })
	db.Create(&model.Country{ Name: "Morocco", Code: "MA" })
	db.Create(&model.Country{ Name: "Mozambique", Code: "MZ" })
	db.Create(&model.Country{ Name: "Myanmar", Code: "MM" })
	db.Create(&model.Country{ Name: "Namibia", Code: "NA" })
	db.Create(&model.Country{ Name: "Nauru", Code: "NR" })
	db.Create(&model.Country{ Name: "Nepal", Code: "NP" })
	db.Create(&model.Country{ Name: "Netherlands", Code: "NL" })
	db.Create(&model.Country{ Name: "Netherlands Antilles", Code: "AN" })
	db.Create(&model.Country{ Name: "New Caledonia", Code: "NC" })
	db.Create(&model.Country{ Name: "New Zealand", Code: "NZ" })
	db.Create(&model.Country{ Name: "Nicaragua", Code: "NI" })
	db.Create(&model.Country{ Name: "Niger", Code: "NE" })
	db.Create(&model.Country{ Name: "Nigeria", Code: "NG" })
	db.Create(&model.Country{ Name: "Niue", Code: "NU" })
	db.Create(&model.Country{ Name: "Norfolk Island", Code: "NF" })
	db.Create(&model.Country{ Name: "Northern Mariana Islands", Code: "MP" })
	db.Create(&model.Country{ Name: "Norway", Code: "NO" })
	db.Create(&model.Country{ Name: "Oman", Code: "OM" })
	db.Create(&model.Country{ Name: "Pakistan", Code: "PK" })
	db.Create(&model.Country{ Name: "Palau", Code: "PW" })
	db.Create(&model.Country{ Name: "Palestinian Territory, Occupied", Code: "PS" })
	db.Create(&model.Country{ Name: "Panama", Code: "PA" })
	db.Create(&model.Country{ Name: "Papua New Guinea", Code: "PG" })
	db.Create(&model.Country{ Name: "Paraguay", Code: "PY" })
	db.Create(&model.Country{ Name: "Peru", Code: "PE" })
	db.Create(&model.Country{ Name: "Philippines", Code: "PH" })
	db.Create(&model.Country{ Name: "Pitcairn", Code: "PN" })
	db.Create(&model.Country{ Name: "Poland", Code: "PL" })
	db.Create(&model.Country{ Name: "Portugal", Code: "PT" })
	db.Create(&model.Country{ Name: "Puerto Rico", Code: "PR" })
	db.Create(&model.Country{ Name: "Qatar", Code: "QA" })
	db.Create(&model.Country{ Name: "Reunion", Code: "RE" })
	db.Create(&model.Country{ Name: "Romania", Code: "RO" })
	db.Create(&model.Country{ Name: "Russian Federation", Code: "RU" })
	db.Create(&model.Country{ Name: "Rwanda", Code: "RW" })
	db.Create(&model.Country{ Name: "Saint Helena", Code: "SH" })
	db.Create(&model.Country{ Name: "Saint Kitts and Nevis", Code: "KN" })
	db.Create(&model.Country{ Name: "Saint Lucia", Code: "LC" })
	db.Create(&model.Country{ Name: "Saint Pierre and Miquelon", Code: "PM" })
	db.Create(&model.Country{ Name: "Saint Vincent and the Grenadines", Code: "VC" })
	db.Create(&model.Country{ Name: "Samoa", Code: "WS" })
	db.Create(&model.Country{ Name: "San Marino", Code: "SM" })
	db.Create(&model.Country{ Name: "Sao Tome and Principe", Code: "ST" })
	db.Create(&model.Country{ Name: "Saudi Arabia", Code: "SA" })
	db.Create(&model.Country{ Name: "Senegal", Code: "SN" })
	db.Create(&model.Country{ Name: "Serbia", Code: "RS" })
	db.Create(&model.Country{ Name: "Seychelles", Code: "SC" })
	db.Create(&model.Country{ Name: "Sierra Leone", Code: "SL" })
	db.Create(&model.Country{ Name: "Singapore", Code: "SG" })
	db.Create(&model.Country{ Name: "Slovakia", Code: "SK" })
	db.Create(&model.Country{ Name: "Slovenia", Code: "SI" })
	db.Create(&model.Country{ Name: "Solomon Islands", Code: "SB" })
	db.Create(&model.Country{ Name: "Somalia", Code: "SO" })
	db.Create(&model.Country{ Name: "South Africa", Code: "ZA" })
	db.Create(&model.Country{ Name: "South Georgia and the South Sandwich Islands", Code: "GS" })
	db.Create(&model.Country{ Name: "Spain", Code: "ES" })
	db.Create(&model.Country{ Name: "Sri Lanka", Code: "LK" })
	db.Create(&model.Country{ Name: "Sudan", Code: "SD" })
	db.Create(&model.Country{ Name: "Suriname", Code: "SR" })
	db.Create(&model.Country{ Name: "Svalbard and Jan Mayen", Code: "SJ" })
	db.Create(&model.Country{ Name: "Swaziland", Code: "SZ" })
	db.Create(&model.Country{ Name: "Sweden", Code: "SE" })
	db.Create(&model.Country{ Name: "Switzerland", Code: "CH" })
	db.Create(&model.Country{ Name: "Syrian Arab Republic", Code: "SY" })
	db.Create(&model.Country{ Name: "Taiwan", Code: "TW" })
	db.Create(&model.Country{ Name: "Tajikistan", Code: "TJ" })
	db.Create(&model.Country{ Name: "Tanzania, United Republic of", Code: "TZ" })
	db.Create(&model.Country{ Name: "Thailand", Code: "TH" })
	db.Create(&model.Country{ Name: "Timor-Leste", Code: "TL" })
	db.Create(&model.Country{ Name: "Togo", Code: "TG" })
	db.Create(&model.Country{ Name: "Tokelau", Code: "TK" })
	db.Create(&model.Country{ Name: "Tonga", Code: "TO" })
	db.Create(&model.Country{ Name: "Trinidad and Tobago", Code: "TT" })
	db.Create(&model.Country{ Name: "Tunisia", Code: "TN" })
	db.Create(&model.Country{ Name: "Turkey", Code: "TR" })
	db.Create(&model.Country{ Name: "Turkmenistan", Code: "TM" })
	db.Create(&model.Country{ Name: "Turks and Caicos Islands", Code: "TC" })
	db.Create(&model.Country{ Name: "Tuvalu", Code: "TV" })
	db.Create(&model.Country{ Name: "Uganda", Code: "UG" })
	db.Create(&model.Country{ Name: "Ukraine", Code: "UA" })
	db.Create(&model.Country{ Name: "United Arab Emirates", Code: "AE" })
	db.Create(&model.Country{ Name: "United Kingdom", Code: "GB" })
	db.Create(&model.Country{ Name: "United States", Code: "US" })
	db.Create(&model.Country{ Name: "United States Minor Outlying Islands", Code: "UM" })
	db.Create(&model.Country{ Name: "Uruguay", Code: "UY" })
	db.Create(&model.Country{ Name: "Uzbekistan", Code: "UZ" })
	db.Create(&model.Country{ Name: "Vanuatu", Code: "VU" })
	db.Create(&model.Country{ Name: "Venezuela", Code: "VE" })
	db.Create(&model.Country{ Name: "Viet Nam", Code: "VN" })
	db.Create(&model.Country{ Name: "Virgin Islands, British", Code: "VG" })
	db.Create(&model.Country{ Name: "Virgin Islands, U.S.", Code: "VI" })
	db.Create(&model.Country{ Name: "Wallis and Futuna", Code: "WF" })
	db.Create(&model.Country{ Name: "Western Sahara", Code: "EH" })
	db.Create(&model.Country{ Name: "Yemen", Code: "YE" })
	db.Create(&model.Country{ Name: "Zambia", Code: "ZM" })
	db.Create(&model.Country{ Name: "Zimbabwe", Code: "ZW" })
}