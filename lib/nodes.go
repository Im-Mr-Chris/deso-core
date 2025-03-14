package lib

type DeSoNode struct {
	// Name of the node, displayed to users
	Name string

	// HTTPs URL to the node or app
	URL string

	// DeSo username of the node owner
	Owner string
}

//
// This list of nodes is maintained by the core DeSo developer team.
//
// If you run a DeSo node that has been online for at least one month you may submit a pull request to add your
// node to the list of nodes.
//
// When submitting a post, add the following to PostExtraData:
//   "Node": "ID"
//
// If your node is in the list then other nodes will be able to know where users are posting and can include
// a link to your node and give you free advertising.
//

var NODES = map[uint64]DeSoNode{
	1: {
		Name:  "DeSo",
		URL:   "https://node.deso.org",
		Owner: "diamondhands",
	},
	2: {
		Name:  "BitClout",
		URL:   "https://bitclout.com",
		Owner: "diamondhands",
	},
	3: {
		Name:  "Diamond",
		URL:   "https://diamondapp.com",
		Owner: "Zordon",
	},
	4: {
		Name:  "CloutFeed",
		URL:   "https://apps.apple.com/app/id1561532815",
		Owner: "Ribal",
	},
	5: {
		Name:  "Flick",
		URL:   "https://flickapp.com",
		Owner: "nigeleccles",
	},
	6: {
		Name:  "tijn's club",
		URL:   "https://tijn.club",
		Owner: "tijn",
	},
	7: {
		Name:  "Nacho Average",
		URL:   "https://nachoaverage.com/",
		Owner: "ClayPerryMusic",
	},
	8: {
		Name:  "love4src",
		URL:   "https://love4src.com",
		Owner: "kanshi",
	},
	9: {
		Name:  "Supernovas",
		URL:   "https://supernovas.app",
		Owner: "fransarthur",
	},
	10: {
		Name:  "GiftClout",
		URL:   "https://members.giftclout.com",
		Owner: "RajLahoti",
	},
	11: {
		Name:  "DeSocialWorld",
		URL:   "https://desocialworld.com",
		Owner: "edokoevoet",
	},
	12: {
		Name:  "NFTz",
		URL:   "https://nftz.zone",
		Owner: "mvanhalen",
	},
	13: {
		Name:  "Cloutible",
		URL:   "https://cloutible.club",
		Owner: "DawaynePerza",
	},
	14: {
		Name:  "Agbegbe",
		URL:   "https://agbegbe.org",
		Owner: "TheParkerazzi",
	},
	15: {
		Name:  "CloutingAround",
		URL:   "https://cloutingaround.dev",
		Owner: "TheParkerazzi",
	},
	16: {
		Name:  "MediaTech",
		URL:   "https://deso.mediatech.ventures",
		Owner: "paulobrien",
	},
	17: {
		Name:  "Mousai",
		URL:   "https://deso.mousai.stream",
		Owner: "marlonjm2k",
	},
	18: {
		Name:  "KoalaTBooks",
		URL:   "https://koalatbooks.com",
		Owner: "chriscelaya",
	},
	19: {
		Name:  "Beyond",
		URL:   "https://beyond.restartu.org",
		Owner: "RestartU",
	},
	20: {
		Name:  "DeverSo",
		URL:   "https://deverso.io/",
		Owner: "Nordian",
	},
	21: {
		Name:  "NinjaNode",
		URL:   "https://node.bitcloutapps.ninja",
		Owner: "mattpitts",
	},
	22: {
		Name:  "Tunel",
		URL:   "https://tunel.app",
		Owner: "hazrodriguez",
	},
}
