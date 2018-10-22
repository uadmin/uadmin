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

var langMap map[string]map[string]string

// Translation is for multilingual fields
type Translation struct {
	Name    string
	Code    string
	Flag    string
	Default bool
	Active  bool
	Value   string
}

// InitalizeLanguage !
func initializeLanguage() {
	// Load Multilanguage
	multiLanguage, err := ioutil.ReadFile("./templates/uadmin/multilingual.json")

	if err != nil {
		multiLanguage = []byte{}
	}

	err = json.Unmarshal(multiLanguage, &langMap)

	if err != nil {
		Trail(ERROR, "uadmin.initializeLanguage.json.Unmarshal %s", err.Error())
	}

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
		[]string{"Abkhaz", "аҧсуа бызшәа, аҧсшәа", "ab"},
		[]string{"Afar", "Afaraf", "aa"},
		[]string{"Afrikaans", "Afrikaans", "af"},
		[]string{"Akan", "Akan", "ak"},
		[]string{"Albanian", "Shqip", "sq"},
		[]string{"Amharic", "አማርኛ", "am"},
		[]string{"Arabic", "العربية", "ar"},
		[]string{"Aragonese", "aragonés", "an"},
		[]string{"Armenian", "Հայերեն", "hy"},
		[]string{"Assamese", "অসমীয়া", "as"},
		[]string{"Avaric", "авар мацӀ, магӀарул мацӀ", "av"},
		[]string{"Avestan", "avesta", "ae"},
		[]string{"Aymara", "aymar aru", "ay"},
		[]string{"Azerbaijani", "azərbaycan dili", "az"},
		[]string{"Bambara", "bamanankan", "bm"},
		[]string{"Bashkir", "башҡорт теле", "ba"},
		[]string{"Basque", "euskara, euskera", "eu"},
		[]string{"Belarusian", "беларуская мова", "be"},
		[]string{"Bengali, Bangla", "বাংলা", "bn"},
		[]string{"Bihari", "भोजपुरी", "bh"},
		[]string{"Bislama", "Bislama", "bi"},
		[]string{"Bosnian", "bosanski jezik", "bs"},
		[]string{"Breton", "brezhoneg", "br"},
		[]string{"Bulgarian", "български език", "bg"},
		[]string{"Burmese", "ဗမာစာ", "my"},
		[]string{"Catalan", "català", "ca"},
		[]string{"Chamorro", "Chamoru", "ch"},
		[]string{"Chechen", "нохчийн мотт", "ce"},
		[]string{"Chichewa, Chewa, Nyanja", "chiCheŵa, chinyanja", "ny"},
		[]string{"Chinese", "中文 (Zhōngwén), 汉语, 漢語", "zh"},
		[]string{"Chuvash", "чӑваш чӗлхи", "cv"},
		[]string{"Cornish", "Kernewek", "kw"},
		[]string{"Corsican", "corsu, lingua corsa", "co"},
		[]string{"Cree", "ᓀᐦᐃᔭᐍᐏᐣ", "cr"},
		[]string{"Croatian", "hrvatski jezik", "hr"},
		[]string{"Czech", "čeština, český jazyk", "cs"},
		[]string{"Danish", "dansk", "da"},
		[]string{"Divehi, Dhivehi, Maldivian", "ދިވެހި", "dv"},
		[]string{"Dutch", "Nederlands, Vlaams", "nl"},
		[]string{"Dzongkha", "རྫོང་ཁ", "dz"},
		[]string{"English", "English", "en"},
		[]string{"Esperanto", "Esperanto", "eo"},
		[]string{"Estonian", "eesti, eesti keel", "et"},
		[]string{"Ewe", "Eʋegbe", "ee"},
		[]string{"Faroese", "føroyskt", "fo"},
		[]string{"Fijian", "vosa Vakaviti", "fj"},
		[]string{"Filipino", "Filipino", "fl"},
		[]string{"Finnish", "suomi, suomen kieli", "fi"},
		[]string{"French", "français, langue française", "fr"},
		[]string{"Fula, Fulah, Pulaar, Pular", "Fulfulde, Pulaar, Pular", "ff"},
		[]string{"Galician", "galego", "gl"},
		[]string{"Georgian", "ქართული", "ka"},
		[]string{"German", "Deutsch", "de"},
		[]string{"Greek (modern)", "ελληνικά", "el"},
		[]string{"Guaraní", "Avañe'ẽ", "gn"},
		[]string{"Gujarati", "ગુજરાતી", "gu"},
		[]string{"Haitian, Haitian Creole", "Kreyòl ayisyen", "ht"},
		[]string{"Hausa", "(Hausa) هَوُسَ", "ha"},
		[]string{"Hebrew (modern)", "עברית", "he"},
		[]string{"Herero", "Otjiherero", "hz"},
		[]string{"Hindi", "हिन्दी, हिंदी", "hi"},
		[]string{"Hiri Motu", "Hiri Motu", "ho"},
		[]string{"Hungarian", "magyar", "hu"},
		[]string{"Interlingua", "Interlingua", "ia"},
		[]string{"Indonesian", "Bahasa Indonesia", "id"},
		[]string{"Interlingue", "Originally called Occidental; then Interlingue after WWII", "ie"},
		[]string{"Irish", "Gaeilge", "ga"},
		[]string{"Igbo", "Asụsụ Igbo", "ig"},
		[]string{"Inupiaq", "Iñupiaq, Iñupiatun", "ik"},
		[]string{"Ido", "Ido", "io"},
		[]string{"Icelandic", "Íslenska", "is"},
		[]string{"Italian", "Italiano", "it"},
		[]string{"Inuktitut", "ᐃᓄᒃᑎᑐᑦ", "iu"},
		[]string{"Japanese", "日本語 (にほんご)", "ja"},
		[]string{"Javanese", "ꦧꦱꦗꦮ, Basa Jawa", "jv"},
		[]string{"Kalaallisut, Greenlandic", "kalaallisut, kalaallit oqaasii", "kl"},
		[]string{"Kannada", "ಕನ್ನಡ", "kn"},
		[]string{"Kanuri", "Kanuri", "kr"},
		[]string{"Kashmiri", "कश्मीरी, كشميري‎", "ks"},
		[]string{"Kazakh", "қазақ тілі", "kk"},
		[]string{"Khmer", "ខ្មែរ, ខេមរភាសា, ភាសាខ្មែរ", "km"},
		[]string{"Kikuyu, Gikuyu", "Gĩkũyũ", "ki"},
		[]string{"Kinyarwanda", "Ikinyarwanda", "rw"},
		[]string{"Kyrgyz", "Кыргызча, Кыргыз тили", "ky"},
		[]string{"Komi", "коми кыв", "kv"},
		[]string{"Kongo", "Kikongo", "kg"},
		[]string{"Korean", "한국어", "ko"},
		[]string{"Kurdish", "Kurdî, كوردی‎", "ku"},
		[]string{"Kwanyama, Kuanyama", "Kuanyama", "kj"},
		[]string{"Latin", "latine, lingua latina", "la"},
		[]string{"Luxembourgish, Letzeburgesch", "Lëtzebuergesch", "lb"},
		[]string{"Ganda", "Luganda", "lg"},
		[]string{"Limburgish, Limburgan, Limburger", "Limburgs", "li"},
		[]string{"Lingala", "Lingála", "ln"},
		[]string{"Lao", "ພາສາລາວ", "lo"},
		[]string{"Lithuanian", "lietuvių kalba", "lt"},
		[]string{"Luba-Katanga", "Tshiluba", "lu"},
		[]string{"Latvian", "latviešu valoda", "lv"},
		[]string{"Manx", "Gaelg, Gailck", "gv"},
		[]string{"Macedonian", "македонски јазик", "mk"},
		[]string{"Malagasy", "fiteny malagasy", "mg"},
		[]string{"Malay", "bahasa Melayu, بهاس ملايو‎", "ms"},
		[]string{"Malayalam", "മലയാളം", "ml"},
		[]string{"Maltese", "Malti", "mt"},
		[]string{"Māori", "te reo Māori", "mi"},
		[]string{"Marathi (Marāṭhī)", "मराठी", "mr"},
		[]string{"Marshallese", "Kajin M̧ajeļ", "mh"},
		[]string{"Mongolian", "Монгол хэл", "mn"},
		[]string{"Nauruan", "Dorerin Naoero", "na"},
		[]string{"Navajo, Navaho", "Diné bizaad", "nv"},
		[]string{"Northern Ndebele", "isiNdebele", "nd"},
		[]string{"Nepali", "नेपाली", "ne"},
		[]string{"Ndonga", "Owambo", "ng"},
		[]string{"Norwegian Bokmål", "Norsk bokmål", "nb"},
		[]string{"Norwegian Nynorsk", "Norsk nynorsk", "nn"},
		[]string{"Norwegian", "Norsk", "no"},
		[]string{"Nuosu", "ꆈꌠ꒿ Nuosuhxop", "ii"},
		[]string{"Southern Ndebele", "isiNdebele", "nr"},
		[]string{"Occitan", "occitan, lenga d'òc", "oc"},
		[]string{"Ojibwe, Ojibwa", "ᐊᓂᔑᓈᐯᒧᐎᓐ", "oj"},
		[]string{"Old Church Slavonic, Church Slavonic, Old Bulgarian", "ѩзыкъ словѣньскъ", "cu"},
		[]string{"Oromo", "Afaan Oromoo", "om"},
		[]string{"Oriya", "ଓଡ଼ିଆ", "or"},
		[]string{"Ossetian, Ossetic", "ирон æвзаг", "os"},
		[]string{"(Eastern) Punjabi", "ਪੰਜਾਬੀ", "pa"},
		[]string{"Pāli", "पाऴि", "pi"},
		[]string{"Persian (Farsi)", "فارسی", "fa"},
		[]string{"Polish", "język polski, polszczyzna", "pl"},
		[]string{"Pashto, Pushto", "پښتو", "ps"},
		[]string{"Portuguese", "Português", "pt"},
		[]string{"Quechua", "Runa Simi, Kichwa", "qu"},
		[]string{"Romansh", "rumantsch grischun", "rm"},
		[]string{"Kirundi", "Ikirundi", "rn"},
		[]string{"Romanian", "Română", "ro"},
		[]string{"Russian", "Русский", "ru"},
		[]string{"Sanskrit (Saṁskṛta)", "संस्कृतम्", "sa"},
		[]string{"Sardinian", "sardu", "sc"},
		[]string{"Sindhi", "सिन्धी, سنڌي، سندھی‎", "sd"},
		[]string{"Northern Sami", "Davvisámegiella", "se"},
		[]string{"Samoan", "gagana fa'a Samoa", "sm"},
		[]string{"Sango", "yângâ tî sängö", "sg"},
		[]string{"Serbian", "српски језик", "sr"},
		[]string{"Scottish Gaelic, Gaelic", "Gàidhlig", "gd"},
		[]string{"Shona", "chiShona", "sn"},
		[]string{"Sinhalese, Sinhala", "සිංහල", "si"},
		[]string{"Slovak", "slovenčina, slovenský jazyk", "sk"},
		[]string{"Slovene", "slovenski jezik, slovenščina", "sl"},
		[]string{"Somali", "Soomaaliga, af Soomaali", "so"},
		[]string{"Southern Sotho", "Sesotho", "st"},
		[]string{"Spanish", "Español", "es"},
		[]string{"Sundanese", "Basa Sunda", "su"},
		[]string{"Swahili", "Kiswahili", "sw"},
		[]string{"Swati", "SiSwati", "ss"},
		[]string{"Swedish", "svenska", "sv"},
		[]string{"Tamil", "தமிழ்", "ta"},
		[]string{"Telugu", "తెలుగు", "te"},
		[]string{"Tajik", "тоҷикӣ, toçikī, تاجیکی‎", "tg"},
		[]string{"Thai", "ไทย", "th"},
		[]string{"Tigrinya", "ትግርኛ", "ti"},
		[]string{"Tibetan Standard, Tibetan, Central", "བོད་ཡིག", "bo"},
		[]string{"Turkmen", "Türkmen, Түркмен", "tk"},
		[]string{"Tagalog", "Wikang Tagalog", "tl"},
		[]string{"Tswana", "Setswana", "tn"},
		[]string{"Tonga (Tonga Islands)", "faka Tonga", "to"},
		[]string{"Turkish", "Türkçe", "tr"},
		[]string{"Tsonga", "Xitsonga", "ts"},
		[]string{"Tatar", "татар теле, tatar tele", "tt"},
		[]string{"Twi", "Twi", "tw"},
		[]string{"Tahitian", "Reo Tahiti", "ty"},
		[]string{"Uyghur", "ئۇيغۇرچە‎, Uyghurche", "ug"},
		[]string{"Ukrainian", "Українська", "uk"},
		[]string{"Urdu", "اردو", "ur"},
		[]string{"Uzbek", "Oʻzbek, Ўзбек, أۇزبېك‎", "uz"},
		[]string{"Venda", "Tshivenḓa", "ve"},
		[]string{"Vietnamese", "Tiếng Việt", "vi"},
		[]string{"Volapük", "Volapük", "vo"},
		[]string{"Walloon", "walon", "wa"},
		[]string{"Welsh", "Cymraeg", "cy"},
		[]string{"Wolof", "Wollof", "wo"},
		[]string{"Western Frisian", "Frysk", "fy"},
		[]string{"Xhosa", "isiXhosa", "xh"},
		[]string{"Yiddish", "ייִדיש", "yi"},
		[]string{"Yoruba", "Yorùbá", "yo"},
		[]string{"Zhuang, Chuang", "Saɯ cueŋƅ, Saw cuengh", "za"},
		[]string{"Zulu", "isiZulu", "zu"},
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
		Trail(WORKING, "Initializing Languages: [%s%d/%d%s]", colors.FG_GREEN_B, i+1, len(langs), colors.FG_NORMAL)
	}
	tx.Commit()
	Trail(OK, "Initializing Languages: [%s%d/%d%s]", colors.FG_GREEN_B, len(langs), len(langs), colors.FG_NORMAL)
}

