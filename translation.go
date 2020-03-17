package uadmin

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"reflect"
	"strings"

	"github.com/uadmin/uadmin/colors"
)

// Translation is for multilingual fields
type translation struct {
	Name     string
	Code     string
	Flag     string
	Default  bool
	Active   bool
	Value    string
	NewValue string
	OldValue string
	Changed  bool
}

// InitalizeLanguage !
func initializeLanguage() {
	// Load Multilanguage
	// multiLanguage, err := ioutil.ReadFile("./templates/uadmin/multilingual.json")

	// if err != nil {
	// 	multiLanguage = []byte{}
	// }
	//
	// err = json.Unmarshal(multiLanguage, &langMap)

	// if err != nil {
	// 	Trail(ERROR, "uadmin.initializeLanguage.json.Unmarshal %s", err.Error())
	// }

	langList := []Language{}
	if Count(&langList, "") != 0 {
		// Setup Active languages
		activeLangs = []Language{}
		Filter(&activeLangs, "active = ?", true)

		// Setup default language
		Get(&defaultLang, "`default` = ?", true)
		return
	}

	langs := [][]string{
		{"Abkhaz", "аҧсуа бызшәа, аҧсшәа", "ab"},
		{"Afar", "Afaraf", "aa"},
		{"Afrikaans", "Afrikaans", "af"},
		{"Akan", "Akan", "ak"},
		{"Albanian", "Shqip", "sq"},
		{"Amharic", "አማርኛ", "am"},
		{"Arabic", "العربية", "ar"},
		{"Aragonese", "aragonés", "an"},
		{"Armenian", "Հայերեն", "hy"},
		{"Assamese", "অসমীয়া", "as"},
		{"Avaric", "авар мацӀ, магӀарул мацӀ", "av"},
		{"Avestan", "avesta", "ae"},
		{"Aymara", "aymar aru", "ay"},
		{"Azerbaijani", "azərbaycan dili", "az"},
		{"Bambara", "bamanankan", "bm"},
		{"Bashkir", "башҡорт теле", "ba"},
		{"Basque", "euskara, euskera", "eu"},
		{"Belarusian", "беларуская мова", "be"},
		{"Bengali, Bangla", "বাংলা", "bn"},
		{"Bihari", "भोजपुरी", "bh"},
		{"Bislama", "Bislama", "bi"},
		{"Bosnian", "bosanski jezik", "bs"},
		{"Breton", "brezhoneg", "br"},
		{"Bulgarian", "български език", "bg"},
		{"Burmese", "ဗမာစာ", "my"},
		{"Catalan", "català", "ca"},
		{"Chamorro", "Chamoru", "ch"},
		{"Chechen", "нохчийн мотт", "ce"},
		{"Chichewa, Chewa, Nyanja", "chiCheŵa, chinyanja", "ny"},
		{"Chinese", "中文 (Zhōngwén), 汉语, 漢語", "zh"},
		{"Chuvash", "чӑваш чӗлхи", "cv"},
		{"Cornish", "Kernewek", "kw"},
		{"Corsican", "corsu, lingua corsa", "co"},
		{"Cree", "ᓀᐦᐃᔭᐍᐏᐣ", "cr"},
		{"Croatian", "hrvatski jezik", "hr"},
		{"Czech", "čeština, český jazyk", "cs"},
		{"Danish", "dansk", "da"},
		{"Divehi, Dhivehi, Maldivian", "ދިވެހި", "dv"},
		{"Dutch", "Nederlands, Vlaams", "nl"},
		{"Dzongkha", "རྫོང་ཁ", "dz"},
		{"English", "English", "en"},
		{"Esperanto", "Esperanto", "eo"},
		{"Estonian", "eesti, eesti keel", "et"},
		{"Ewe", "Eʋegbe", "ee"},
		{"Faroese", "føroyskt", "fo"},
		{"Fijian", "vosa Vakaviti", "fj"},
		{"Filipino", "Filipino", "fl"},
		{"Finnish", "suomi, suomen kieli", "fi"},
		{"French", "français, langue française", "fr"},
		{"Fula, Fulah, Pulaar, Pular", "Fulfulde, Pulaar, Pular", "ff"},
		{"Galician", "galego", "gl"},
		{"Georgian", "ქართული", "ka"},
		{"German", "Deutsch", "de"},
		{"Greek (modern)", "ελληνικά", "el"},
		{"Guaraní", "Avañe'ẽ", "gn"},
		{"Gujarati", "ગુજરાતી", "gu"},
		{"Haitian, Haitian Creole", "Kreyòl ayisyen", "ht"},
		{"Hausa", "(Hausa) هَوُسَ", "ha"},
		{"Hebrew (modern)", "עברית", "he"},
		{"Herero", "Otjiherero", "hz"},
		{"Hindi", "हिन्दी, हिंदी", "hi"},
		{"Hiri Motu", "Hiri Motu", "ho"},
		{"Hungarian", "magyar", "hu"},
		{"Interlingua", "Interlingua", "ia"},
		{"Indonesian", "Bahasa Indonesia", "id"},
		{"Interlingue", "Originally called Occidental; then Interlingue after WWII", "ie"},
		{"Irish", "Gaeilge", "ga"},
		{"Igbo", "Asụsụ Igbo", "ig"},
		{"Inupiaq", "Iñupiaq, Iñupiatun", "ik"},
		{"Ido", "Ido", "io"},
		{"Icelandic", "Íslenska", "is"},
		{"Italian", "Italiano", "it"},
		{"Inuktitut", "ᐃᓄᒃᑎᑐᑦ", "iu"},
		{"Japanese", "日本語 (にほんご)", "ja"},
		{"Javanese", "ꦧꦱꦗꦮ, Basa Jawa", "jv"},
		{"Kalaallisut, Greenlandic", "kalaallisut, kalaallit oqaasii", "kl"},
		{"Kannada", "ಕನ್ನಡ", "kn"},
		{"Kanuri", "Kanuri", "kr"},
		{"Kashmiri", "कश्मीरी, كشميري‎", "ks"},
		{"Kazakh", "қазақ тілі", "kk"},
		{"Khmer", "ខ្មែរ, ខេមរភាសា, ភាសាខ្មែរ", "km"},
		{"Kikuyu, Gikuyu", "Gĩkũyũ", "ki"},
		{"Kinyarwanda", "Ikinyarwanda", "rw"},
		{"Kyrgyz", "Кыргызча, Кыргыз тили", "ky"},
		{"Komi", "коми кыв", "kv"},
		{"Kongo", "Kikongo", "kg"},
		{"Korean", "한국어", "ko"},
		{"Kurdish", "Kurdî, كوردی‎", "ku"},
		{"Kwanyama, Kuanyama", "Kuanyama", "kj"},
		{"Latin", "latine, lingua latina", "la"},
		{"Luxembourgish, Letzeburgesch", "Lëtzebuergesch", "lb"},
		{"Ganda", "Luganda", "lg"},
		{"Limburgish, Limburgan, Limburger", "Limburgs", "li"},
		{"Lingala", "Lingála", "ln"},
		{"Lao", "ພາສາລາວ", "lo"},
		{"Lithuanian", "lietuvių kalba", "lt"},
		{"Luba-Katanga", "Tshiluba", "lu"},
		{"Latvian", "latviešu valoda", "lv"},
		{"Manx", "Gaelg, Gailck", "gv"},
		{"Macedonian", "македонски јазик", "mk"},
		{"Malagasy", "fiteny malagasy", "mg"},
		{"Malay", "bahasa Melayu, بهاس ملايو‎", "ms"},
		{"Malayalam", "മലയാളം", "ml"},
		{"Maltese", "Malti", "mt"},
		{"Māori", "te reo Māori", "mi"},
		{"Marathi (Marāṭhī)", "मराठी", "mr"},
		{"Marshallese", "Kajin M̧ajeļ", "mh"},
		{"Mongolian", "Монгол хэл", "mn"},
		{"Nauruan", "Dorerin Naoero", "na"},
		{"Navajo, Navaho", "Diné bizaad", "nv"},
		{"Northern Ndebele", "isiNdebele", "nd"},
		{"Nepali", "नेपाली", "ne"},
		{"Ndonga", "Owambo", "ng"},
		{"Norwegian Bokmål", "Norsk bokmål", "nb"},
		{"Norwegian Nynorsk", "Norsk nynorsk", "nn"},
		{"Norwegian", "Norsk", "no"},
		{"Nuosu", "ꆈꌠ꒿ Nuosuhxop", "ii"},
		{"Southern Ndebele", "isiNdebele", "nr"},
		{"Occitan", "occitan, lenga d'òc", "oc"},
		{"Ojibwe, Ojibwa", "ᐊᓂᔑᓈᐯᒧᐎᓐ", "oj"},
		{"Old Church Slavonic, Church Slavonic, Old Bulgarian", "ѩзыкъ словѣньскъ", "cu"},
		{"Oromo", "Afaan Oromoo", "om"},
		{"Oriya", "ଓଡ଼ିଆ", "or"},
		{"Ossetian, Ossetic", "ирон æвзаг", "os"},
		{"(Eastern) Punjabi", "ਪੰਜਾਬੀ", "pa"},
		{"Pāli", "पाऴि", "pi"},
		{"Persian (Farsi)", "فارسی", "fa"},
		{"Polish", "język polski, polszczyzna", "pl"},
		{"Pashto, Pushto", "پښتو", "ps"},
		{"Portuguese", "Português", "pt"},
		{"Quechua", "Runa Simi, Kichwa", "qu"},
		{"Romansh", "rumantsch grischun", "rm"},
		{"Kirundi", "Ikirundi", "rn"},
		{"Romanian", "Română", "ro"},
		{"Russian", "Русский", "ru"},
		{"Sanskrit (Saṁskṛta)", "संस्कृतम्", "sa"},
		{"Sardinian", "sardu", "sc"},
		{"Sindhi", "सिन्धी, سنڌي، سندھی‎", "sd"},
		{"Northern Sami", "Davvisámegiella", "se"},
		{"Samoan", "gagana fa'a Samoa", "sm"},
		{"Sango", "yângâ tî sängö", "sg"},
		{"Serbian", "српски језик", "sr"},
		{"Scottish Gaelic, Gaelic", "Gàidhlig", "gd"},
		{"Shona", "chiShona", "sn"},
		{"Sinhalese, Sinhala", "සිංහල", "si"},
		{"Slovak", "slovenčina, slovenský jazyk", "sk"},
		{"Slovene", "slovenski jezik, slovenščina", "sl"},
		{"Somali", "Soomaaliga, af Soomaali", "so"},
		{"Southern Sotho", "Sesotho", "st"},
		{"Spanish", "Español", "es"},
		{"Sundanese", "Basa Sunda", "su"},
		{"Swahili", "Kiswahili", "sw"},
		{"Swati", "SiSwati", "ss"},
		{"Swedish", "svenska", "sv"},
		{"Tamil", "தமிழ்", "ta"},
		{"Telugu", "తెలుగు", "te"},
		{"Tajik", "тоҷикӣ, toçikī, تاجیکی‎", "tg"},
		{"Thai", "ไทย", "th"},
		{"Tigrinya", "ትግርኛ", "ti"},
		{"Tibetan Standard, Tibetan, Central", "བོད་ཡིག", "bo"},
		{"Turkmen", "Türkmen, Түркмен", "tk"},
		{"Tagalog", "Wikang Tagalog", "tl"},
		{"Tswana", "Setswana", "tn"},
		{"Tonga (Tonga Islands)", "faka Tonga", "to"},
		{"Turkish", "Türkçe", "tr"},
		{"Tsonga", "Xitsonga", "ts"},
		{"Tatar", "татар теле, tatar tele", "tt"},
		{"Twi", "Twi", "tw"},
		{"Tahitian", "Reo Tahiti", "ty"},
		{"Uyghur", "ئۇيغۇرچە‎, Uyghurche", "ug"},
		{"Ukrainian", "Українська", "uk"},
		{"Urdu", "اردو", "ur"},
		{"Uzbek", "Oʻzbek, Ўзбек, أۇزبېك‎", "uz"},
		{"Venda", "Tshivenḓa", "ve"},
		{"Vietnamese", "Tiếng Việt", "vi"},
		{"Volapük", "Volapük", "vo"},
		{"Walloon", "walon", "wa"},
		{"Welsh", "Cymraeg", "cy"},
		{"Wolof", "Wollof", "wo"},
		{"Western Frisian", "Frysk", "fy"},
		{"Xhosa", "isiXhosa", "xh"},
		{"Yiddish", "ייִדיש", "yi"},
		{"Yoruba", "Yorùbá", "yo"},
		{"Zhuang, Chuang", "Saɯ cueŋƅ, Saw cuengh", "za"},
		{"Zulu", "isiZulu", "zu"},
	}
	activeLangs = []Language{}
	tx := db.Begin()
	for i, lang := range langs {
		l := Language{
			EnglishName: lang[0],
			Name:        lang[1],
			Code:        lang[2],
			Active:      false,
		}
		if l.Code == "en" {
			l.AvailableInGui = true
			l.Active = true
			l.Default = true
		}
		tx.Create(&l)

		if l.Active {
			activeLangs = append(activeLangs, l)
		}
		if l.Default {
			defaultLang = l
		}
		Trail(WORKING, "Initializing Languages: [%s%d/%d%s]", colors.FGGreenB, i+1, len(langs), colors.FGNormal)
	}
	tx.Commit()
	Trail(OK, "Initializing Languages: [%s%d/%d%s]", colors.FGGreenB, len(langs), len(langs), colors.FGNormal)
}

