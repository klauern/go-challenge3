package mosaic

import (
	"strings"
	"testing"
)

func TestJsonUnmarshalInts(t *testing.T) {
	photos := Photos{}
	photos.Unmarshal(strings.NewReader(photos_json))

	cases := []struct {
		item string
		in   int
		want int
	}{
		{"photos.Page", photos.Page, 1},
		{"photos.Pages", photos.Pages, 5},
		{"photos.PerPage", photos.PerPage, 100},
		{"photos.PhotoList[0].Farm", photos.PhotoList[0].Farm, 8},
		{"photos.PhotoList[0].IsPublic", photos.PhotoList[0].IsPublic, 1},
		{"photos.PhotoList[0].IsFriendly", photos.PhotoList[0].IsFriendly, 0},
		{"photos.PhotoList[0].IsFamily", photos.PhotoList[0].IsFamily, 0},
	}

	for _, c := range cases {
		if c.in != c.want {
			t.Errorf("Parsing %q = %d, want %d", c.item, c.in, c.want)
		}
	}
}

func TestJsonUnmarshalStrings(t *testing.T) {
	photos := new(Photos)
	photos.Unmarshal(strings.NewReader(photos_json))

	cases := []struct {
		item string
		in   string
		want string
	}{
		{"photos.Total", photos.Total, "500"},
		{"photos.Stat", photos.Stat, ""},
		{"photos.PhotoList[0].Owner", photos.PhotoList[0].Owner, "60509750@N08"},
		{"photos.PhotoList[0].Id", photos.PhotoList[0].Id, "17256178635"},
		{"photos.PhotoList[0].Secret", photos.PhotoList[0].Secret, "1d410796ef"},
		{"photos.PhotoList[0].Server", photos.PhotoList[0].Server, "7724"},
		{"photos.PhotoList[0].Title", photos.PhotoList[0].Title, "Saitama sakura, Japan [Explore nr 1 - thank you all!]"},
	}

	for _, c := range cases {
		if c.in != c.want {
			t.Errorf("Parsing %q = %d, want %d", c.item, c.in, c.want)
		}
	}
}

