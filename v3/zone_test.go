package geohex

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Encode", func() {
	var zone *Zone
	var err error

	for _, tc := range testCases {
		It("should encode to "+tc.code, func() {
			zone, err = Encode(tc.lat, tc.lon, tc.level)
			Expect(err).To(BeNil())
			Expect(zone.Code).To(Equal(tc.code))
		})
	}

	It("should wrap the position and level", func() {
		zone, err := Encode(85.04354094565655, 89.2529296875, 3)
		Expect(err).To(BeNil())
		Expect(zone.Level()).To(Equal(3))
		Expect(zone.Pos.X).To(Equal(271))
		Expect(zone.Pos.Y).To(Equal(150))
	})

})

var _ = Describe("Decode", func() {

	for _, tc := range testCases {
		It("should encode from "+tc.code, func() {
			exp := NewLL(tc.lat, tc.lon).Point().Position(zooms[tc.level]).LL()
			act, err := Decode(tc.code)
			Expect(err).To(BeNil())
			Expect(act).To(Equal(exp))
		})
	}

})

var testCases = []struct {
	lat   float64
	lon   float64
	level int
	code  string
}{
	{33.35137950146622, 135.6104480957031, 0, "XM"},
	{-1.150500622870068, -9.233301904296898, 0, "OY"},
	{-2.7315738409448347, 178.9405262207031, 0, "GI"},
	{4.6451064186120155, -174.81923940429692, 0, "QU"},
	{-23.490737690233317, -158.29580190429692, 0, "GH"},
	{-76.24971820610453, 176.3916980957031, 0, "EU"},
	{-85.25233946359653, 108.89169809570309, 0, "CI"},
	{-84.8900362156735, -83.7645519042969, 0, "BV"},
	{-47.22500123755036, -102.0458019042969, 0, "Fb"},
	{66.684617798596, -87.9833019042969, 0, "SV"},
	{84.64823341488797, -177.9833019042969, 0, "TO"},
	{82.27244849463305, 172.87607309570308, 0, "TK"},
	{82.99315605554614, -35.952051904296894, 0, "bD"},
	{76.61704901633043, 112.40732309570309, 0, "ZA"},
	{75.07094913292545, 24.516698095703102, 0, "aX"},
	{42.35263222764025, 30.141698095703102, 0, "PZ"},
	{-46.7453426931156, 61.7823230957031, 0, "OK"},
	{-78.40414562396407, -16.967676904296898, 1, "CZ3"},
	{-84.73838712095338, 119.53125, 1, "CI8"},
	{-85.01376874698711, 177.0948230957031, 1, "DO0"},
	{-75.40885422846453, -175.78125, 1, "EX2"},
	{-77.91566898632583, 175.78125, 1, "LK1"},
	{-51.39920565355377, 59.0625, 1, "OK4"},
	{-12.476571162995887, -175.87392690429692, 1, "GI3"},
	{-15.28418511407642, 174.375, 1, "Mc1"},
	{7.092954137951323, 178.5010730957031, 1, "QU0"},
	{15.361236013307861, -177.28017690429692, 1, "QU7"},
	{-16.63619187839765, -161.71875, 1, "GH8"},
	{-7.01366792756663, -1.40625, 1, "OY0"},
	{-51.618016548773696, -97.03125, 1, "Fb3"},
	{31.95216223802497, 140.625, 1, "XM4"},
	{38.88480596176492, 39.2823230957031, 1, "PZ0"},
	{70.61261423801925, -87.1875, 1, "SW6"},
	{65.94647177615738, -77.34375, 1, "SV3"},
	{76.51681887717322, 18.28125, 1, "aX0"},
	{79.44701237483491, 18.188573095703102, 1, "aX8"},
	{82.12933120948381, -33.842676904296894, 1, "aZ2"},
	{84.26717240431665, 178.59375, 1, "TO4"},
	{83.52016238353204, 174.375, 1, "aB6"},
	{81.30832090051811, 172.96875, 1, "aA1"},
	{81.09321385260837, -171.5625, 1, "TK3"},
	{74.4021625984244, 119.53125, 1, "ZA4"},
	{0.07989077839459471, -61.967676904296894, 1, "PA0"},
	{-84.54136107313407, -77.34375, 1, "BV8"},
	{-85.05112877980659, 130.78125, 2, "CI76"},
	{-85.04423207151001, 111.00107309570309, 2, "CI55"},
	{-84.94383661482844, -178.505859375, 2, "DO07"},
	{-84.92125351985577, 177.5342762207031, 2, "KE05"},
	{-79.23718500609334, -174.7265625, 2, "EU06"},
	{-78.61424480931518, 174.2823230957031, 2, "LK13"},
	{-27.059125784374054, -157.5, 2, "GH32"},
	{-47.989921667414166, 52.03125, 2, "OK51"},
	{-11.178401873711772, -178.59375, 2, "GI47"},
	{-8.328127830560586, 175.6885730957031, 2, "Mc53"},
	{7.092954137951794, 179.9073230957031, 2, "QU08"},
	{6.918485343601025, -165.50283315429692, 2, "OC58"},
	{19.476950206488414, 173.408203125, 2, "XL60"},
	{35.239085491966435, 137.0166980957031, 2, "XM56"},
	{22.34267675373961, 128.5791980957031, 2, "PS62"},
	{42.87010132809042, 142.6416980957031, 2, "XX03"},
	{40.97989806962013, 45.703125, 2, "PZ38"},
	{36.03133177633187, 38.671875, 2, "PZ04"},
	{40.97989806962013, 33.046875, 2, "PZ18"},
	{-12.133085454027368, -0.09267690429689779, 2, "OU88"},
	{3.5134210456400443, -4.921875, 2, "OY53"},
	{4.214943141390651, -59.765625, 2, "PA04"},
	{-51.618016548773696, -105.46875, 2, "Fb17"},
	{-51.618016548773696, -93.515625, 2, "Fb34"},
	{-44.28239868260753, -95.7176769042969, 2, "Fb86"},
	{-17.30868788677002, -167.34375, 2, "GH58"},
	{81.5816722445285, -165.3270519042969, 2, "TK63"},
	{83.72428346705478, -177.9833019042969, 2, "TO07"},
	{84.44784177073191, 176.39169809570308, 2, "aE42"},
	{80.4157074446218, 175.78125, 2, "aA02"},
	{82.45925547160363, -35.952051904296894, 2, "bD31"},
	{84.64823341488803, -51.420801904296894, 2, "bE73"},
	{85.03524922742024, 90.0831043457031, 2, "bb33"},
	{85.0511287798066, 179.912109375, 2, "TO84"},
	{84.5162211763226, -177.275390625, 2, "TO47"},
	{83.163581025311, 171.826171875, 2, "aB61"},
	{73.52839948765174, 119.1796875, 2, "ZA08"},
	{78.56048828398782, 10.8984375, 2, "aX55"},
	{75.84516854027044, 21.09375, 2, "aX04"},
	{76.59669080369527, 33.3936512207031, 2, "YF58"},
	{66.65297740055279, -80.859375, 2, "SV40"},
	{47.98456741915722, -108.2860362792969, 2, "RX15"},
	{-25.4901725033649, -1.0594737792969227, 2, "OU40"},
	{-85.05112877980659, -82.96875, 2, "BV56"},
	{-85.05112877980659, 128.3203125, 3, "CI771"},
	{-85.06248881025151, 179.3408203125, 3, "KE045"},
	{-85.03784245526992, 179.23095703125, 3, "KE056"},
	{-84.96579544427098, -177.28017690429692, 3, "DO077"},
	{-84.95930495623834, -70.83984375, 3, "BV778"},
	{-85.06626970363817, -82.177734375, 3, "BV568"},
	{-78.4905516616031, -12.12890625, 3, "CZ335"},
	{-76.55774293896553, 175.4296875, 3, "LK570"},
	{-75.53762333496856, -175.41250112304692, 3, "EU868"},
	{-49.382372787009544, 64.6875, 3, "OK468"},
	{-0.7031073524364783, 0.703125, 3, "OY443"},
	{-10.487811882056683, 0.0, 3, "OY004"},
	{-47.5172006978394, -109.6875, 3, "Fb526"},
	{-15.28418511407642, -166.640625, 3, "GH588"},
	{-12.197523144307803, 179.66562387695296, 3, "GI440"},
	{-11.509395970397813, -172.60000112304704, 3, "GI622"},
	{6.678485780144587, -173.30312612304704, 3, "QU340"},
	{4.214943141390651, 177.1875, 3, "XK026"},
	{34.02782288983624, 137.47812387695296, 3, "XM454"},
	{22.59372606392931, 128.49609375, 3, "PS624"},
	{44.33956524809713, 142.294921875, 3, "XX038"},
	{40.04443758460856, 135.52734375, 3, "XU630"},
	{39.095962936305476, 45.3515625, 3, "PZ342"},
	{46.08292326673606, 33.239842626952964, 3, "PZ580"},
	{7.710991655433229, 4.21875, 3, "OY864"},
	{-3.513421045640032, -7.734375, 3, "OY145"},
	{2.8113711933311403, -59.4140625, 3, "PA032"},
	{27.695900709062915, -115.47109487304704, 3, "OI758"},
	{67.7427590666639, -84.7265625, 3, "SV428"},
	{67.06743335108298, -74.1796875, 3, "SV385"},
	{77.07878389624943, 23.90625, 3, "aX355"},
	{78.56325923135681, 18.122655126952964, 3, "aX568"},
	{83.599030708362, -33.046875, 3, "bD744"},
	{83.55971676457146, -49.921875, 3, "bD515"},
	{85.04354094565655, 89.2529296875, 3, "bb337"},
	{84.9785740239428, 179.6484375, 3, "TO808"},
	{84.5162211763226, 175.2978515625, 3, "aE428"},
	{84.55521389397988, -177.52187612304704, 3, "TO474"},
	{83.70109641601667, -176.748046875, 3, "TO073"},
	{83.27255896509399, 173.5400390625, 3, "aB648"},
	{70.64639982680607, 97.70761606445296, 3, "YG168"},
	{-85.05870502661288, -76.46484375, 4, "BV7511"},
	{-85.0283304449441, -75.9375, 4, "BV7541"},
	{-85.04354094565655, 177.978515625, 4, "KE0532"},
	{-84.87239245838492, -177.47793081054706, 4, "DO3203"},
	{-77.595847144363, -170.62246206054706, 4, "EU6235"},
	{-78.34941069014627, 171.5625, 4, "LK1425"},
	{-47.989921667414166, 64.6875, 4, "OK7172"},
	{-78.9039293885709, -16.875, 4, "CZ0737"},
	{-55.321195006419345, -104.52871206054704, 4, "Fb0200"},
	{-45.08903556483102, -95.625, 4, "Fb8608"},
	{-23.241346102386135, -160.3125, 4, "GH4405"},
	{-11.178401873711772, 170.15625, 4, "Mc2651"},
	{-1.4061088354351594, -177.1875, 4, "OC2121"},
	{17.97873309555617, -172.96875, 4, "QU7871"},
	{11.178401873711785, 174.375, 4, "XK4212"},
	{-50.736455137010644, 63.28125, 4, "OK4362"},
	{-10.387634030186756, 5.158787939452964, 4, "OX2538"},
	{3.6150874592425724, -0.46621206054703634, 4, "OY5663"},
	{0.0, -60.46875, 4, "OG6663"},
	{35.460669951495305, 45.0, 4, "PZ0631"},
	{70.61261423801925, -82.96875, 4, "SV8518"},
	{81.09321385260837, -177.1875, 4, "TK3246"},
	{81.5182718765338, 171.5625, 4, "aA1825"},
	{83.35951133035452, 174.375, 4, "aB6727"},
	{85.05870502661288, -171.650390625, 4, "TR1202"},
	{82.9403268016951, -40.78125, 4, "bD4054"},
	{77.76758238272801, 20.390625, 4, "aX4486"},
	{76.01609366420995, 117.421875, 4, "ZA5685"},
	{35.38904996691167, 139.7021484375, 4, "XM4881"},
	{40.27952566881291, 135.0, 4, "XU6302"},
	{43.61221676817573, 143.1298828125, 4, "XX0337"},
	{32.80574473290688, 151.5673828125, 4, "XM6425"},
	{22.715390019335942, 129.5947265625, 4, "PS6270"},
	{-85.08136441846642, -79.453125, 5, "BV80302"},
	{-85.05112877980659, -78.92578125, 5, "BV80373"},
	{-84.92832092949963, -179.47265625, 5, "DO08354"},
	{-84.95930495623834, 178.2421875, 5, "KE05758"},
	{-78.73350050778467, -176.1328125, 5, "EU31135"},
	{-77.46299575294476, 177.60019418945296, 5, "LK45042"},
	{-78.20656311074711, -15.46875, 5, "CZ35087"},
	{-54.572061655658516, 64.3359375, 5, "OK06445"},
	{-46.31658418182218, -100.1953125, 5, "Fb80457"},
	{-17.644022027872722, -160.6640625, 5, "GH80422"},
	{-7.174237640658758, -172.73183706054704, 5, "GI74753"},
	{-6.476099848825542, 178.47910043945296, 5, "Mc80240"},
	{2.649746548777155, 176.72128793945296, 5, "XK01265"},
	{18.49286089121868, -174.84121206054704, 5, "QU86344"},
	{4.565473550710291, 9.84375, 5, "OY77334"},
	{-2.460181181020993, -9.140625, 5, "OY15454"},
	{3.5134210456400443, -60.46875, 5, "PA04170"},
	{21.94304553343818, 127.96875, 5, "PS62113"},
	{24.459869026557232, 133.91855356445296, 5, "PS68342"},
	{37.64903402157866, 145.810546875, 5, "XM78145"},
	{35.02999636902565, 138.33984375, 5, "XM48257"},
	{43.51668853502906, 141.416015625, 5, "XX03156"},
	{43.068887774169625, 43.59375, 5, "PZ47685"},
	{66.51326044311185, -78.75, 5, "SV40654"},
	{25.733322875649765, -115.25136831054692, 5, "OI71873"},
	{76.9999351181161, 18.984375, 5, "aX40264"},
	{74.95939165894974, 118.828125, 5, "ZA44552"},
	{83.27770503961696, -38.671875, 5, "bD47117"},
	{82.02137801950887, 172.265625, 5, "aA51615"},
	{82.02137801950887, -174.375, 5, "TK71615"},
	{84.32968167034325, -173.6107433105468, 5, "TO38458"},
	{85.04354094565655, 173.3203125, 5, "aE58484"},
	{-85.035941506574, -88.41796875, 6, "BV553581"},
	{-85.02070774312593, 118.4765625, 6, "CI802417"},
	{-85.1114157806266, 124.1015625, 6, "CI750110"},
	{-85.14128398117636, 177.890625, 6, "KE018862"},
	{-85.05112877980659, 178.2421875, 6, "KE053505"},
	{-85.01949481381588, -177.1263683105467, 6, "DO074664"},
	{-78.1344931829381, -16.5234375, 6, "CZ351822"},
	{-77.8418477505252, 176.1328125, 6, "LK412420"},
	{-77.15405630624801, -173.6107433105467, 6, "EU708101"},
	{-44.59046718130883, 60.8203125, 6, "OK847038"},
	{-21.616579336740593, -166.9921875, 6, "GH501658"},
	{-3.8503134325701494, 179.0064441894533, 6, "Mc842644"},
	{-4.915832801313164, -176.484375, 6, "GI836575"},
	{6.678485780329861, -174.3138683105467, 6, "QU316588"},
	{9.462844802358449, 174.7876941894533, 6, "XK175485"},
	{-48.45835188280864, -96.6796875, 6, "Fb711277"},
	{4.565473550710291, 2.4609375, 6, "OY728344"},
	{2.4601811810210052, -60.46875, 6, "PA016372"},
	{10.14193168613103, -180.0, 6, "QU408880"},
	{32.54681317351514, 138.8671875, 6, "XM442337"},
	{21.94304553343818, 127.6171875, 6, "PS387785"},
	{43.83452678223684, 143.0859375, 6, "XX037037"},
	{39.63953756436671, 36.9140625, 6, "PZ173676"},
	{65.80277639340238, -78.3984375, 6, "SV321726"},
	{39.78550768453777, -104.6605479980467, 6, "PC828583"},
	{24.379841066454198, -114.5042979980467, 6, "OI713375"},
	{82.40242347938855, -40.078125, 6, "bD080050"},
	{78.9039293885709, 18.28125, 6, "aX817100"},
	{75.49715731893085, 123.046875, 6, "ZA712478"},
	{81.06812649632941, 179.4019520019533, 6, "aA166315"},
	{82.83322366624273, -177.0824229980467, 6, "TK873482"},
	{84.25098665623719, -177.0824229980467, 6, "TO433402"},
	{85.00664369479912, 175.0074207519533, 6, "aE586342"},
	{85.03974267783215, 88.41796875, 6, "bb335332"},
	{-85.04733631224822, -79.6728515625, 7, "BV8032788"},
	{-85.05112877980659, 128.0126953125, 7, "CI7714524"},
	{-85.05207644397981, 177.220458984375, 7, "KE0514686"},
	{-85.05112877980659, 177.36328125, 7, "KE0517300"},
	{-85.05207644397981, -177.352294921875, 7, "DO0732687"},
	{-85.05018093458115, -177.396240234375, 7, "DO0732768"},
	{-78.20656311074711, -179.47265625, 7, "EU4031833"},
	{-76.32524311648964, -179.8180187011717, 7, "EU8072131"},
	{-75.6442438813013, 174.5569812988283, 7, "LK8255485"},
	{-78.06198918665973, -13.359375, 7, "CZ3800810"},
	{-52.26815737376816, 56.42578125, 7, "OK1624776"},
	{-47.32190612004859, 68.0335437988283, 7, "OK7464515"},
	{-48.45835188280864, -104.58984375, 7, "Fb5316746"},
	{-51.835777520452474, -99.140625, 7, "Fb4033527"},
	{-21.453068633086772, -162.24609375, 7, "GH4540858"},
	{-12.382928338487396, -178.76953125, 7, "GI4351804"},
	{-4.102390498486627, -177.5328624511717, 7, "GI8387484"},
	{4.327479492112243, -177.1812999511717, 7, "QU0706771"},
	{14.714023344909329, -179.9937999511717, 7, "QU4884047"},
	{19.168745200320153, 174.4251453613283, 7, "XK8251271"},
	{5.266007882805498, -60.8203125, 7, "PA0565644"},
	{6.165370163645122, 0.5774891113283047, 7, "OY8320288"},
	{-2.1088986592431254, 4.21875, 7, "OY3556486"},
	{45.336701909968106, 39.7265625, 7, "PZ8016814"},
	{39.90973623453719, 36.2109375, 7, "PZ1746604"},
	{22.43134015636061, 128.408203125, 7, "PS6242083"},
	{35.817813158696616, 137.900390625, 7, "XM5643552"},
	{40.212440718286466, 135.7470703125, 7, "XU6303820"},
	{38.20365531807151, 145.6787109375, 7, "XM7851800"},
	{43.866218006556394, 142.119140625, 7, "XX0345653"},
	{30.84821145703608, 156.1438953613283, 7, "XM6306347"},
	{21.453068633086783, 178.857421875, 7, "XK8827707"},
	{42.032974332441405, 35.15625, 7, "PZ4253332"},
	{30.92363839408563, -116.4928233886717, 7, "OI8776718"},
	{41.131270837014384, -95.3990733886717, 7, "PF2148656"},
	{75.40885422846455, 118.125, 7, "ZA4587731"},
	{66.79190947341796, -77.34375, 7, "SV4335020"},
	{69.72675675727797, -86.6100108886717, 7, "SV5843218"},
	{76.01609366420995, 23.90625, 7, "aX0682573"},
	{78.66951243420486, 17.804051611328305, 7, "aX8100870"},
	{83.19489563661588, -41.484375, 7, "bD4424804"},
	{84.18123616265304, -48.992823388671695, 7, "bE6100113"},
	{81.87364125482827, 171.9140625, 7, "aA5027217"},
	{81.87364125482827, -172.6171875, 7, "TK7070414"},
	{84.40594104126978, -171.2109375, 7, "TO6240542"},
	{84.95930495623834, -170.5078125, 7, "TO7775738"},
	{85.00759985598985, 171.7884266113283, 7, "aE5821675"},
	{85.05089183547521, 89.41497802734375, 7, "bb3371862"},
	{-85.05112877980659, -86.484375, 8, "BV54870250"},
	{-84.89434602717608, -74.96621968994123, 8, "BV78253862"},
	{-84.95930495623834, 116.015625, 8, "CI57758513"},
	{-85.05112877980659, 125.859375, 8, "CI75600226"},
	{-85.04733631224822, 179.384765625, 8, "KE05603615"},
	{-85.04461761710783, -178.94082906494123, 8, "DO07152160"},
	{-78.02557363284087, 178.41796875, 8, "LK41063674"},
	{-76.78350465189983, -179.38028218994123, 8, "EU72201265"},
	{-78.27820145542812, -12.3046875, 8, "CZ34762511"},
	{-50.736455137010644, -94.921875, 8, "Fb38241503"},
	{-48.05704077771629, 58.97909281005877, 8, "OK48511367"},
	{-16.057790563159426, -177.97403218994123, 8, "GI32042116"},
	{-8.506430649777663, -175.86465718994123, 8, "GI71565064"},
	{-21.28937435586041, -162.421875, 8, "GH45442518"},
	{-5.015809753574582, 6.9478428100587735, 8, "OY33228746"},
	{2.8113711933311403, -59.0625, 8, "PA03278283"},
	{11.178401873711785, 175.078125, 8, "XK42174351"},
	{12.28489094009069, -179.20450093994123, 8, "QU47122108"},
	{29.22889003019423, 140.9765625, 8, "XM32230826"},
	{35.23544008841136, 138.43221781005877, 8, "XM56301016"},
	{21.453068633086783, 129.7265625, 8, "PS61272751"},
	{31.11753019833845, 153.37362406005877, 8, "XM60832123"},
	{40.58058466412761, 134.296875, 8, "XU63124187"},
	{43.507650004750154, 141.77206156005877, 8, "XX03173236"},
	{41.37680856570233, 47.109375, 8, "PZ38755703"},
	{46.00356753757519, 36.12753031005877, 8, "PZ57757416"},
	{-39.85185459365323, -129.63418843994123, 8, "Fc51504013"},
	{0.7031073524364909, -58.359375, 8, "OW22586402"},
	{36.092306116910116, -111.35293843994123, 8, "PC52458806"},
	{66.82161538287745, -93.77481343994123, 8, "SV20852101"},
	{77.61770905279676, 24.609375, 8, "aX46486040"},
	{74.95939165894974, 129.375, 8, "ZA62754057"},
	{83.44032649527307, -52.03125, 8, "bD28367162"},
	{80.76061470752451, 177.890625, 8, "aA05722747"},
	{82.50466921991138, -178.85293843994123, 8, "TK83151057"},
	{83.83804706159876, -176.04043843994123, 8, "TO31130720"},
	{84.61492390223334, 172.00643656005877, 8, "aE51331450"},
	{85.04923290826919, 88.857421875, 8, "bb33488116"},
	{-85.02070774312593, -80.5078125, 9, "BV804240512"},
	{-85.25894723497555, 126.5625, 9, "CI708565326"},
	{-84.9901001802348, 124.453125, 9, "CI754848748"},
	{-84.96316466091037, 176.7919921875, 9, "KE055580431"},
	{-85.05065487982755, 178.5113525390625, 9, "KE053480025"},
	{-85.05065487982755, -177.2314453125, 9, "DO073506567"},
	{-78.34941069014627, -12.65625, 9, "CZ347013544"},
	{-77.91566898632583, 178.59375, 9, "LK413474586"},
	{-78.06198918665973, -175.78125, 9, "EU354580674"},
	{-48.92249926375824, 61.875, 9, "OK474555348"},
	{-47.989921667414166, -99.84375, 9, "Fb484868583"},
	{-25.79989118208832, -165.9375, 9, "GH147081657"},
	{-9.797132738815217, -170.64615133056623, 9, "GI730641802"},
	{-6.316766263346645, 176.69759866943377, 9, "Mc576514101"},
	{-8.407168163601074, -0.703125, 9, "OY041035522"},
	{2.809896281504813, -4.0055263305662265, 9, "OY537117080"},
	{9.100638641982364, -61.66177633056623, 9, "PA402570238"},
	{6.315298538330033, 177.890625, 9, "XK058877078"},
	{17.977328506752137, -179.43521383056623, 9, "QU832828267"},
	{40.97989806962013, 45.3515625, 9, "PZ382714446"},
	{36.59788913307022, 36.2109375, 9, "PZ028668414"},
	{20.960060644697037, 127.83041116943377, 9, "PS383657231"},
	{30.447400642478012, 151.73666116943377, 9, "XM602176827"},
	{34.884719615800655, 138.37728616943377, 9, "XM482452815"},
	{43.32410342218456, 142.59603616943377, 9, "XX033123366"},
	{28.612163074991937, -115.09927633056623, 9, "OI860616105"},
	{46.315564263468865, -108.77115133056623, 9, "RX114077466"},
	{66.51326044311185, -76.2890625, 9, "SV355072242"},
	{70.84418808231302, 98.29916116943377, 9, "YG413103084"},
	{75.49678751018848, 121.50228616943377, 9, "ZA486483578"},
	{77.54209596075547, 16.875, 9, "aX426174078"},
	{78.49025701498567, 27.283536169433773, 9, "aX747576237"},
	{81.5182718765338, 173.3203125, 9, "aA184816376"},
	{83.8297867078656, -175.91958883056623, 9, "TO310523631"},
	{84.40579709229564, 176.34603616943377, 9, "aE423806423"},
	{84.89701560751682, 87.40072366943377, 9, "bb303856635"},
	{85.04970694406315, 176.890869140625, 9, "aE815178640"},
	{65.94647177615738, -84.375, 10, "SV1700305142"},
	{-85.05112877980659, -87.022705078125, 10, "BV5482764471"},
	{-85.05018093458115, -74.02587890625, 10, "BV7560068257"},
	{-85.14128398117636, 110.7421875, 10, "CI5263615540"},
	{-85.1114157806266, 123.3984375, 10, "CI7265571616"},
	{-85.08136441846642, 124.1015625, 10, "CI7505181555"},
	{-85.05112877980659, 116.71875, 10, "CI5736546251"},
	{-85.1114157806266, 176.484375, 10, "KE0273028158"},
	{-85.05112877980659, 178.9453125, 10, "KE0536546500"},
	{-85.05112877980659, -177.890625, 10, "DO0716764500"},
	{-78.20656311074711, 179.6484375, 10, "LK4016113514"},
	{-76.95198835609855, 178.87838480224627, 10, "LK4827777232"},
	{-76.6712633375603, -173.73880269775373, 10, "EU7445272188"},
	{-51.17934297928927, 56.6015625, 10, "OK1777354303"},
	{-78.420193275912, -14.765625, 10, "CZ3420476772"},
	{-50.95842672335992, -97.03125, 10, "Fb4334484641"},
	{-15.961329081596647, -166.9921875, 10, "GH5856584625"},
	{-14.569063542697467, 176.24166605224627, 10, "Mc1730824803"},
	{-8.05922962720018, -177.1875, 10, "GI7234521114"},
	{10.487811882056683, -174.375, 10, "QU3851424320"},
	{17.97873309555617, 176.484375, 10, "XK8152020110"},
	{3.8642546157214084, -61.171875, 10, "PA0426257051"},
	{0.0, -4.21875, 10, "OY4244664644"},
	{2.848301563708052, 8.194791052246273, 10, "OY7353226324"},
	{-3.8642546157213955, 6.6796875, 10, "OY3444432334"},
	{17.308687886770034, 178.9453125, 10, "XK8136543688"},
	{20.74959624356506, -179.58352926025373, 10, "QU8831183154"},
	{21.779905342529645, 129.375, 10, "PS6206841033"},
	{27.326543260529995, 139.10787698974627, 10, "XM0565771230"},
	{34.69976336974951, 141.56881448974627, 10, "XM4863207510"},
	{43.28812706641589, 142.97506448974627, 10, "XX0330807788"},
	{30.858527238035272, 152.64303323974627, 10, "XM6054354111"},
	{40.17887331434696, 31.2890625, 10, "PZ1548687734"},
	{45.91573852227541, 43.48287698974627, 10, "PZ7555070851"},
	{27.839076094777816, -115.83984375, 10, "OI7581576564"},
	{47.989921667414166, -108.28125, 10, "RX1535370770"},
	{67.7427590666639, -72.7734375, 10, "SV7031782060"},
	{71.30079291637452, -85.95703125, 10, "SW6431665682"},
	{76.63922560965885, 17.40234375, 10, "aX1617612856"},
	{78.48038815858817, 21.158658239746273, 10, "aX7255422107"},
	{75.40885422846455, 127.6171875, 10, "ZA7323011017"},
	{80.13996914990166, 178.48287698974627, 10, "aA0130411728"},
	{82.41891476310745, -179.40774801025373, 10, "TK8071386820"},
	{84.38379917653626, -179.40774801025373, 10, "TO4434206021"},
	{85.0511287798066, 89.06478881835938, 10, "bb3372146218"},
	{-85.02070774312593, -85.4296875, 11, "BV57264048634"},
	{-85.02070774312593, 114.2578125, 11, "CI57246068382"},
	{-85.02070774312593, 177.890625, 11, "KE05443747884"},
	{-85.08136441846642, -178.9453125, 11, "DO04773823048"},
	{-85.08136441846642, -178.9453125, 11, "DO04773823048"},
	{-85.05112877980659, -179.296875, 11, "DO07200264504"},
	{-78.20656311074711, -13.0078125, 11, "CZ34837816828"},
	{-77.69287033641926, 174.0234375, 11, "LK18781747622"},
	{-77.28204307452988, -177.18988302001935, 11, "EU47467037811"},
	{-53.330872983017045, -102.65625, 11, "Fb05827021141"},
	{-47.75409797968002, 62.578125, 11, "OK72311284586"},
	{-20.3034175184893, -168.046875, 11, "GH51383310718"},
	{-10.351498489623582, -177.89300802001935, 11, "GI47416263465"},
	{-3.3750724311904046, -177.89300802001935, 11, "GI87045185255"},
	{4.002527103350086, -177.89300802001935, 11, "QU07002763476"},
	{20.05593126519445, 179.912109375, 11, "XK85663663436"},
	{18.56294744288831, 178.41796875, 11, "XK84201138888"},
	{22.105998799750566, 128.232421875, 11, "PS62147032100"},
	{26.947688436082007, 146.15972635498065, 11, "XM30178015132"},
	{29.88461929615859, 153.10308572998065, 11, "XM60054311737"},
	{32.99023555965106, 139.74609375, 11, "XM44560457183"},
	{36.00125691778753, 137.72222635498065, 11, "XM56456700534"},
	{43.32517767999296, 141.6796875, 11, "XX03131786872"},
	{38.8225909761771, 38.671875, 11, "PZ16386410216"},
	{28.116713872100473, -113.82074239501935, 11, "OI78178446320"},
	{46.41222705726201, -106.08636739501935, 11, "RX13217487787"},
	{66.23145747862573, -80.15625, 11, "SV40163034106"},
	{77.157162522661, 18.28125, 11, "aX41314430862"},
	{75.14077784070429, 120.9375, 11, "ZA47238521110"},
	{81.06018065538063, 170.94488260498065, 11, "aA11857485652"},
	{81.48680323045866, -179.91449239501935, 11, "TK43222807032"},
	{83.54985076685978, -178.2421875, 11, "TO03574622544"},
	{84.73838712095339, 172.353515625, 11, "aE54114446763"},
	{85.05041791286303, 88.1597900390625, 11, "bb33540877533"},
	{-85.05065487982755, -79.82666015625, 12, "BV803254745602"},
	{-85.05112877980659, 120.706787109375, 12, "CI803456424553"},
	{-85.05207644397981, 178.868408203125, 12, "KE053652024200"},
	{-85.05112877980659, 178.934326171875, 12, "KE053654462750"},
	{-85.05112877980659, -178.253173828125, 12, "DO071704644727"},
	{-79.44522664358297, -13.702334191894394, 12, "CY281612610161"},
	{-77.91566898632583, 172.96875, 12, "LK181655132036"},
	{-76.18499546094715, -179.296875, 12, "EU832016181016"},
	{-57.36436751908976, -102.9992091918944, 12, "FY636801583240"},
	{-49.837982453084834, 61.875, 12, "OK446658243285"},
	{-27.74562238880946, -162.06170919189435, 12, "GH058068815001"},
	{-13.308212759686633, 177.5476658081056, 12, "Mc410587505268"},
	{-2.881425891238504, -178.23358419189435, 12, "GI871735246286"},
	{10.4188347787547, -175.42108419189435, 12, "QU460311054340"},
	{13.855313517221834, 176.1414158081056, 12, "XK533828300764"},
	{-5.615985819155327, 4.921875, 12, "OY313181386161"},
	{4.845945917726586, 0.36016580810560583, 12, "OY803527038734"},
	{3.5134210456400443, -59.765625, 12, "PA043284261600"},
	{35.40351865149676, 44.657040808105606, 12, "PZ060785824214"},
	{22.59372606392931, 128.3203125, 12, "PS625060382624"},
	{28.860219752558816, 144.8523533081056, 12, "XM317558565378"},
	{29.47418457855351, 153.9929783081056, 12, "XL822458403751"},
	{42.5013892782751, 143.0945408081056, 12, "XP251226173262"},
	{26.36840163579232, -113.89764669189435, 12, "OI745463230145"},
	{43.52955855271223, -107.92108419189435, 12, "RU681285781273"},
	{66.37275500247456, -76.9921875, 12, "SV352321156833"},
	{74.4964131169431, 116.71875, 12, "ZA412658306821"},
	{77.157162522661, 18.6328125, 12, "aX413323833627"},
	{82.1183836069127, -43.59375, 12, "bD023073480358"},
	{82.26169873683153, 175.078125, 12, "aA570232153327"},
	{82.29949965286332, -179.99139669189435, 12, "TK800808780528"},
	{84.11778923555221, -177.53045919189435, 12, "TO328348003525"},
	{84.76413141109205, 178.9539158081056, 12, "aE563763855337"},
	{85.05089183547521, 88.98101806640625, 12, "bb337212207184"},
	{-85.02070774312593, -79.453125, 13, "BV8046354823266"},
	{-84.9901001802348, 113.5546875, 13, "CI5803560714155"},
	{-85.05397122930692, 179.6044921875, 13, "KE0482706175341"},
	{-85.05302392713516, 179.6044921875, 13, "KE0482731186381"},
	{-85.05207644397981, 179.769287109375, 13, "KE0485035421662"},
	{-85.05207644397981, 179.813232421875, 13, "KE0485060618642"},
	{-85.05018093458115, 179.813232421875, 13, "KE0485068654767"},
	{-85.05207644397981, -179.0496826171875, 13, "DO0712537663424"},
	{-85.05112877980659, -179.044189453125, 13, "DO0712564265481"},
	{-77.50411917973987, 177.36328125, 13, "LK4501241604860"},
	{-75.97257944809174, -179.84857442626935, 13, "EU8407181150325"},
	{-78.9039293885709, -12.3046875, 13, "CZ3061211527483"},
	{-49.837982453084834, -97.3828125, 13, "Fb4625116215430"},
	{-47.279229002570816, 61.5234375, 13, "OK7250326542858"},
	{-22.593726063929296, -166.2890625, 13, "GH4221515066426"},
	{-11.523087506868512, 178.2421875, 13, "Mc4421818162227"},
	{-2.460181181020993, -179.296875, 13, "GI8724851414563"},
	{9.449061826881419, -179.296875, 13, "QU4074815410453"},
	{20.632784250388028, 179.296875, 13, "XK8810462030471"},
	{3.5134210456400443, -60.8203125, 13, "PA0414822216031"},
	{-0.3515602939922709, -3.1640625, 13, "OY4236836048713"},
	{0.3555766933805418, 6.303769323730606, 13, "OY7001555744550"},
	{31.353636941500987, 137.4609375, 13, "XM4117607703260"},
	{22.59372606392931, 129.375, 13, "PS6246856326218"},
	{29.23239505219198, 153.6084568237306, 13, "XL5887834871345"},
	{42.814468134371985, 142.0068943237306, 13, "XX0066438378463"},
	{42.5530802889558, 38.671875, 13, "PZ4537105130820"},
	{27.997947874626394, -116.7431056762694, 13, "OI8337351073052"},
	{46.31935818342445, -108.3056056762694, 13, "RX1135780670728"},
	{66.93006025862448, -80.5078125, 13, "SV4163853373183"},
	{76.9206135182968, 20.7421875, 13, "aX4034808877353"},
	{75.14077784070429, 121.2890625, 13, "ZA4726354081160"},
	{82.58610635020881, -40.4296875, 13, "bD0853382488383"},
	{81.92318632602199, 168.75, 13, "aA2756770452053"},
	{82.1664460084773, -172.6171875, 13, "TK7434850466772"},
	{83.90548439564024, -179.32123067626938, 13, "TO0846631354832"},
	{84.60821787945922, 179.62408182373065, 13, "aE4805277813220"},
	{85.051099162384, 88.67683410644531, 13, "bb3356324634667"},
	{-84.9901001802348, -79.453125, 14, "BV80717180121576"},
	{-85.14128398117636, 118.125, 14, "CI56436210413512"},
	{-85.08136441846642, 118.828125, 14, "CI56773820550701"},
	{-85.035941506574, 121.11328125, 14, "CI80611150250021"},
	{-85.05112877980659, 177.5390625, 14, "KE05167224507116"},
	{-85.035941506574, 177.71484375, 14, "KE05405867244763"},
	{-85.0513657128222, -178.33969116210938, 14, "DO07170215584802"},
	{-85.05124724772884, -178.33969116210938, 14, "DO07170226100057"},
	{-85.05101030905541, -178.33969116210938, 14, "DO07170227640163"},
	{-78.34941069014627, 178.154296875, 14, "LK16473602511456"},
	{-78.09010081233451, -178.18242910156235, 14, "EU40673078217161"},
	{-78.4905516616031, -15.46875, 14, "CZ31831222362536"},
	{-53.330872983017045, -97.03125, 14, "Fb31200035818717"},
	{-49.925654392595526, 58.59491464843764, 14, "OK44243402806508"},
	{-8.407168163601074, -2.8125, 14, "OY01571576207503"},
	{4.079240211993051, -1.170710351562363, 14, "OY56733420402656"},
	{-15.961329081596647, -159.609375, 14, "GH84065825033010"},
	{-10.621569004992997, 173.90741464843762, 14, "Mc50342481388451"},
	{-5.751376102926013, -173.43633535156232, 14, "GI78081171860100"},
	{5.48056406874262, -174.13946035156232, 14, "QU30541311610415"},
	{12.42172340890993, 179.53241464843762, 14, "XK44828780263624"},
	{5.61598581915534, -58.359375, 14, "PA07571045254517"},
	{21.162545457735018, 127.50116464843764, 14, "PS38383877824230"},
	{28.18452118500034, 145.07928964843762, 14, "XM31385143063872"},
	{29.416781643647624, 155.62616464843762, 14, "XL82447011727335"},
	{41.93183471670877, 141.56366464843762, 14, "XP22503400845082"},
	{38.27268853598097, 34.453125, 14, "PZ13526267300787"},
	{28.18452118500034, -112.96758535156236, 14, "OI78437338044154"},
	{44.49349282335009, -106.63946035156236, 14, "RU68880348017456"},
	{65.31006522139135, -84.84258535156236, 14, "SV13411223332736"},
	{77.28260220160595, 24.844914648437637, 14, "aX38226176387034"},
	{74.01954331150228, 123.046875, 14, "ZA35122807551718"},
	{82.85338229176081, -37.96875, 14, "bD32858648748733"},
	{81.5182718765338, 175.078125, 14, "aA18678460283758"},
	{81.30832090051811, -177.1875, 14, "TK35253726505866"},
	{84.11102985143594, -175.54571035156232, 14, "TO35083224064751"},
	{84.52840283662707, 177.42303964843762, 14, "aE45177603556405"},
	{85.05012909407333, 88.71932029724121, 14, "bb33563401287846"},
	{85.0511287798066, 88.7024974822998, 14, "bb33563422636071"},
	{-85.05491835052943, -80.244140625, 15, "BV801644720107680"},
	{-85.05112877980659, -80.244140625, 15, "BV801672224571144"},
	{-85.04354094565655, -76.81640625, 15, "BV752066087034602"},
	{-84.92832092949963, 113.203125, 15, "CI584051202811030"},
	{-85.05112877980659, 178.70361328125, 15, "KE053702425073344"},
	{-85.05101030905541, 178.6212158203125, 15, "KE053474671871414"},
	{-85.05112877980659, -179.47265625, 15, "DO048672464801367"},
	{-85.04923290826916, -179.49462890625, 15, "DO048680771383421"},
	{-78.63000556774834, -16.875, 15, "CZ312460160330640"},
	{-76.9999351181161, 178.59375, 15, "LK482446631807432"},
	{-77.69287033641926, -179.6484375, 15, "EU443120677000135"},
	{-48.224672649565186, 61.171875, 15, "OK486263022716030"},
	{-47.75409797968002, -101.6015625, 15, "Fb563137601173436"},
	{-24.5271348225978, -160.3125, 15, "GH405610617770460"},
	{-10.627811586244796, 175.58411714782733, 15, "Mc428447007470857"},
	{-8.895519360279467, -174.9236953521727, 15, "GI714676472576137"},
	{9.655313611088724, -177.7361953521727, 15, "QU430846555413857"},
	{13.453737213419247, 177.1875, 15, "XK455868523435074"},
	{28.613459424004414, 144.4921875, 15, "XM314886135481601"},
	{22.268764039073965, 129.0234375, 15, "PS624076167125758"},
	{30.401661513177608, 155.41321871032733, 15, "XL828827023882162"},
	{42.77151327199906, 142.75696871032733, 15, "XP252183615213788"},
	{-0.3515602939922709, -0.3515625, 15, "OY441655876848342"},
	{0.6485910937534255, -2.438343789672672, 15, "OY450583438341310"},
	{1.4061088354351594, -59.4140625, 15, "PA007023204648554"},
	{27.010563124730247, -115.28990628967267, 15, "OI754343508378851"},
	{47.989921667414166, -108.6328125, 15, "RX153284780083552"},
	{44.33956524809713, 32.6953125, 15, "PZ541437131806627"},
	{65.07213008560697, -80.859375, 15, "SV081012647253500"},
	{73.82482034613932, 116.71875, 15, "ZA162768706170274"},
	{78.27820145542813, 16.171875, 15, "aX538661552510561"},
	{83.40004205976699, -38.671875, 15, "bD475537206875044"},
	{82.30889251821553, 176.1328125, 15, "aA573170435224014"},
	{82.07002819448266, -177.890625, 15, "TK720137660817775"},
	{83.8992664793694, -177.8680312896727, 15, "TO078855632751174"},
	{84.50259500130507, -179.9774062896727, 15, "TO448407012467243"},
	{84.8609006021274, 178.96790621032733, 15, "aE801157747467437"},
	{85.05112507763846, 89.37952995300293, 15, "bb337184418811744"},
}