// Translate is used to get a translation from a multilingual fields
func Translate(raw string, lang string, args ...bool) string {
	var langParser map[string]json.RawMessage
	err := json.Unmarshal([]byte(raw), &langParser)
	if err != nil {
		return raw
	}
	transtedStr := string(langParser[lang])

	if len(transtedStr) > 2 {
		return transtedStr[1 : len(transtedStr)-1]
	}
	if len(args) > 0 && !args[0] {
		return ""
	}

	Get(&defaultLang, "`default` = ?", true)
	transtedStr = string(langParser[defaultLang.Code])

	if len(transtedStr) > 2 {
		return transtedStr[1 : len(transtedStr)-1]
	}
	return ""
}

// Tf is a function for translating strings into any given language
// Parameters:
// ===========
//   path (string): This is where to get the translation from. It is in the
//                  format of "GROUPNAME/FILENAME" for example: "uadmin/system"
//   lang (string): Is the language code. If empty string is passed we will use
//                  the default language.
//   term (string): The term to translate.
//   args (...interface{}): Is a list of args to fill the term with place holders
func Tf(path string, lang string, term string, args ...interface{}) string {
	var err error
	var buf []byte
	if lang == "" {
		lang = defaultLang.Code
	}

	// Check if the path if for an existing model schema
	pathParts := strings.Split(path, "/")
	isSchemaFile := false
	if len(pathParts) > 2 {
		path = strings.Join(pathParts[0:2], "/")
		isSchemaFile = true
	}

	if langMapCache == nil {
		langMapCache = map[string][]byte{}
	}

	// Check if the translation is cached
	fileName := "./static/i18n/" + strings.ToLower(path) + "." + lang + ".json"
	var ok bool
	if CacheTranslation {
		buf, ok = langMapCache[fileName]
		if !ok {
			if _, err = os.Stat(fileName); os.IsNotExist(err) {
				Trail(WARNING, "Unrecognized path (%s) - fileName:%s", path, fileName)
				return term
			}
			buf, err = ioutil.ReadFile(fileName)
			if err != nil {
				Trail(ERROR, "Unable to read language file (%s)", fileName)
				return term
			}
			langMapCache[fileName] = buf
		}
	} else {
		if _, err = os.Stat(fileName); os.IsNotExist(err) {
			Trail(WARNING, "Unrecognized path (%s) - fileName:%s", path, fileName)
			return term
		}
		buf, err = ioutil.ReadFile(fileName)
		if err != nil {
			Trail(ERROR, "Unable to read language file (%s)", fileName)
			return term
		}
	}

	// Check if it is a schema file or custom files
	if !isSchemaFile {
		langMap := map[string]string{}
		err = json.Unmarshal(buf, &langMap)
		if err != nil {
			Trail(ERROR, "Unknown translation file format (%s)", path)
			return term
		}

		// If the term exists, then return it
		if val, ok := langMap[term]; ok {
			return strings.TrimPrefix(val, translateMe)
		}

		// If it doesn't exist then add it to the file
		if lang != "en" {
			Tf(path, "en", term, args...)
			langMap[term] = translateMe + term
		} else {
			langMap[term] = term
		}

		// Save the file with the new term
		saveLangFile(langMap, fileName)
	} else {
		langMap := structLanguage{}
		err = json.Unmarshal(buf, &langMap)
		if err != nil {
			Trail(ERROR, "Unknown translation file format (%s)", path)
			return term
		}

		// If the term exists, then return it
		// First: ErrMsg
		if strings.ToLower(pathParts[3]) == "errmsg" {
			if val, ok := langMap.Fields[pathParts[2]].ErrMsg[term]; ok {
				return strings.TrimPrefix(val, translateMe)
			}

			// If it doesn't exist then add it to the file
			if lang != "en" {
				Tf(strings.Join(pathParts, "/"), "en", term, args...)
				langMap.Fields[pathParts[2]].ErrMsg[term] = translateMe + term
			} else {
				langMap.Fields[pathParts[2]].ErrMsg[term] = term
			}

			// Save the file with the new term
			saveLangFile(langMap, fileName)
		}
		// TODO: add other parts of the structLang in here
	}
	return term
}

