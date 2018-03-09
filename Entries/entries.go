package fukushu

import "fmt"

type Entry struct {
  Ji          string
  Reading     string
  Translation string
}

func (e Entry) ToString() string {
  return fmt.Sprintf(`<b>%s</b>
  [<i>%s</i>]

  %s`, e.Ji, e.Reading, e.Translation)
}

var Entries = []Entry{
  Entry{
    Ji:          "党",
    Reading:     "トウ・なかま・むら",
    Translation: "party; faction; clique",
  },
  Entry{
    Ji:          "協",
    Reading:     "キョウ",
    Translation: "co-; cooperation",
  },
  Entry{
    Ji:          "総",
    Reading:     "ソウ・す・べて・すべ・て・ふさ",
    Translation: "general; whole; all; full; total",
  },
  Entry{
    Ji:          "区",
    Reading:     "ク・オウ・コウ",
    Translation: "ward; district",
  },
  Entry{
    Ji:          "領",
    Reading:     "リョウ・えり",
    Translation: "jurisdiction; dominion; territory; fief; reign",
  },
  Entry{
    Ji:          "県",
    Reading:     "ケン・か・ける",
    Translation: "prefecture",
  },
  Entry{
    Ji:          "設",
    Reading:     "セツ・もう・ける",
    Translation: "establishment; provision; prepare",
  },
  Entry{
    Ji:          "改",
    Reading:     "カイ・あらた・める・あらた・まる",
    Translation: "reformation; change; modify; mend; renew; examine; inspect; search",
  },
  Entry{
    Ji:          "府",
    Reading:     "フ",
    Translation: "borough; urban prefecture; govt office; representative body; storehouse",
  },
  Entry{
    Ji:          "査",
    Reading:     "サ",
    Translation: "investigate",
  },
  Entry{
    Ji:          "委",
    Reading:     "イ・ゆだ・ねる",
    Translation: "committee; entrust to; leave to; devote; discard",
  },
  Entry{
    Ji:          "軍",
    Reading:     "グン・いくさ",
    Translation: "army; force; troops; war; battle",
  },
  Entry{
    Ji:          "団",
    Reading:     "ダン・トン・かたまり・まる・い",
    Translation: "group; association",
  },
  Entry{
    Ji:          "各",
    Reading:     "カク・おのおの",
    Translation: "each; every; either",
  },
  Entry{
    Ji:          "島",
    Reading:     "トウ・しま",
    Translation: "island",
  },
  Entry{
    Ji:          "革",
    Reading:     "カク・かわ",
    Translation: "leather; skin; reform; become serious",
  },
  Entry{
    Ji:          "村",
    Reading:     "ソン・むら",
    Translation: "village; town",
  },
  Entry{
    Ji:          "勢",
    Reading:     "セイ・ゼイ・いきお・い・はずみ",
    Translation: "forces; energy; military strength",
  },
  Entry{
    Ji:          "減",
    Reading:     "ゲン・へ・る・へ・らす",
    Translation: "dwindle; decrease; reduce; decline; curtail; get hungry",
  },
  Entry{
    Ji:          "再",
    Reading:     "サイ・サ・ふたた・び",
    Translation: "again; twice; second time",
  },
  Entry{
    Ji:          "税",
    Reading:     "ゼイ",
    Translation: "tax; duty",
  },
  Entry{
    Ji:          "営",
    Reading:     "エイ・いとな・む・いとな・み",
    Translation: "occupation; camp; perform; build; conduct (business)",
  },
  Entry{
    Ji:          "比",
    Reading:     "ヒ・くら・べる",
    Translation: "compare; race; ratio; Philipines",
  },
  Entry{
    Ji:          "防",
    Reading:     "ボウ・ふせ・ぐ",
    Translation: "ward off; defend; protect; resist",
  },
  Entry{
    Ji:          "補",
    Reading:     "ホ・おぎな・う",
    Translation: "supplement; supply; make good; offset; compensate; assistant; learner",
  },
  Entry{
    Ji:          "境",
    Reading:     "キョウ・ケイ・さかい",
    Translation: "boundary; border; region",
  },
  Entry{
    Ji:          "導",
    Reading:     "ドウ・みちび・く",
    Translation: "guidance; leading; conduct; usher",
  },
  Entry{
    Ji:          "副",
    Reading:     "フク",
    Translation: "vice-; assistant; aide; duplicate; copy",
  },
  Entry{
    Ji:          "算",
    Reading:     "サン・そろ",
    Translation: "calculate; divining; number; abacus; probability",
  },
  Entry{
    Ji:          "輸",
    Reading:     "ユ・シュ",
    Translation: "transport; send; be inferior",
  },
  Entry{
    Ji:          "述",
    Reading:     "ジュツ・の・べる",
    Translation: "mention; state; speak; relate",
  },
  Entry{
    Ji:          "線",
    Reading:     "セン・すじ",
    Translation: "line; track",
  },
  Entry{
    Ji:          "農",
    Reading:     "ノウ",
    Translation: "agriculture; farmers",
  },
  Entry{
    Ji:          "州",
    Reading:     "シュウ・ス・す",
    Translation: "state; province",
  },
  Entry{
    Ji:          "武",
    Reading:     "ブ・ム・たけ・し",
    Translation: "warrior; military; chivalry; arms",
  },
  Entry{
    Ji:          "象",
    Reading:     "ショウ・ゾウ・かたど・る",
    Translation: "elephant; pattern after; imitate; image; shape; sign (of the times",
  },
  Entry{
    Ji:          "域",
    Reading:     "イキ",
    Translation: "range; region; limits; stage; level",
  },
  Entry{
    Ji:          "額",
    Reading:     "ガク・ひたい",
    Translation: "forehead; tablet; plaque; framed picture; sum; amount; volume",
  },
  Entry{
    Ji:          "欧",
    Reading:     "オウ・うた・う・は・く",
    Translation: "Europe",
  },
  Entry{
    Ji:          "担",
    Reading:     "タン・かつ・ぐ・にな・う",
    Translation: "shouldering; carry; raise; bear",
  },
  Entry{
    Ji:          "準",
    Reading:     "ジュン・じゅん・じる・じゅん・ずる・なぞら・える・のり・ひと・しい・みずもり",
    Translation: "semi-; correspond to; proportionate to; conform; imitate",
  },
  Entry{
    Ji:          "賞",
    Reading:     "ショウ・ほ・める",
    Translation: "prize; reward; praise",
  },
  Entry{
    Ji:          "辺",
    Reading:     "ヘン・あた・り・ほと・り・-べ",
    Translation: "environs; boundary; border; vicinity",
  },
  Entry{
    Ji:          "造",
    Reading:     "ゾウ・つく・る・つく・り・-づく・り",
    Translation: "create; make; structure; physique",
  },
  Entry{
    Ji:          "被",
    Reading:     "ヒ・こうむ・る・おお・う・かぶ・る・かぶ・せる",
    Translation: "incur; cover; veil; brood over; shelter; wear; put on; be exposed (film); receiving",
  },
  Entry{
    Ji:          "技",
    Reading:     "ギ・わざ",
    Translation: "skill; art; craft; ability; feat; performance; vocation; arts",
  },
  Entry{
    Ji:          "低",
    Reading:     "テイ・ひく・い・ひく・める・ひく・まる",
    Translation: "lower; short; humble",
  },
  Entry{
    Ji:          "復",
    Reading:     "フク・また",
    Translation: "restore; return to; revert; resume",
  },
  Entry{
    Ji:          "移",
    Reading:     "イ・うつ・る・うつ・す",
    Translation: "shift; move; change; drift; catch (cold, fire); pass into",
  },
  Entry{
    Ji:          "個",
    Reading:     "コ・カ",
    Translation: "individual; counter for articles",
  },
  Entry{
    Ji:          "門",
    Reading:     "モン・かど・と",
    Translation: "gate; counter for cannons",
  },
  Entry{
    Ji:          "課",
    Reading:     "カ",
    Translation: "chapter; lesson; section; department; division; counter for chapters (of a book)",
  },
  Entry{
    Ji:          "脳",
    Reading:     "ノウ・ドウ・のうずる",
    Translation: "brain; memory",
  },
  Entry{
    Ji:          "極",
    Reading:     "キョク・ゴク・きわ・める・きわ・まる・きわ・まり・きわ・み・き・める・-ぎ・め・き・まる",
    Translation: "poles; settlement; conclusion; end; highest rank; electric poles; very; extremely; most; highly; 10**48",
  },
  Entry{
    Ji:          "含",
    Reading:     "ガン・ふく・む・ふく・める",
    Translation: "include; bear in mind; understand; cherish",
  },
  Entry{
    Ji:          "蔵",
    Reading:     "ゾウ・ソウ・くら・おさ・める・かく・れる",
    Translation: "storehouse; hide; own; have; possess",
  },
  Entry{
    Ji:          "量",
    Reading:     "リョウ・はか・る",
    Translation: "quantity; measure; weight; amount; consider; estimate; surmise",
  },
  Entry{
    Ji:          "型",
    Reading:     "ケイ・かた・-がた",
    Translation: "mould; type; model",
  },
  Entry{
    Ji:          "況",
    Reading:     "キョウ・まし・て・いわ・んや・おもむき",
    Translation: "condition; situation",
  },
  Entry{
    Ji:          "針",
    Reading:     "シン・はり",
    Translation: "needle; pin; staple; stinger",
  },
  Entry{
    Ji:          "専",
    Reading:     "セン・もっぱ・ら",
    Translation: "specialty; exclusive; mainly; solely",
  },
  Entry{
    Ji:          "谷",
    Reading:     "コク・たに・きわ・まる",
    Translation: "valley",
  },
  Entry{
    Ji:          "史",
    Reading:     "シ",
    Translation: "history; chronicle",
  },
  Entry{
    Ji:          "階",
    Reading:     "カイ・きざはし",
    Translation: "storey; stair; counter for storeys of a building",
  },
  Entry{
    Ji:          "管",
    Reading:     "カン・くだ",
    Translation: "pipe; tube; wind instrument; drunken talk; control; jurisdiction",
  },
  Entry{
    Ji:          "兵",
    Reading:     "ヘイ・ヒョウ・つわもの",
    Translation: "soldier; private; troops; army; warfare; strategy; tactics",
  },
  Entry{
    Ji:          "接",
    Reading:     "セツ・ショウ・つ・ぐ",
    Translation: "touch; contact; adjoin; piece together",
  },
  Entry{
    Ji:          "細",
    Reading:     "サイ・ほそ・い・ほそ・る・こま・か・こま・かい",
    Translation: "dainty; get thin; taper; slender; narrow",
  },
  Entry{
    Ji:          "効",
    Reading:     "コウ・き・く・ききめ・なら・う",
    Translation: "merit; efficacy; efficiency; benefit",
  },
  Entry{
    Ji:          "丸",
    Reading:     "ガン・まる・まる・める・まる・い",
    Translation: "round; full (month); perfection; -ship; pills; make round; roll up; curl up; seduce; explain away",
  },
  Entry{
    Ji:          "湾",
    Reading:     "ワン・いりえ",
    Translation: "gulf; bay; inlet",
  },
  Entry{
    Ji:          "録",
    Reading:     "ロク・しる・す・と・る",
    Translation: "record",
  },
  Entry{
    Ji:          "省",
    Reading:     "セイ・ショウ・かえり・みる・はぶ・く",
    Translation: "focus; government ministry; conserve",
  },
  Entry{
    Ji:          "旧",
    Reading:     "キュウ・ふる・い・もと",
    Translation: "old times; old things; old friend; former; ex",
  },
  Entry{
    Ji:          "橋",
    Reading:     "キョウ・はし",
    Translation: "bridge",
  },
  Entry{
    Ji:          "岸",
    Reading:     "ガン・きし",
    Translation: "beach",
  },
  Entry{
    Ji:          "周",
    Reading:     "シュウ・まわ・り",
    Translation: "circumference; circuit; lap",
  },
  Entry{
    Ji:          "材",
    Reading:     "ザイ",
    Translation: "lumber; log; timber; wood; materials; ingredients; talent",
  },
  Entry{
    Ji:          "戸",
    Reading:     "コ・と",
    Translation: "door; counter for houses; door radical (no・ 63",
  },
  Entry{
    Ji:          "央",
    Reading:     "オウ",
    Translation: "center; middle",
  },
  Entry{
    Ji:          "券",
    Reading:     "ケン",
    Translation: "ticket",
  },
  Entry{
    Ji:          "編",
    Reading:     "ヘン・あ・む・-あ・み",
    Translation: "compilation; knit; plait; braid; twist; editing; completed poem; part of a book",
  },
  Entry{
    Ji:          "捜",
    Reading:     "ソウ・シュ・シュウ・さが・す",
    Translation: "search; look for; locate",
  },
  Entry{
    Ji:          "竹",
    Reading:     "チク・たけ",
    Translation: "bamboo",
  },
  Entry{
    Ji:          "超",
    Reading:     "チョウ・こ・える・こ・す",
    Translation: "transcend; super-; ultra",
  },
  Entry{
    Ji:          "並",
    Reading:     "ヘイ・ホウ・な・み・なみ・なら・べる・なら・ぶ・なら・びに",
    Translation: "row; and; besides; as well as; line up; rank with; rival; equal",
  },
  Entry{
    Ji:          "療",
    Reading:     "リョウ",
    Translation: "heal; cure",
  },
  Entry{
    Ji:          "採",
    Reading:     "サイ・と・る",
    Translation: "pick; take; fetch; take up",
  },
  Entry{
    Ji:          "森",
    Reading:     "シン・もり",
    Translation: "forest; woods",
  },
  Entry{
    Ji:          "競",
    Reading:     "キョウ・ケイ・きそ・う・せ・る・くら・べる",
    Translation: "emulate; compete with; bid; sell at auction; bout; contest; race",
  },
  Entry{
    Ji:          "介",
    Reading:     "カイ",
    Translation: "jammed in; shellfish; mediate; concern oneself with",
  },
  Entry{
    Ji:          "根",
    Reading:     "コン・ね・-ね",
    Translation: "root; radical; head (pimple",
  },
  Entry{
    Ji:          "販",
    Reading:     "ハン",
    Translation: "marketing; sell; trade",
  },
  Entry{
    Ji:          "歴",
    Reading:     "レキ・レッキ",
    Translation: "curriculum; continuation; passage of time",
  },
  Entry{
    Ji:          "将",
    Reading:     "ショウ・ソウ・まさ・に・はた・まさ・ひきい・る・もって",
    Translation: "leader; commander; general; admiral; or; and again; soon; from now on; just about",
  },
  Entry{
    Ji:          "幅",
    Reading:     "フク・はば",
    Translation: "hanging scroll; width",
  },
  Entry{
    Ji:          "般",
    Reading:     "ハン",
    Translation: "carrier; carry; all",
  },
  Entry{
    Ji:          "貿",
    Reading:     "ボウ",
    Translation: "trade; exchange",
  },
  Entry{
    Ji:          "講",
    Reading:     "コウ",
    Translation: "lecture; club; association",
  },
  Entry{
    Ji:          "林",
    Reading:     "リン・はやし",
    Translation: "grove; forest",
  },
  Entry{
    Ji:          "装",
    Reading:     "ソウ・ショウ・よそお・う・よそお・い",
    Translation: "attire; dress; pretend; disguise; profess",
  },
  Entry{
    Ji:          "諸",
    Reading:     "ショ・もろ",
    Translation: "various; many; several; together",
  },
  Entry{
    Ji:          "劇",
    Reading:     "ゲキ",
    Translation: "drama; play",
  },
  Entry{
    Ji:          "河",
    Reading:     "カ・かわ",
    Translation: "river",
  },
  Entry{
    Ji:          "航",
    Reading:     "コウ",
    Translation: "navigate; sail; cruise; fly",
  },
  Entry{
    Ji:          "鉄",
    Reading:     "テツ・くろがね",
    Translation: "iron",
  },
  Entry{
    Ji:          "児",
    Reading:     "ジ・ニ・ゲイ・こ・-こ・-っこ",
    Translation: "newborn babe; child; young of animals",
  },
  Entry{
    Ji:          "禁",
    Reading:     "キン",
    Translation: "prohibition; ban; forbid",
  },
  Entry{
    Ji:          "印",
    Reading:     "イン・しるし・-じるし・しる・す",
    Translation: "stamp; seal; mark; imprint; symbol; emblem; trademark; evidence; souvenir; India",
  },
  Entry{
    Ji:          "逆",
    Reading:     "ギャク・ゲキ・さか・さか・さ・さか・らう",
    Translation: "inverted; reverse; opposite; wicked",
  },
  Entry{
    Ji:          "換",
    Reading:     "カン・か・える・-か・える・か・わる",
    Translation: "interchange; period; change; convert; replace; renew",
  },
  Entry{
    Ji:          "久",
    Reading:     "キュウ・ク・ひさ・しい",
    Translation: "long time; old story",
  },
  Entry{
    Ji:          "短",
    Reading:     "タン・みじか・い",
    Translation: "short; brevity; fault; defect; weak point",
  },
  Entry{
    Ji:          "油",
    Reading:     "ユ・ユウ・あぶら",
    Translation: "oil; fat",
  },
  Entry{
    Ji:          "暴",
    Reading:     "ボウ・バク・あば・く・あば・れる",
    Translation: "outburst; rave; fret; force; violence; cruelty; outrage",
  },
  Entry{
    Ji:          "輪",
    Reading:     "リン・わ",
    Translation: "wheel; ring; circle; link; loop; counter for wheels and flowers",
  },
  Entry{
    Ji:          "占",
    Reading:     "セン・し・める・うらな・う",
    Translation: "fortune-telling; divining; forecasting; occupy; hold; have; get; take",
  },
  Entry{
    Ji:          "植",
    Reading:     "ショク・う・える・う・わる",
    Translation: "plant",
  },
  Entry{
    Ji:          "清",
    Reading:     "セイ・ショウ・シン・きよ・い・きよ・まる・きよ・める",
    Translation: "pure; purify; cleanse; exorcise; Manchu dynasty",
  },
  Entry{
    Ji:          "倍",
    Reading:     "バイ",
    Translation: "double; twice; times; fold",
  },
  Entry{
    Ji:          "均",
    Reading:     "キン・なら・す",
    Translation: "level; average",
  },
  Entry{
    Ji:          "億",
    Reading:     "オク",
    Translation: "hundred million; 10**8",
  },
  Entry{
    Ji:          "圧",
    Reading:     "アツ・エン・オウ・お・す・へ・す・おさ・える・お・さえる",
    Translation: "pressure; push; overwhelm; oppress; dominate",
  },
  Entry{
    Ji:          "芸",
    Reading:     "ゲイ・ウン・う・える・のり・わざ",
    Translation: "technique; art; craft; performance; acting; trick; stunt",
  },
  Entry{
    Ji:          "署",
    Reading:     "ショ",
    Translation: "signature; govt office; police station",
  },
  Entry{
    Ji:          "伸",
    Reading:     "シン・の・びる・の・ばす・の・べる・の・す",
    Translation: "expand; stretch; extend; lengthen; increase",
  },
  Entry{
    Ji:          "停",
    Reading:     "テイ・と・める・と・まる",
    Translation: "halt; stopping",
  },
  Entry{
    Ji:          "爆",
    Reading:     "バク・は・ぜる",
    Translation: "bomb; burst open; pop; split",
  },
  Entry{
    Ji:          "陸",
    Reading:     "リク・ロク・おか",
    Translation: "land; six",
  },
  Entry{
    Ji:          "玉",
    Reading:     "ギョク・たま・たま-・-だま",
    Translation: "jewel; ball",
  },
  Entry{
    Ji:          "波",
    Reading:     "ハ・なみ",
    Translation: "waves; billows; Poland",
  },
  Entry{
    Ji:          "帯",
    Reading:     "タイ・お・びる・おび",
    Translation: "sash; belt; obi; zone; region",
  },
  Entry{
    Ji:          "延",
    Reading:     "エン・の・びる・の・べる・の・べ・の・ばす",
    Translation: "prolong; stretching",
  },
  Entry{
    Ji:          "羽",
    Reading:     "ウ・は・わ・はね",
    Translation: "feathers; counter for birds, rabbits",
  },
  Entry{
    Ji:          "固",
    Reading:     "コ・かた・める・かた・まる・かた・まり・かた・い",
    Translation: "harden; set; clot; curdle",
  },
  Entry{
    Ji:          "則",
    Reading:     "ソク・のっと・る",
    Translation: "rule; follow; based on; model after",
  },
  Entry{
    Ji:          "乱",
    Reading:     "ラン・ロン・みだ・れる・みだ・る・みだ・す・みだ・おさ・める・わた・る",
    Translation: "riot; war; disorder; disturb",
  },
  Entry{
    Ji:          "普",
    Reading:     "フ・あまね・く・あまねし",
    Translation: "universal; wide(ly); generally; Prussia",
  },
  Entry{
    Ji:          "測",
    Reading:     "ソク・はか・る",
    Translation: "fathom; plan; scheme; measure",
  },
  Entry{
    Ji:          "豊",
    Reading:     "ホウ・ブ・ゆた・か・とよ",
    Translation: "bountiful; excellent; rich",
  },
  Entry{
    Ji:          "厚",
    Reading:     "コウ・あつ・い・あか",
    Translation: "thick; heavy; rich; kind; cordial; brazen; shameless",
  },
  Entry{
    Ji:          "齢",
    Reading:     "レイ・よわい・とし",
    Translation: "age",
  },
  Entry{
    Ji:          "囲",
    Reading:     "イ・かこ・む・かこ・う・かこ・い",
    Translation: "surround; besiege; store; paling; enclosure; encircle; preserve; keep",
  },
  Entry{
    Ji:          "卒",
    Reading:     "ソツ・シュツ・そっ・する・お・える・お・わる・ついに・にわか",
    Translation: "graduate; soldier; private; die",
  },
  Entry{
    Ji:          "略",
    Reading:     "リャク・ほぼ・おか・す・おさ・める・はかりごと・はか・る・はぶ・く・りゃく・す・りゃく・する",
    Translation: "abbreviation; omission; outline; shorten; capture; plunder",
  },
  Entry{
    Ji:          "承",
    Reading:     "ショウ・ジョウ・うけたまわ・る・う・ける・ささ・げる・とど・める・たす・ける・こ・らす・つい・で・すく・う",
    Translation: "acquiesce; hear; listen to; be informed; receive",
  },
  Entry{
    Ji:          "順",
    Reading:     "ジュン",
    Translation: "obey; order; turn; right; docility; occasion",
  },
  Entry{
    Ji:          "岩",
    Reading:     "ガン・いわ",
    Translation: "boulder; rock; cliff",
  },
  Entry{
    Ji:          "練",
    Reading:     "レン・ね・る・ね・り",
    Translation: "practice; gloss; train; drill; polish; refine",
  },
  Entry{
    Ji:          "軽",
    Reading:     "ケイ・かる・い・かろ・やか・かろ・んじる",
    Translation: "lightly; trifling; unimportant",
  },
  Entry{
    Ji:          "了",
    Reading:     "リョウ",
    Translation: "complete; finish",
  },
  Entry{
    Ji:          "庁",
    Reading:     "チョウ・テイ・やくしょ",
    Translation: "government office",
  },
  Entry{
    Ji:          "城",
    Reading:     "ジョウ・しろ",
    Translation: "castle",
  },
  Entry{
    Ji:          "患",
    Reading:     "カン・わずら・う",
    Translation: "afflicted; disease; suffer from; be ill",
  },
  Entry{
    Ji:          "層",
    Reading:     "ソウ",
    Translation: "stratum; social class; layer; story; floor",
  },
  Entry{
    Ji:          "版",
    Reading:     "ハン",
    Translation: "printing block; printing plate; edition; impression; label",
  },
  Entry{
    Ji:          "令",
    Reading:     "レイ",
    Translation: "orders; ancient laws; command; decree",
  },
  Entry{
    Ji:          "角",
    Reading:     "カク・かど・つの",
    Translation: "angle; corner; square; horn; antlers",
  },
  Entry{
    Ji:          "絡",
    Reading:     "ラク・から・む・から・まる",
    Translation: "entwine; coil around; get caught in",
  },
  Entry{
    Ji:          "損",
    Reading:     "ソン・そこ・なう・そこな・う・-そこ・なう・そこ・ねる・-そこ・ねる",
    Translation: "damage; loss; disadvantage; hurt; injure",
  },
  Entry{
    Ji:          "募",
    Reading:     "ボ・つの・る",
    Translation: "recruit; campaign; gather (contributions); enlist; grow violent",
  },
  Entry{
    Ji:          "裏",
    Reading:     "リ・うら",
    Translation: "back; amidst; in; reverse; inside; palm; sole; rear; lining; wrong side",
  },
  Entry{
    Ji:          "仏",
    Reading:     "ブツ・フツ・ほとけ",
    Translation: "Buddha; the dead; France",
  },
  Entry{
    Ji:          "績",
    Reading:     "セキ",
    Translation: "exploits; unreeling cocoons",
  },
  Entry{
    Ji:          "築",
    Reading:     "チク・きず・く",
    Translation: "fabricate; build; construct",
  },
  Entry{
    Ji:          "貨",
    Reading:     "カ・たから",
    Translation: "freight; goods; property",
  },
  Entry{
    Ji:          "混",
    Reading:     "コン・ま・じる・-ま・じり・ま・ざる・ま・ぜる・こ・む",
    Translation: "mix; blend; confuse",
  },
  Entry{
    Ji:          "昇",
    Reading:     "ショウ・のぼ・る",
    Translation: "rise up",
  },
  Entry{
    Ji:          "池",
    Reading:     "チ・いけ",
    Translation: "pond; cistern; pool; reservoir",
  },
  Entry{
    Ji:          "血",
    Reading:     "ケツ・ち",
    Translation: "blood",
  },
  Entry{
    Ji:          "温",
    Reading:     "オン・あたた・か・あたた・かい・あたた・まる・あたた・める・ぬく",
    Translation: "warm",
  },
  Entry{
    Ji:          "季",
    Reading:     "キ",
    Translation: "seasons",
  },
  Entry{
    Ji:          "星",
    Reading:     "セイ・ショウ・ほし・-ぼし",
    Translation: "star; spot; dot; mark",
  },
  Entry{
    Ji:          "永",
    Reading:     "エイ・なが・い",
    Translation: "eternity; long; lengthy",
  },
  Entry{
    Ji:          "著",
    Reading:     "チョ・チャク・あらわ・す・いちじる・しい",
    Translation: "renowned; publish; write; remarkable; phenomenal; put on; don; wear; arrival; finish (race); counter for suits of clothing; literary work",
  },
  Entry{
    Ji:          "誌",
    Reading:     "シ",
    Translation: "document; records",
  },
  Entry{
    Ji:          "庫",
    Reading:     "コ・ク・くら",
    Translation: "warehouse; storehouse",
  },
  Entry{
    Ji:          "刊",
    Reading:     "カン",
    Translation: "publish; carve; engrave",
  },
  Entry{
    Ji:          "像",
    Reading:     "ゾウ",
    Translation: "statue; picture; image; figure; portrait",
  },
  Entry{
    Ji:          "香",
    Reading:     "コウ・キョウ・か・かお・り・かお・る",
    Translation: "incense; smell; perfume",
  },
  Entry{
    Ji:          "坂",
    Reading:     "ハン・さか",
    Translation: "slope; incline; hill",
  },
  Entry{
    Ji:          "底",
    Reading:     "テイ・そこ",
    Translation: "bottom; sole; depth; bottom price; base; kind; sort",
  },
  Entry{
    Ji:          "布",
    Reading:     "フ・ぬの",
    Translation: "linen; cloth",
  },
  Entry{
    Ji:          "寺",
    Reading:     "ジ・てら",
    Translation: "Buddhist temple",
  },
  Entry{
    Ji:          "宇",
    Reading:     "ウ",
    Translation: "leaves; roof; house; heaven",
  },
  Entry{
    Ji:          "巨",
    Reading:     "キョ",
    Translation: "gigantic; big; large; great",
  },
  Entry{
    Ji:          "震",
    Reading:     "シン・ふる・う・ふる・える",
    Translation: "quake; shake; tremble; quiver; shiver",
  },
  Entry{
    Ji:          "希",
    Reading:     "キ・ケ・まれ",
    Translation: "hope; beg; request; pray; beseech; Greece; dilute (acid); rare; few; phenomenal",
  },
  Entry{
    Ji:          "触",
    Reading:     "ショク・ふ・れる・さわ・る・さわ",
    Translation: "contact; touch; feel; hit; proclaim; announce; conflict",
  },
  Entry{
    Ji:          "依",
    Reading:     "イ・エ・よ・る",
    Translation: "reliant; depend on; consequently; therefore; due to",
  },
  Entry{
    Ji:          "籍",
    Reading:     "セキ",
    Translation: "enroll; domiciliary register; membership",
  },
  Entry{
    Ji:          "汚",
    Reading:     "オ・けが・す・けが・れる・けが・らわしい・よご・す・よご・れる・きたな・い",
    Translation: "dirty; pollute; disgrace; rape; defile",
  },
  Entry{
    Ji:          "枚",
    Reading:     "マイ・バイ",
    Translation: "sheet of.; counter for flat thin objects or sheets",
  },
  Entry{
    Ji:          "複",
    Reading:     "フク",
    Translation: "duplicate; double; compound; multiple",
  },
  Entry{
    Ji:          "郵",
    Reading:     "ユウ",
    Translation: "mail; stagecoach stop",
  },
  Entry{
    Ji:          "仲",
    Reading:     "チュウ・なか",
    Translation: "go-between; relationship",
  },
  Entry{
    Ji:          "栄",
    Reading:     "エイ・ヨウ・さか・える・は・え・-ば・え・は・える・え",
    Translation: "flourish; prosperity; honor; glory; splendor",
  },
  Entry{
    Ji:          "札",
    Reading:     "サツ・ふだ",
    Translation: "tag; paper money; counter for bonds; placard; bid",
  },
  Entry{
    Ji:          "板",
    Reading:     "ハン・バン・いた",
    Translation: "plank; board; plate; stage",
  },
  Entry{
    Ji:          "骨",
    Reading:     "コツ・ほね",
    Translation: "skeleton; bone; remains; frame",
  },
  Entry{
    Ji:          "傾",
    Reading:     "ケイ・かたむ・く・かたむ・ける・かたぶ・く・かた・げる・かし・げる",
    Translation: "lean; incline; tilt; trend; wane; sink; ruin; bias",
  },
  Entry{
    Ji:          "届",
    Reading:     "カイ・とど・ける・-とど・け・とど・く",
    Translation: "deliver; reach; arrive; report; notify; forward",
  },
  Entry{
    Ji:          "巻",
    Reading:     "カン・ケン・ま・く・まき・ま・き",
    Translation: "scroll; volume; book; part; roll up; wind up; tie; coil; counter for texts (or book scrolls",
  },
  Entry{
    Ji:          "燃",
    Reading:     "ネン・も・える・も・やす・も・す",
    Translation: "burn; blaze; glow",
  },
  Entry{
    Ji:          "跡",
    Reading:     "セキ・あと",
    Translation: "tracks; mark; print; impression",
  },
  Entry{
    Ji:          "包",
    Reading:     "ホウ・つつ・む・くる・む",
    Translation: "wrap; pack up; cover; conceal",
  },
  Entry{
    Ji:          "駐",
    Reading:     "チュウ",
    Translation: "stop-over; reside in; resident",
  },
  Entry{
    Ji:          "弱",
    Reading:     "ジャク・よわ・い・よわ・る・よわ・まる・よわ・める",
    Translation: "weak; frail",
  },
  Entry{
    Ji:          "紹",
    Reading:     "ショウ",
    Translation: "introduce; inherit; help",
  },
  Entry{
    Ji:          "雇",
    Reading:     "コ・やと・う",
    Translation: "employ; hire",
  },
  Entry{
    Ji:          "替",
    Reading:     "タイ・か・える・か・え-・か・わる",
    Translation: "exchange; spare; substitute; per",
  },
  Entry{
    Ji:          "預",
    Reading:     "ヨ・あず・ける・あず・かる",
    Translation: "deposit; custody; leave with; entrust to",
  },
  Entry{
    Ji:          "焼",
    Reading:     "ショウ・や・く・や・き・や・き-・-や・き・や・ける",
    Translation: "bake; burning",
  },
  Entry{
    Ji:          "簡",
    Reading:     "カン",
    Translation: "simplicity; brevity",
  },
  Entry{
    Ji:          "章",
    Reading:     "ショウ",
    Translation: "badge; chapter; composition; poem; design",
  },
  Entry{
    Ji:          "臓",
    Reading:     "ゾウ・はらわた",
    Translation: "entrails; viscera; bowels",
  },
  Entry{
    Ji:          "律",
    Reading:     "リツ・リチ・レツ",
    Translation: "rhythm; law; regulation; gauge; control",
  },
  Entry{
    Ji:          "贈",
    Reading:     "ゾウ・ソウ・おく・る",
    Translation: "presents; send; give to; award to; confer on; presenting something",
  },
  Entry{
    Ji:          "照",
    Reading:     "ショウ・て・る・て・らす・て・れる",
    Translation: "illuminate; shine; compare; bashful",
  },
  Entry{
    Ji:          "薄",
    Reading:     "ハク・うす・い・うす-・-うす・うす・める・うす・まる・うす・らぐ・うす・ら-・うす・れる・すすき",
    Translation: "dilute; thin; weak (tea",
  },
  Entry{
    Ji:          "群",
    Reading:     "グン・む・れる・む・れ・むら・むら・がる",
    Translation: "flock; group; crowd; herd; swarm; cluster",
  },
  Entry{
    Ji:          "秒",
    Reading:     "ビョウ",
    Translation: "second (1/60 minute",
  },
  Entry{
    Ji:          "奥",
    Reading:     "オウ・おく・おく・まる・くま",
    Translation: "heart; interior",
  },
  Entry{
    Ji:          "詰",
    Reading:     "キツ・キチ・つ・める・つ・め・-づ・め・つ・まる・つ・む",
    Translation: "packed; close; pressed; reprove; rebuke; blame",
  },
  Entry{
    Ji:          "双",
    Reading:     "ソウ・ふた・たぐい・ならぶ・ふたつ",
    Translation: "pair; set; comparison; counter for pairs",
  },
  Entry{
    Ji:          "刺",
    Reading:     "シ・さ・す・さ・さる・さ・し・さし・とげ",
    Translation: "thorn; pierce; stab; prick; sting; calling card",
  },
  Entry{
    Ji:          "純",
    Reading:     "ジュン",
    Translation: "genuine; purity; innocence; net (profit)",
  },
  Entry{
    Ji:          "翌",
    Reading:     "ヨク",
    Translation: "the following; next",
  },
  Entry{
    Ji:          "快",
    Reading:     "カイ・こころよ・い",
    Translation: "cheerful; pleasant; agreeable; comfortable",
  },
  Entry{
    Ji:          "片",
    Reading:     "ヘン・かた-・かた",
    Translation: "one-sided; leaf; sheet; right-side kata radical (no・ 91",
  },
  Entry{
    Ji:          "敬",
    Reading:     "ケイ・キョウ・うやま・う",
    Translation: "awe; respect; honor; revere",
  },
  Entry{
    Ji:          "悩",
    Reading:     "ノウ・なや・む・なや・ます・なや・ましい・なやみ",
    Translation: "trouble; worry; in pain; distress; illness",
  },
  Entry{
    Ji:          "泉",
    Reading:     "セン・いずみ",
    Translation: "spring; fountain",
  },
  Entry{
    Ji:          "皮",
    Reading:     "ヒ・かわ",
    Translation: "pelt; skin; hide; leather; skin radical (no・ 107",
  },
  Entry{
    Ji:          "漁",
    Reading:     "ギョ・リョウ・あさ・る",
    Translation: "fishing; fishery",
  },
  Entry{
    Ji:          "荒",
    Reading:     "コウ・あら・い・あら-・あ・れる・あ・らす・-あ・らし・すさ・む",
    Translation: "laid waste; rough; rude; wild",
  },
  Entry{
    Ji:          "貯",
    Reading:     "チョ・た・める・たくわ・える",
    Translation: "savings; store; lay in; keep; wear mustache",
  },
  Entry{
    Ji:          "硬",
    Reading:     "コウ・かた・い",
    Translation: "stiff; hard",
  },
  Entry{
    Ji:          "埋",
    Reading:     "マイ・う・める・う・まる・う・もれる・うず・める・うず・まる・い・ける",
    Translation: "bury; be filled up; embedded",
  },
  Entry{
    Ji:          "柱",
    Reading:     "チュウ・はしら",
    Translation: "pillar; post; cylinder; support",
  },
  Entry{
    Ji:          "祭",
    Reading:     "サイ・まつ・る・まつ・り・まつり",
    Translation: "ritual; offer prayers; celebrate; deify; enshrine; worship",
  },
  Entry{
    Ji:          "袋",
    Reading:     "タイ・ダイ・ふくろ",
    Translation: "sack; bag; pouch",
  },
  Entry{
    Ji:          "筆",
    Reading:     "ヒツ・ふで",
    Translation: "writing brush; writing; painting brush; handwriting",
  },
  Entry{
    Ji:          "訓",
    Reading:     "クン・キン・おし・える・よ・む・くん・ずる",
    Translation: "instruction; Japanese character reading; explanation; read",
  },
  Entry{
    Ji:          "浴",
    Reading:     "ヨク・あ・びる・あ・びせる",
    Translation: "bathe; be favored with; bask in",
  },
  Entry{
    Ji:          "童",
    Reading:     "ドウ・わらべ",
    Translation: "juvenile; child",
  },
  Entry{
    Ji:          "宝",
    Reading:     "ホウ・たから",
    Translation: "treasure; wealth; valuables",
  },
  Entry{
    Ji:          "封",
    Reading:     "フウ・ホウ",
    Translation: "seal; closing",
  },
  Entry{
    Ji:          "胸",
    Reading:     "キョウ・むね・むな-",
    Translation: "bosom; breast; chest; heart; feelings",
  },
  Entry{
    Ji:          "砂",
    Reading:     "サ・シャ・すな",
    Translation: "sand",
  },
  Entry{
    Ji:          "塩",
    Reading:     "エン・しお",
    Translation: "salt",
  },
  Entry{
    Ji:          "賢",
    Reading:     "ケン・かしこ・い",
    Translation: "intelligent; wise; wisdom; cleverness",
  },
  Entry{
    Ji:          "腕",
    Reading:     "ワン・うで",
    Translation: "arm; ability; talent",
  },
  Entry{
    Ji:          "兆",
    Reading:     "チョウ・きざ・す・きざ・し",
    Translation: "portent; 10**12; trillion; sign; omen; symptoms",
  },
  Entry{
    Ji:          "床",
    Reading:     "ショウ・とこ・ゆか",
    Translation: "bed; counter for beds; floor; padding; tatami",
  },
  Entry{
    Ji:          "毛",
    Reading:     "モウ・け",
    Translation: "fur; hair; feather; down",
  },
  Entry{
    Ji:          "緑",
    Reading:     "リョク・ロク・みどり",
    Translation: "green",
  },
  Entry{
    Ji:          "尊",
    Reading:     "ソン・たっと・い・とうと・い・たっと・ぶ・とうと・ぶ",
    Translation: "revered; valuable; precious; noble; exalted",
  },
  Entry{
    Ji:          "祝",
    Reading:     "シュク・シュウ・いわ・う",
    Translation: "celebrate; congratulate",
  },
  Entry{
    Ji:          "柔",
    Reading:     "ジュウ・ニュウ・やわ・らか・やわ・らかい・やわ・やわ・ら",
    Translation: "tender; weakness; gentleness; softness",
  },
  Entry{
    Ji:          "殿",
    Reading:     "デン・テン・との・-どの",
    Translation: "Mr.; hall; mansion; palace; temple; lord",
  },
  Entry{
    Ji:          "濃",
    Reading:     "ノウ・こ・い",
    Translation: "concentrated; thick; dark; undiluted",
  },
  Entry{
    Ji:          "液",
    Reading:     "エキ",
    Translation: "fluid; liquid; juice; sap; secretion",
  },
  Entry{
    Ji:          "衣",
    Reading:     "イ・エ・ころも・きぬ・-ぎ",
    Translation: "garment; clothes; dressing",
  },
  Entry{
    Ji:          "肩",
    Reading:     "ケン・かた",
    Translation: "shoulder",
  },
  Entry{
    Ji:          "零",
    Reading:     "レイ・ぜろ・こぼ・す・こぼ・れる",
    Translation: "zero; spill; overflow; nothing; cipher",
  },
  Entry{
    Ji:          "幼",
    Reading:     "ヨウ・おさな・い",
    Translation: "infancy; childhood",
  },
  Entry{
    Ji:          "荷",
    Reading:     "カ・に",
    Translation: "baggage; shoulder-pole load; bear (a burden); shoulder (a gun); load; cargo; freight",
  },
  Entry{
    Ji:          "泊",
    Reading:     "ハク・と・まる・と・める",
    Translation: "overnight; put up at; ride at anchor; 3-day stay",
  },
  Entry{
    Ji:          "黄",
    Reading:     "コウ・オウ・き・こ-",
    Translation: "yellow",
  },
  Entry{
    Ji:          "甘",
    Reading:     "カン・あま・い・あま・える・あま・やかす・うま・い",
    Translation: "sweet; coax; pamper; be content; sugary",
  },
  Entry{
    Ji:          "臣",
    Reading:     "シン・ジン",
    Translation: "retainer; subject",
  },
  Entry{
    Ji:          "浅",
    Reading:     "セン・あさ・い",
    Translation: "shallow; superficial; frivolous; wretched; shameful",
  },
  Entry{
    Ji:          "掃",
    Reading:     "ソウ・シュ・は・く",
    Translation: "sweep; brush",
  },
  Entry{
    Ji:          "雲",
    Reading:     "ウン・くも・-ぐも",
    Translation: "cloud",
  },
  Entry{
    Ji:          "掘",
    Reading:     "クツ・ほ・る",
    Translation: "dig; delve; excavate",
  },
  Entry{
    Ji:          "捨",
    Reading:     "シャ・す・てる",
    Translation: "discard; throw away; abandon; resign; reject; sacrifice",
  },
  Entry{
    Ji:          "軟",
    Reading:     "ナン・やわ・らか・やわ・らかい",
    Translation: "soft",
  },
  Entry{
    Ji:          "沈",
    Reading:     "チン・ジン・しず・む・しず・める",
    Translation: "sink; be submerged; subside; be depressed; aloes",
  },
  Entry{
    Ji:          "凍",
    Reading:     "トウ・こお・る・こご・える・こご・る・い・てる・し・みる",
    Translation: "frozen; congeal; refrigerate",
  },
  Entry{
    Ji:          "乳",
    Reading:     "ニュウ・ちち・ち",
    Translation: "milk; breasts",
  },
  Entry{
    Ji:          "恋",
    Reading:     "レン・こ・う・こい・こい・しい",
    Translation: "romance; in love; yearn for; miss; darling",
  },
  Entry{
    Ji:          "紅",
    Reading:     "コウ・ク・べに・くれない・あか・い",
    Translation: "crimson; deep red",
  },
  Entry{
    Ji:          "郊",
    Reading:     "コウ",
    Translation: "outskirts; suburbs; rural area",
  },
  Entry{
    Ji:          "腰",
    Reading:     "ヨウ・こし",
    Translation: "loins; hips; waist; low wainscoting",
  },
  Entry{
    Ji:          "炭",
    Reading:     "タン・すみ",
    Translation: "charcoal; coal",
  },
  Entry{
    Ji:          "踊",
    Reading:     "ヨウ・おど・る",
    Translation: "jump; dance; leap; skip",
  },
  Entry{
    Ji:          "冊",
    Reading:     "サツ・サク・ふみ",
    Translation: "tome; counter for books; volume",
  },
  Entry{
    Ji:          "勇",
    Reading:     "ユウ・いさ・む",
    Translation: "courage; cheer up; be in high spirits; bravery; heroism",
  },
  Entry{
    Ji:          "械",
    Reading:     "カイ・かせ",
    Translation: "contraption; fetter; machine; instrument",
  },
  Entry{
    Ji:          "菜",
    Reading:     "サイ・な",
    Translation: "vegetable; side dish; greens",
  },
  Entry{
    Ji:          "珍",
    Reading:     "チン・めずら・しい・たから",
    Translation: "rare; curious; strange",
  },
  Entry{
    Ji:          "卵",
    Reading:     "ラン・たまご",
    Translation: "egg; ovum; spawn; roe",
  },
  Entry{
    Ji:          "湖",
    Reading:     "コ・みずうみ",
    Translation: "lake",
  },
  Entry{
    Ji:          "喫",
    Reading:     "キツ・の・む",
    Translation: "consume; eat; drink; smoke; receive (a blow",
  },
  Entry{
    Ji:          "干",
    Reading:     "カン・ほ・す・ほ・し-・-ぼ・し・ひ・る",
    Translation: "dry; parch",
  },
  Entry{
    Ji:          "虫",
    Reading:     "チュウ・キ・むし",
    Translation: "insect; bug; temper",
  },
  Entry{
    Ji:          "刷",
    Reading:     "サツ・す・る・-ず・り・-ずり・は・く",
    Translation: "printing; print; brush",
  },
  Entry{
    Ji:          "湯",
    Reading:     "トウ・ゆ",
    Translation: "hot water; bath; hot spring",
  },
  Entry{
    Ji:          "溶",
    Reading:     "ヨウ・と・ける・と・かす・と・く",
    Translation: "melt; dissolve; thaw",
  },
  Entry{
    Ji:          "鉱",
    Reading:     "コウ・あらがね",
    Translation: "mineral; ore",
  },
  Entry{
    Ji:          "涙",
    Reading:     "ルイ・レイ・なみだ",
    Translation: "tears; sympathy",
  },
  Entry{
    Ji:          "匹",
    Reading:     "ヒツ・ひき",
    Translation: "equal; head; counter for small animals; roll of cloth",
  },
  Entry{
    Ji:          "孫",
    Reading:     "ソン・まご",
    Translation: "grandchild; descendants",
  },
  Entry{
    Ji:          "鋭",
    Reading:     "エイ・するど・い",
    Translation: "pointed; sharpness; edge; weapon; sharp; violent",
  },
  Entry{
    Ji:          "枝",
    Reading:     "シ・えだ",
    Translation: "bough; branch; twig; limb; counter for branches",
  },
  Entry{
    Ji:          "塗",
    Reading:     "ト・ぬ・る・ぬ・り・まみ・れる",
    Translation: "paint; plaster; daub; smear; coating",
  },
  Entry{
    Ji:          "軒",
    Reading:     "ケン・のき",
    Translation: "flats; counter for houses; eaves",
  },
  Entry{
    Ji:          "毒",
    Reading:     "ドク",
    Translation: "poison; virus; venom; germ; harm; injury; spite",
  },
  Entry{
    Ji:          "叫",
    Reading:     "キョウ・さけ・ぶ",
    Translation: "shout; exclaim; yell",
  },
  Entry{
    Ji:          "拝",
    Reading:     "ハイ・おが・む・おろが・む",
    Translation: "worship; adore; pray to",
  },
  Entry{
    Ji:          "氷",
    Reading:     "ヒョウ・こおり・ひ・こお・る",
    Translation: "icicle; ice; hail; freeze; congeal",
  },
  Entry{
    Ji:          "乾",
    Reading:     "カン・ケン・かわ・く・かわ・かす・ほ・す・ひ・る・いぬい",
    Translation: "drought; dry; dessicate; drink up; heaven; emperor",
  },
  Entry{
    Ji:          "棒",
    Reading:     "ボウ",
    Translation: "rod; stick; cane; pole; club; line",
  },
  Entry{
    Ji:          "祈",
    Reading:     "キ・いの・る",
    Translation: "pray; wish",
  },
  Entry{
    Ji:          "拾",
    Reading:     "シュウ・ジュウ・ひろ・う",
    Translation: "pick up; gather; find; go on foot; ten",
  },
  Entry{
    Ji:          "粉",
    Reading:     "フン・デシメートル・こ・こな",
    Translation: "flour; powder; dust",
  },
  Entry{
    Ji:          "糸",
    Reading:     "シ・いと",
    Translation: "thread",
  },
  Entry{
    Ji:          "綿",
    Reading:     "メン・わた",
    Translation: "cotton",
  },
  Entry{
    Ji:          "汗",
    Reading:     "カン・あせ",
    Translation: "sweat; perspire",
  },
  Entry{
    Ji:          "銅",
    Reading:     "ドウ・あかがね",
    Translation: "copper",
  },
  Entry{
    Ji:          "湿",
    Reading:     "シツ・シュウ・しめ・る・しめ・す・うるお・う・うるお・す",
    Translation: "damp; wet; moist",
  },
  Entry{
    Ji:          "瓶",
    Reading:     "ビン・かめ",
    Translation: "flower pot; bottle; vial; jar; jug; vat; urn",
  },
  Entry{
    Ji:          "咲",
    Reading:     "ショウ・さ・く・-ざき",
    Translation: "blossom; bloom",
  },
  Entry{
    Ji:          "召",
    Reading:     "ショウ・め・す",
    Translation: "seduce; call; send for; wear; put on; ride in; buy; eat; drink; catch (cold",
  },
  Entry{
    Ji:          "缶",
    Reading:     "カン・かま",
    Translation: "tin can; container; jar radical (no. 121)",
  },
  Entry{
    Ji:          "隻",
    Reading:     "セキ",
    Translation: "vessels; counter for ships; fish; birds; arrows; one of a pair",
  },
  Entry{
    Ji:          "脂",
    Reading:     "シ・あぶら",
    Translation: "fat; grease; tallow; lard; rosin; gum; tar",
  },
  Entry{
    Ji:          "蒸",
    Reading:     "ジョウ・セイ・む・す・む・れる・む・らす",
    Translation: "steam; heat; sultry; foment; get musty",
  },
  Entry{
    Ji:          "肌",
    Reading:     "キ・はだ",
    Translation: "texture; skin; body; grain",
  },
  Entry{
    Ji:          "耕",
    Reading:     "コウ・たがや・す",
    Translation: "till; plow; cultivate",
  },
  Entry{
    Ji:          "鈍",
    Reading:     "ドン・にぶ・い・にぶ・る・にぶ-・なま・る・なまく・ら",
    Translation: "dull; slow; foolish; blunt",
  },
  Entry{
    Ji:          "泥",
    Reading:     "デイ・ナイ・デ・ニ・どろ",
    Translation: "mud; mire; adhere to; be attached to",
  },
  Entry{
    Ji:          "隅",
    Reading:     "グウ・すみ",
    Translation: "corner; nook",
  },
  Entry{
    Ji:          "灯",
    Reading:     "トウ・ひ・ほ-・ともしび・とも・す・あかり",
    Translation: "lamp; a light; light; counter for lights",
  },
  Entry{
    Ji:          "辛",
    Reading:     "シン・から・い・つら・い・-づら・い・かのと",
    Translation: "spicy; bitter; hot; acrid",
  },
  Entry{
    Ji:          "磨",
    Reading:     "マ・みが・く・す・る",
    Translation: "grind; polish; scour; improve; brush (teeth",
  },
  Entry{
    Ji:          "麦",
    Reading:     "バク・むぎ",
    Translation: "barley; wheat",
  },
  Entry{
    Ji:          "姓",
    Reading:     "セイ・ショウ",
    Translation: "surname",
  },
  Entry{
    Ji:          "筒",
    Reading:     "トウ・つつ",
    Translation: "cylinder; pipe; tube; gun barrel; sleeve",
  },
  Entry{
    Ji:          "鼻",
    Reading:     "ビ・はな",
    Translation: "nose; snout",
  },
  Entry{
    Ji:          "粒",
    Reading:     "リュウ・つぶ",
    Translation: "grains; drop; counter for tiny particles",
  },
  Entry{
    Ji:          "詞",
    Reading:     "シ",
    Translation: "part of speech; words; poetry",
  },
  Entry{
    Ji:          "胃",
    Reading:     "イ",
    Translation: "stomach; paunch; crop; craw",
  },
  Entry{
    Ji:          "畳",
    Reading:     "ジョウ・チョウ・たた・む・たたみ・かさ・なる",
    Translation: "tatami mat; counter for tatami mats; fold; shut up; do away with",
  },
  Entry{
    Ji:          "膚",
    Reading:     "フ・はだ",
    Translation: "skin; body; grain; texture; disposition",
  },
  Entry{
    Ji:          "机",
    Reading:     "キ・つくえ",
    Translation: "desk; table",
  },
  Entry{
    Ji:          "濯",
    Reading:     "タク・すす・ぐ・ゆす・ぐ",
    Translation: "laundry; wash; pour on; rinse",
  },
  Entry{
    Ji:          "塔",
    Reading:     "トウ",
    Translation: "pagoda; tower; steeple",
  },
  Entry{
    Ji:          "灰",
    Reading:     "カイ・はい",
    Translation: "ashes; puckery juice; cremate",
  },
  Entry{
    Ji:          "沸",
    Reading:     "フツ・わ・く・わ・かす",
    Translation: "seethe; boil; ferment; uproar; breed",
  },
  Entry{
    Ji:          "菓",
    Reading:     "カ",
    Translation: "candy; cakes; fruit",
  },
  Entry{
    Ji:          "帽",
    Reading:     "ボウ・モウ・ずきん・おお・う",
    Translation: "cap; headgear",
  },
  Entry{
    Ji:          "枯",
    Reading:     "コ・か・れる・か・らす",
    Translation: "wither; die; dry up; be seasoned",
  },
  Entry{
    Ji:          "涼",
    Reading:     "リョウ・すず・しい・すず・む・すず・やか・うす・い・ひや・す・まことに",
    Translation: "refreshing; nice and cool",
  },
  Entry{
    Ji:          "舟",
    Reading:     "シュウ・ふね・ふな-・-ぶね",
    Translation: "boat; ship",
  },
  Entry{
    Ji:          "貝",
    Reading:     "バイ・かい",
    Translation: "shellfish",
  },
  Entry{
    Ji:          "符",
    Reading:     "フ",
    Translation: "token; sign; mark; tally; charm",
  },
  Entry{
    Ji:          "憎",
    Reading:     "ゾウ・にく・む・にく・い・にく・らしい・にく・しみ",
    Translation: "hate; detest",
  },
  Entry{
    Ji:          "皿",
    Reading:     "ベイ・さら",
    Translation: "dish; a helping; plate",
  },
  Entry{
    Ji:          "肯",
    Reading:     "コウ・がえんじ・る",
    Translation: "agreement; consent; comply with",
  },
  Entry{
    Ji:          "燥",
    Reading:     "ソウ・はしゃ・ぐ",
    Translation: "parch; dry up",
  },
  Entry{
    Ji:          "畜",
    Reading:     "チク",
    Translation: "livestock; domestic fowl and animals",
  },
  Entry{
    Ji:          "挟",
    Reading:     "キョウ・ショウ・はさ・む・はさ・まる・わきばさ・む・さしはさ・む",
    Translation: "pinch; between",
  },
  Entry{
    Ji:          "曇",
    Reading:     "ドン・くも・る",
    Translation: "cloudy weather; cloud up",
  },
  Entry{
    Ji:          "滴",
    Reading:     "テキ・しずく・したた・る",
    Translation: "drip; drop",
  },
  Entry{
    Ji:          "伺",
    Reading:     "シ・うかが・う",
    Translation: "pay respects; visit; ask; inquire; question; implore",
  },
}