// translate is used to get a translation from a multilingual fields
func translate(raw string, lang string, args ...bool) string {
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
	fileName := "./static/i18n/" + path + "." + lang + ".json"
	if _, err = os.Stat(fileName); os.IsNotExist(err) {
		Trail(WARNING, "Unrecognized path (%s) - fileName:%s", path, fileName)
		return term
	}
	buf, err = ioutil.ReadFile(fileName)
	if err != nil {
		Trail(ERROR, "Unable to read language file (%s)", fileName)
		return term
	}
	langMap := map[string]string{}
	err = json.Unmarshal(buf, &langMap)
	if err != nil {
		Trail(ERROR, "Unknown translation file format (%s)", path)
		return term
	}
	if val, ok := langMap[term]; ok {
		return strings.TrimPrefix(val, translateMe)
	}
	if lang != "en" {
		Tf(path, "en", term, args...)
		langMap[term] = translateMe + term
	} else {
		langMap[term] = term
		//saveLangFile(langMap, fileName)
		//pathParts := strings.Split(path, "/")
		//syncCustomTranslation(pathParts[0], pathParts[1])
	}

	saveLangFile(langMap, fileName)
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
		f.ErrMsg = strings.TrimPrefix(structLang.Fields[f.Name].ErrMsg[f.ErrMsg], translateMe)
		s.Fields[i] = f
	}
}

func getLanguage(r *http.Request) Language {
	langCookie, err := r.Cookie("language")

	if err != nil || langCookie == nil {
		return Language{}
	}
	lang := langCookie.Value
	language := Language{}
	Get(&language, "code=?", lang)
	return language
}