// Translate Model
func translateSchema(s *ModelSchema, lang string) {
	if lang == "" {
		lang = "en"
	}

	pkgName := fmt.Sprint(reflect.TypeOf(models[s.ModelName]))
	pkgName = strings.Split(pkgName, ".")[0]
	fileName := "./static/i18n/" + pkgName + "/" + s.ModelName + "." + lang + ".json"

	// If the model language file doessn't exist, then return
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		return
	}
	buf, err := ioutil.ReadFile(fileName)
	if err != nil {
		Trail(ERROR, "Unable to read language file (%s)", fileName)
		return
	}
	// Parse the file and traslate the schema
	structLang := structLanguage{}
	err = json.Unmarshal(buf, &structLang)
	if err != nil {
		Trail(ERROR, "Invalid format for language file (%s)", fileName)
		return
	}
	s.DisplayName = strings.TrimPrefix(structLang.DisplayName, translateMe)
	for i, f := range s.Fields {
		f.DisplayName = strings.TrimPrefix(structLang.Fields[f.Name].DisplayName, translateMe)
		f.Help = strings.TrimPrefix(structLang.Fields[f.Name].Help, translateMe)
		f.PatternMsg = strings.TrimPrefix(structLang.Fields[f.Name].PatternMsg, translateMe)
		if _, ok := structLang.Fields[f.Name].ErrMsg[f.ErrMsg]; ok {
			f.ErrMsg = strings.TrimPrefix(structLang.Fields[f.Name].ErrMsg[f.ErrMsg], translateMe)
		} else {

		}
		for k, v := range structLang.Fields[f.Name].Choices {
			for index := range f.Choices {
				if f.Choices[index].K == uint(k) {
					f.Choices[index].V = v
					break
				}
			}
		}

		s.Fields[i] = f
	}
}

func getLanguage(r *http.Request) Language {
	langCookie, err := r.Cookie("language")

	if err != nil || langCookie == nil {
		return defaultLang
	}

	for _, l := range activeLangs {
		if l.Code == langCookie.Value {
			return l
		}
	}
	return defaultLang
}