const photos_json string = `{ "photos": { "page": 1, "pages": 5, "perpage": 100, "total": "500",
    "photo": [
      { "id": "17256178635", "owner": "60509750@N08", "secret": "1d410796ef", "server": "7724", "farm": 8, "title": "Saitama sakura, Japan [Explore nr 1 - thank you all!]", "ispublic": 1, "isfriend": 0, "isfamily": 0 },
      { "id": "17257735765", "owner": "128044807@N04", "secret": "4112eb6000", "server": "7647", "farm": 8, "title": "Staub & Mopeds @ Hausham", "ispublic": 1, "isfriend": 0, "isfamily": 0 },
      { "id": "17252996935", "owner": "87699814@N06", "secret": "58476d93f8", "server": "8746", "farm": 9, "title": "Morgens im Luch", "ispublic": 1, "isfriend": 0, "isfamily": 0 },
      { "id": "16632879074", "owner": "66460484@N06", "secret": "a5631826f2", "server": "8787", "farm": 9, "title": "trekkers", "ispublic": 1, "isfriend": 0, "isfamily": 0 },
      { "id": "17248431802", "owner": "98512152@N03", "secret": "c84e3f4a39", "server": "7695", "farm": 8, "title": "So Deep is the Night", "ispublic": 1, "isfriend": 0, "isfamily": 0 },
      { "id": "17253954111", "owner": "73331227@N05", "secret": "4ef15365d8", "server": "8802", "farm": 9, "title": "The Building Beyond", "ispublic": 1, "isfriend": 0, "isfamily": 0 },
      { "id": "17045641937", "owner": "94321001@N00", "secret": "083cab811e", "server": "8686", "farm": 9, "title": "the unbreakable circle", "ispublic": 1, "isfriend": 0, "isfamily": 0 },
      { "id": "17069058000", "owner": "68313863@N05", "secret": "f6844dcf41", "server": "7639", "farm": 8, "title": "The Wise Owl", "ispublic": 1, "isfriend": 0, "isfamily": 0 },
      { "id": "17063669360", "owner": "74970900@N08", "secret": "743b1ac367", "server": "8713", "farm": 9, "title": "Amén!", "ispublic": 1, "isfriend": 0, "isfamily": 0 },
      { "id": "17256908452", "owner": "109262173@N05", "secret": "60564fca33", "server": "8794", "farm": 9, "title": "Let's go!", "ispublic": 1, "isfriend": 0, "isfamily": 0 },
      { "id": "16633972103", "owner": "21692577@N02", "secret": "1a1841096f", "server": "7713", "farm": 8, "title": "Sunset at Percé, Gaspésie, Canada", "ispublic": 1, "isfriend": 0, "isfamily": 0 },
      { "id": "17233330966", "owner": "115925985@N04", "secret": "657fe06043", "server": "8755", "farm": 9, "title": "Any of you on Instagram?", "ispublic": 1, "isfriend": 0, "isfamily": 0 },
      { "id": "17067077320", "owner": "79688583@N03", "secret": "a5208dd706", "server": "8773", "farm": 9, "title": "Heavy Venice", "ispublic": 1, "isfriend": 0, "isfamily": 0 },
      { "id": "16632813993", "owner": "50566514@N02", "secret": "b31c32d1e0", "server": "7708", "farm": 8, "title": "The Dawn.", "ispublic": 1, "isfriend": 0, "isfamily": 0 },
      { "id": "16631770783", "owner": "98227287@N05", "secret": "9f0129e199", "server": "7672", "farm": 8, "title": "Red gates", "ispublic": 1, "isfriend": 0, "isfamily": 0 },
      { "id": "16631700743", "owner": "106365084@N07", "secret": "123be29c2a", "server": "8739", "farm": 9, "title": "HFF.....(Explore...24\/4\/15)", "ispublic": 1, "isfriend": 0, "isfamily": 0 },
      { "id": "17046497097", "owner": "74269181@N00", "secret": "bd9b48c19d", "server": "7725", "farm": 8, "title": "'Early Spring'", "ispublic": 1, "isfriend": 0, "isfamily": 0 },
      { "id": "17063554858", "owner": "125269494@N02", "secret": "0f74868e73", "server": "8693", "farm": 9, "title": "Morning has broken", "ispublic": 1, "isfriend": 0, "isfamily": 0 },
      { "id": "17070673139", "owner": "97623054@N05", "secret": "7ab07c00be", "server": "8750", "farm": 9, "title": "Blue Mountains, Australia", "ispublic": 1, "isfriend": 0, "isfamily": 0 },
      { "id": "17258800005", "owner": "10817753@N03", "secret": "959fe42d25", "server": "8767", "farm": 9, "title": "Stockholm, April 24, 2015", "ispublic": 1, "isfriend": 0, "isfamily": 0 },
      { "id": "17065484889", "owner": "127686746@N07", "secret": "9b5bf42943", "server": "8750", "farm": 9, "title": "Morgens am See", "ispublic": 1, "isfriend": 0, "isfamily": 0 },
      { "id": "17066981860", "owner": "82379842@N02", "secret": "e9049b2bb0", "server": "7637", "farm": 8, "title": "Abandoned FF Explored", "ispublic": 1, "isfriend": 0, "isfamily": 0 },
      { "id": "17066115589", "owner": "128653174@N02", "secret": "c3b12e61e7", "server": "8745", "farm": 9, "title": "Sunset behind the Süleymaniye", "ispublic": 1, "isfriend": 0, "isfamily": 0 },
      { "id": "16637648683", "owner": "110746823@N02", "secret": "e80dea5633", "server": "8725", "farm": 9, "title": "Tulip at Keukenhof", "ispublic": 1, "isfriend": 0, "isfamily": 0 },
      { "id": "17251130351", "owner": "28726735@N08", "secret": "28abd82f01", "server": "7651", "farm": 8, "title": "entwined {explored}", "ispublic": 1, "isfriend": 0, "isfamily": 0 },
      { "id": "17073186079", "owner": "11248435@N04", "secret": "fb74f8b3e8", "server": "8711", "farm": 9, "title": ".", "ispublic": 1, "isfriend": 0, "isfamily": 0 },
      { "id": "17254266501", "owner": "79159583@N08", "secret": "494aa5856c", "server": "7722", "farm": 8, "title": "cold place warm light", "ispublic": 1, "isfriend": 0, "isfamily": 0 },
      { "id": "17049009117", "owner": "81950116@N07", "secret": "77e21340d1", "server": "8757", "farm": 9, "title": "Jay", "ispublic": 1, "isfriend": 0, "isfamily": 0 },
      { "id": "17256500372", "owner": "75924892@N08", "secret": "77685c428f", "server": "7709", "farm": 8, "title": "Worms Head Long Exposure [EXPLORE 24-04-15]", "ispublic": 1, "isfriend": 0, "isfamily": 0 },
      { "id": "17249244932", "owner": "56662332@N04", "secret": "d25bf5e9a0", "server": "7610", "farm": 8, "title": "His And Hers!", "ispublic": 1, "isfriend": 0, "isfamily": 0 },
      { "id": "17253713205", "owner": "69656176@N05", "secret": "cdb0407415", "server": "8742", "farm": 9, "title": "Matera di notte", "ispublic": 1, "isfriend": 0, "isfamily": 0 },
      { "id": "17063473308", "owner": "7121172@N02", "secret": "32ab02ab98", "server": "7608", "farm": 8, "title": "Whatcom Falls Park, Infrared, Bellingham, Washington, explored #28", "ispublic": 1, "isfriend": 0, "isfamily": 0 },
      { "id": "17256590891", "owner": "125312165@N02", "secret": "a3e9769cf2", "server": "8740", "farm": 9, "title": "Sunset over Burgh Island", "ispublic": 1, "isfriend": 0, "isfamily": 0 },
      { "id": "17224282036", "owner": "59354145@N06", "secret": "9c5026de82", "server": "8820", "farm": 9, "title": "Deciduous Azalea   (explored)", "ispublic": 1, "isfriend": 0, "isfamily": 0 },
      { "id": "16629319704", "owner": "126616561@N08", "secret": "3d7f8aa58e", "server": "7695", "farm": 8, "title": "Magical sunrise at Rapa Nui", "ispublic": 1, "isfriend": 0, "isfamily": 0 },
      { "id": "17251737601", "owner": "90909195@N05", "secret": "494bba7cbf", "server": "7614", "farm": 8, "title": "Şile balıkçı tekneleri - Bahadırın köyü civarları - Dublin", "ispublic": 1, "isfriend": 0, "isfamily": 0 },
      { "id": "17232271386", "owner": "46005255@N02", "secret": "db4024f36a", "server": "7669", "farm": 8, "title": "Fell on Black Days", "ispublic": 1, "isfriend": 0, "isfamily": 0 },
      { "id": "17224120056", "owner": "60529780@N02", "secret": "17c5aaaa31", "server": "8822", "farm": 9, "title": "memories of summer", "ispublic": 1, "isfriend": 0, "isfamily": 0 },
      { "id": "17254944812", "owner": "85328224@N05", "secret": "6541e08b89", "server": "7668", "farm": 8, "title": "Cliffs of Moher, West of Ireland.", "ispublic": 1, "isfriend": 0, "isfamily": 0 },
      { "id": "17253724095", "owner": "53662163@N06", "secret": "a7b430101f", "server": "7722", "farm": 8, "title": "in commemoration", "ispublic": 1, "isfriend": 0, "isfamily": 0 },
      { "id": "17065562460", "owner": "55732328@N07", "secret": "be85a9dbaf", "server": "7664", "farm": 8, "title": "|| The New Change ||", "ispublic": 1, "isfriend": 0, "isfamily": 0 },
      { "id": "17071765729", "owner": "77194288@N04", "secret": "29f40467ab", "server": "7584", "farm": 8, "title": "Spring In Hood River Valley", "ispublic": 1, "isfriend": 0, "isfamily": 0 },
      { "id": "17062121048", "owner": "107793886@N08", "secret": "81c1db2802", "server": "7617", "farm": 8, "title": "The Illustrator", "ispublic": 1, "isfriend": 0, "isfamily": 0 },
      { "id": "17071422379", "owner": "58519338@N08", "secret": "0566610b8a", "server": "7640", "farm": 8, "title": "Red-necked Grebe", "ispublic": 1, "isfriend": 0, "isfamily": 0 },
      { "id": "16631480974", "owner": "41042844@N04", "secret": "e6f6825ab2", "server": "8776", "farm": 9, "title": "Tuscany, San Gimignano, via San Giovanni", "ispublic": 1, "isfriend": 0, "isfamily": 0 },
      { "id": "17067733518", "owner": "64056743@N07", "secret": "7f4b1c239d", "server": "8807", "farm": 9, "title": "An Oasis of the Concrete Jungle", "ispublic": 1, "isfriend": 0, "isfamily": 0 },
      { "id": "17047730547", "owner": "62990032@N03", "secret": "38a36a9354", "server": "7664", "farm": 8, "title": "Commuting", "ispublic": 1, "isfriend": 0, "isfamily": 0 },
      { "id": "17258664445", "owner": "55862496@N05", "secret": "04ea597fa3", "server": "8760", "farm": 9, "title": "Black biting fly", "ispublic": 1, "isfriend": 0, "isfamily": 0 },
      { "id": "16632016033", "owner": "58745059@N06", "secret": "e3b92d8888", "server": "7584", "farm": 8, "title": "Last days at my Tea Blossom", "ispublic": 1, "isfriend": 0, "isfamily": 0 },
      { "id": "17259640915", "owner": "77575637@N02", "secret": "80df5581dc", "server": "7653", "farm": 8, "title": "Hood River", "ispublic": 1, "isfriend": 0, "isfamily": 0 },
      { "id": "17052321777", "owner": "92748408@N06", "secret": "ffeb116413", "server": "7595", "farm": 8, "title": "Spanish Peaks", "ispublic": 1, "isfriend": 0, "isfamily": 0 },
      { "id": "17256293725", "owner": "81092501@N00", "secret": "a234e418e2", "server": "8694", "farm": 9, "title": "Ati-atihan Festival Participant, Kalibo, Aklan, Philippines", "ispublic": 1, "isfriend": 0, "isfamily": 0 },
      { "id": "17228482856", "owner": "12987346@N07", "secret": "9320599ecf", "server": "7706", "farm": 8, "title": "tranquillo mattino", "ispublic": 1, "isfriend": 0, "isfamily": 0 },
      { "id": "17253049225", "owner": "57982491@N06", "secret": "0dc140f4c8", "server": "8770", "farm": 9, "title": "Clevedon, Somerset, UK", "ispublic": 1, "isfriend": 0, "isfamily": 0 },
      { "id": "17070679909", "owner": "117605304@N07", "secret": "fe42efdafa", "server": "7641", "farm": 8, "title": "Shafer Trail Panorama: in EXPLORE 04\/24\/2015", "ispublic": 1, "isfriend": 0, "isfamily": 0 },
      { "id": "17258589805", "owner": "53609194@N00", "secret": "604ac8def5", "server": "8703", "farm": 9, "title": "Sunfall in April", "ispublic": 1, "isfriend": 0, "isfamily": 0 },
      { "id": "17251138242", "owner": "66592939@N02", "secret": "5e98b56249", "server": "7617", "farm": 8, "title": "A ninguna parte", "ispublic": 1, "isfriend": 0, "isfamily": 0 },
      { "id": "17230418266", "owner": "24796741@N05", "secret": "c205d81c5b", "server": "7643", "farm": 8, "title": "2 Parked Cadillacs - circa 1943", "ispublic": 1, "isfriend": 0, "isfamily": 0 },
      { "id": "16630982073", "owner": "34721981@N06", "secret": "6e5e5f91d0", "server": "8703", "farm": 9, "title": "Canned", "ispublic": 1, "isfriend": 0, "isfamily": 0 },
      { "id": "16637348943", "owner": "48150568@N06", "secret": "0914748d28", "server": "8821", "farm": 9, "title": "Industrial space", "ispublic": 1, "isfriend": 0, "isfamily": 0 },
      { "id": "17257166945", "owner": "64047197@N02", "secret": "629ea27107", "server": "7601", "farm": 8, "title": "SM268321", "ispublic": 1, "isfriend": 0, "isfamily": 0 },
      { "id": "17254404182", "owner": "86656817@N08", "secret": "0a96a62eb2", "server": "7641", "farm": 8, "title": "Greenish blue carabid", "ispublic": 1, "isfriend": 0, "isfamily": 0 },
      { "id": "17063149398", "owner": "45417826@N03", "secret": "2a56c7162b", "server": "7713", "farm": 8, "title": "Potter Falls (Explored)", "ispublic": 1, "isfriend": 0, "isfamily": 0 },
      { "id": "17045362717", "owner": "16934617@N07", "secret": "1770556e7d", "server": "8751", "farm": 9, "title": "Paloma Orgullosa (Explored).", "ispublic": 1, "isfriend": 0, "isfamily": 0 },
      { "id": "16630715063", "owner": "22269025@N07", "secret": "3d0d2ef49d", "server": "7656", "farm": 8, "title": "Take me home", "ispublic": 1, "isfriend": 0, "isfamily": 0 },
      { "id": "17071058039", "owner": "53196512@N07", "secret": "3ef4c022e7", "server": "8719", "farm": 9, "title": "Zlatorog (Mala Savica Rapids Lit Two Ways),  Gorenjska Slovenia", "ispublic": 1, "isfriend": 0, "isfamily": 0 },
      { "id": "17253037011", "owner": "63198319@N05", "secret": "9f62237785", "server": "8707", "farm": 9, "title": "", "ispublic": 1, "isfriend": 0, "isfamily": 0 },
      { "id": "17252054771", "owner": "62544416@N04", "secret": "87de57eb65", "server": "7606", "farm": 8, "title": "Poppies field [Explore nr 60]", "ispublic": 1, "isfriend": 0, "isfamily": 0 },
      { "id": "16636969134", "owner": "100850197@N06", "secret": "f941c0466b", "server": "7597", "farm": 8, "title": "liencres", "ispublic": 1, "isfriend": 0, "isfamily": 0 },
      { "id": "17068274350", "owner": "89187987@N03", "secret": "4dc0d8d150", "server": "8689", "farm": 9, "title": "…and they lived happily ever after", "ispublic": 1, "isfriend": 0, "isfamily": 0 },
      { "id": "17066445890", "owner": "61997941@N00", "secret": "8a0d2e448b", "server": "8805", "farm": 9, "title": "Beech-tree woods along the Thuringian Werra trail", "ispublic": 1, "isfriend": 0, "isfamily": 0 },
      { "id": "16635687274", "owner": "78597264@N00", "secret": "997f97ef5c", "server": "7668", "farm": 8, "title": "Anemone", "ispublic": 1, "isfriend": 0, "isfamily": 0 },
      { "id": "16632614774", "owner": "17516529@N00", "secret": "94694223c9", "server": "7701", "farm": 8, "title": "Foggy Fence Friday", "ispublic": 1, "isfriend": 0, "isfamily": 0 },
      { "id": "17250575512", "owner": "78683307@N04", "secret": "b3d3effe5e", "server": "8763", "farm": 9, "title": "YU Photograph 20150146", "ispublic": 1, "isfriend": 0, "isfamily": 0 },
      { "id": "17231937716", "owner": "93319057@N07", "secret": "4d1a81f71a", "server": "8719", "farm": 9, "title": "Casablanca - tea poring training [explored]", "ispublic": 1, "isfriend": 0, "isfamily": 0 },
      { "id": "17255935115", "owner": "60179651@N03", "secret": "9ef8e55611", "server": "7707", "farm": 8, "title": "Tropical Storm", "ispublic": 1, "isfriend": 0, "isfamily": 0 },
      { "id": "17254107292", "owner": "37193596@N00", "secret": "f8aac0e17b", "server": "8794", "farm": 9, "title": "Sunrise", "ispublic": 1, "isfriend": 0, "isfamily": 0 },
      { "id": "17250081641", "owner": "40696541@N06", "secret": "ebf6890af6", "server": "7665", "farm": 8, "title": "Dahlia (Close-up)", "ispublic": 1, "isfriend": 0, "isfamily": 0 },
      { "id": "17064942828", "owner": "28478695@N00", "secret": "b4ceb636cc", "server": "8690", "farm": 9, "title": "", "ispublic": 1, "isfriend": 0, "isfamily": 0 },
      { "id": "17062385268", "owner": "105827505@N05", "secret": "2fae9982c3", "server": "8815", "farm": 9, "title": "Red Barn in Spring # 66", "ispublic": 1, "isfriend": 0, "isfamily": 0 },
      { "id": "17068220550", "owner": "79397274@N00", "secret": "ec62aaa560", "server": "8785", "farm": 9, "title": "Arch.'60 - Mareggiata a Sant'Angelo d' Ischia", "ispublic": 1, "isfriend": 0, "isfamily": 0 },
      { "id": "17063534098", "owner": "60612262@N08", "secret": "180f920b78", "server": "8749", "farm": 9, "title": "Golden Crowned Sparrow - Zonotrichia atricapilla", "ispublic": 1, "isfriend": 0, "isfamily": 0 },
      { "id": "16637954353", "owner": "129084494@N04", "secret": "d4c34142af", "server": "7719", "farm": 8, "title": "Eternal Ache", "ispublic": 1, "isfriend": 0, "isfamily": 0 },
      { "id": "17258079705", "owner": "127735759@N07", "secret": "17a8eb9fb1", "server": "7598", "farm": 8, "title": "Margarita\/Daisy [Explore]", "ispublic": 1, "isfriend": 0, "isfamily": 0 },
      { "id": "16632193484", "owner": "39844181@N04", "secret": "1a28c80ee1", "server": "7658", "farm": 8, "title": "Make a wish - EXPLORED 4.24.15", "ispublic": 1, "isfriend": 0, "isfamily": 0 },
      { "id": "17256300015", "owner": "72484280@N08", "secret": "8126ba1d92", "server": "8775", "farm": 9, "title": "Where's all the sun gone?", "ispublic": 1, "isfriend": 0, "isfamily": 0 },
      { "id": "17254170032", "owner": "44532761@N02", "secret": "061a2eabb0", "server": "8763", "farm": 9, "title": "Estrella", "ispublic": 1, "isfriend": 0, "isfamily": 0 },
      { "id": "17065856519", "owner": "55066950@N08", "secret": "e727fdec71", "server": "8815", "farm": 9, "title": "Texas Size", "ispublic": 1, "isfriend": 0, "isfamily": 0 },
      { "id": "17063963499", "owner": "101866926@N03", "secret": "bf230b1bef", "server": "8716", "farm": 9, "title": "Golden Hour", "ispublic": 1, "isfriend": 0, "isfamily": 0 },
      { "id": "17048547787", "owner": "34479299@N08", "secret": "7b9f0c6069", "server": "8716", "farm": 9, "title": "are you looking at me...(explored)", "ispublic": 1, "isfriend": 0, "isfamily": 0 },
      { "id": "17248207972", "owner": "95607642@N04", "secret": "173e0d02a7", "server": "8771", "farm": 9, "title": "Everlasting", "ispublic": 1, "isfriend": 0, "isfamily": 0 },
      { "id": "17068936249", "owner": "86458486@N02", "secret": "232fb843cc", "server": "7663", "farm": 8, "title": "Hff!!... :)", "ispublic": 1, "isfriend": 0, "isfamily": 0 },
      { "id": "17066390378", "owner": "59298007@N03", "secret": "c76e5a9779", "server": "8702", "farm": 9, "title": "", "ispublic": 1, "isfriend": 0, "isfamily": 0 },
      { "id": "17254416571", "owner": "94466116@N02", "secret": "0a5ed61ba4", "server": "7703", "farm": 8, "title": "On top of the Empire State Building", "ispublic": 1, "isfriend": 0, "isfamily": 0 },
      { "id": "17069862189", "owner": "37995657@N00", "secret": "458ce82684", "server": "7666", "farm": 8, "title": "Polaroid Week Day Five", "ispublic": 1, "isfriend": 0, "isfamily": 0 },
      { "id": "17071779749", "owner": "53421519@N02", "secret": "63fb010248", "server": "7670", "farm": 8, "title": "Red Sunset", "ispublic": 1, "isfriend": 0, "isfamily": 0 },
      { "id": "17068837959", "owner": "29194479@N04", "secret": "ce15bf9f6c", "server": "8771", "farm": 9, "title": "Wild Idaho", "ispublic": 1, "isfriend": 0, "isfamily": 0 },
      { "id": "17063228110", "owner": "95833184@N05", "secret": "8c5e685d51", "server": "8717", "farm": 9, "title": "temporary fix", "ispublic": 1, "isfriend": 0, "isfamily": 0 },
      { "id": "16634155163", "owner": "44319079@N06", "secret": "05c6a73243", "server": "7699", "farm": 8, "title": "Mt. Fuji and Cherry Blssoms 富士山と桜", "ispublic": 1, "isfriend": 0, "isfamily": 0 },
      { "id": "17256907122", "owner": "67485055@N04", "secret": "5aec94961b", "server": "7615", "farm": 8, "title": "gargouille", "ispublic": 1, "isfriend": 0, "isfamily": 0 }
    ] }, "stat": "ok" }`
