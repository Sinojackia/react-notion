package scheme

type number int

type boolean bool

type any interface{}

//======================================================================================
type BaseValueTypeKind string

const (
	Page           BaseValueTypeKind = "page"
	Text           BaseValueTypeKind = "text"
	BulletedList   BaseValueTypeKind = "bulleted_list"
	NumberedList   BaseValueTypeKind = "numbered_list"
	Header         BaseValueTypeKind = "header"
	SubHeader      BaseValueTypeKind = "sub_header"
	SubSubHeader   BaseValueTypeKind = "sub_sub_header"
	Quote          BaseValueTypeKind = "quote"
	ColumnList     BaseValueTypeKind = "column_list"
	Divider        BaseValueTypeKind = "divider"
	Image          BaseValueTypeKind = "image"
	Embed          BaseValueTypeKind = "embed"
	Figma          BaseValueTypeKind = "figma"
	Video          BaseValueTypeKind = "video"
	CollectionView BaseValueTypeKind = "collection_view"
	Column         BaseValueTypeKind = "column"
	Code           BaseValueTypeKind = "code"
	Callout        BaseValueTypeKind = "callout"
	Toggle         BaseValueTypeKind = "toggle"
	Bookmark       BaseValueTypeKind = "bookmark"
	Tweet          BaseValueTypeKind = "tweet"
	ToDo           BaseValueTypeKind = "to_do"
	Gallery        BaseValueTypeKind = "gallery"
	Table          BaseValueTypeKind = "table"
)

type BaseValueType struct {
	ID                string   `json:"id"`
	Version           number   `json:"version"`
	CreatedTime       number   `json:"created_time"`
	LastEditedTime    number   `json:"last_edited_time"`
	ParentId          string   `json:"parent_id"`
	ParentTable       string   `json:"parent_table"`
	Alive             boolean  `json:"alive"`
	CreatedByTable    string   `json:"created_by_table"`
	CreatedById       string   `json:"created_by_id"`
	LastEditedByTable string   `json:"last_edited_by_table"`
	LastEditedById    string   `json:"last_edited_by_id"`
	SpaceId           string   `json:"space_id"`
	Properties        any      `json:"properties"`
	Content           []string `json:"content"`
}

type PageValueType struct {
	BaseValueType `json:",inline"`
	Type          BaseValueTypeKind `json:"type"` // type=page
	Properties    struct {
		Title []DecorationType `json:"title"`
	} `json:"properties"`
	Format struct {
		PageFullWidth     boolean `json:"page_full_width"`
		PageSmallText     boolean `json:"page_small_text"`
		PageCoverPosition number  `json:"page_cover_position"`
		BlockLocked       boolean `json:"block_locked"`
		BlockLockedBy     string  `json:"block_locked_by"`
		PageCover         string  `json:"page_cover"`
		PageIcon          string  `json:"page_icon"`
	} `json:"format"`
	Permissions []struct {
		Role string `json:"role"`
		Type string `json:"type"`
	} `json:"permissions"`
	FileIds []string `json:"file_ids"`
}

type TextValueType struct {
	BaseValueType `json:",inline"`
	Type          BaseValueTypeKind `json:"type"` // type=text
}

type BulletedListValueType struct {
	BaseValueType `json:",inline"`
	Type          BaseValueTypeKind `json:"type"` // type=bulleted_list
}

type NumberedListValueType struct {
	BaseValueType `json:",inline"`
	Type          BaseValueTypeKind `json:"type"` // type=numbered_list
}

type HeaderValueType struct {
	BaseValueType `json:",inline"`
	Type          BaseValueTypeKind `json:"type"` // type=header
}

type SubHeaderValueType struct {
	BaseValueType `json:",inline"`
	Type          BaseValueTypeKind `json:"type"` // type=sub_header
}

type SubSubHeaderValueType struct {
	BaseValueType `json:",inline"`
	Type          BaseValueTypeKind `json:"type"` // type=sub_sub_header
}

type QuoteValueType struct {
	BaseValueType `json:",inline"`
	Type          BaseValueTypeKind `json:"type"` // type=quote
}

type ColumnListValueType struct {
	BaseValueType `json:",inline"`
	Type          BaseValueTypeKind `json:"type"` // type=column_list
}

type DividerValueType struct {
	BaseValueType `json:",inline"`
	Type          BaseValueTypeKind `json:"type"` // type=divider
}

type ContentValueType struct {
	BaseValueType `json:",inline"`
	Properties    struct {
		Source  [][]string       `json:"source"`
		Caption []DecorationType `json:"caption"`
	} `json:"properties"`
	Format struct {
		BlockWidth         number  `json:"block_width"`
		BlockHeight        number  `json:"block_height"`
		DisplaySource      string  `json:"display_source"`
		BlockFullWidth     boolean `json:"block_full_width"`
		BlockPageWidth     boolean `json:"block_page_width"`
		BlockAspectRatio   number  `json:"block_aspect_ratio"`
		BlockPreserveScale boolean `json:"block_preserve_scale"`
	} `json:"format"`
	FileIds []string `json:"file_ids"`
}

type ImageValueType struct {
	ContentValueType `json:",inline"`
	Type             BaseValueTypeKind `json:"type"` // type=image
}

type EmbedValueType struct {
	ContentValueType `json:",inline"`
	Type             BaseValueTypeKind `json:"type"` // type=embed
}

type FigmaValueType struct {
	ContentValueType `json:",inline"`
	Type             BaseValueTypeKind `json:"type"` // type=figma
}

type VideoValueType struct {
	ContentValueType `json:",inline"`
	Type             BaseValueTypeKind `json:"type"` // type=video
}

type CollectionValueType struct {
	ContentValueType `json:",inline"`
	Type             BaseValueTypeKind `json:"type"` // type=collection_view
}

type ColumnValueType struct {
	BaseValueType `json:",inline"`
	Type          BaseValueTypeKind `json:"type"` // type=column
	Format        struct {
		ColumnRatio number `json:"column_ratio"`
	} `json:"format"`
}

type CodeValueType struct {
	BaseValueType `json:",inline"`
	Type          BaseValueTypeKind `json:"type"` // type=code
	Properties    struct {
		Title    []DecorationType `json:"title"`
		Language []DecorationType `json:"language"`
	} `json:"properties"`
}

type CalloutValueType struct {
	BaseValueType `json:",inline"`
	Type          BaseValueTypeKind `json:"type"` // type=callout
	Properties    struct {
		Title []DecorationType `json:"title"`
	} `json:"properties"`
	Format struct {
		PageIcon   string    `json:"page_icon"`
		BlockColor ColorKind `json:"block_color"`
	} `json:"format"`
}

type ToggleValueType struct {
	BaseValueType `json:",inline"`
	Type          BaseValueTypeKind `json:"type"` // type=toggle
	Properties    struct {
		Title []DecorationType `json:"title"`
	} `json:"properties"`
}

type BookmarkValueType struct {
	BaseValueType `json:",inline"`
	Type          BaseValueTypeKind `json:"type"` // type=bookmark
	Properties    struct {
		Title       []DecorationType `json:"title"`
		Link        []DecorationType `json:"link"`
		Description []DecorationType `json:"description"`
	} `json:"properties"`
	Format struct {
		BlockColor    string `json:"block_color"`
		BookmarkIcon  string `json:"bookmark_icon"`
		BookmarkCover string `json:"bookmark_cover"`
	} `json:"format"`
}

type TweetType struct {
	BaseValueType `json:",inline"`
	Type          BaseValueTypeKind `json:"type"` // type=tweet
	Properties    struct {
		Source [][]string `json:"source"`
	} `json:"properties"`
}

type CheckedType string

const (
	Yes CheckedType = "Yes"
	No              = "No"
)

type TodoValueType struct {
	BaseValueType `json:",inline"`
	Type          BaseValueTypeKind `json:"type"` // type=to_do
	Properties    struct {
		title   []DecorationType `json:"title"`
		Checked []CheckedType    `json:"checked"` //"Yes" "No"
	} `json:"properties"`
}

type BlockValueType struct {
	PageValueType         `json:",inline"`
	TextValueType         `json:",inline"`
	BulletedListValueType `json:",inline"`
	NumberedListValueType `json:",inline"`
	HeaderValueType       `json:",inline"`
	SubHeaderValueType    `json:",inline"`
	SubSubHeaderValueType `json:",inline"`
	QuoteValueType        `json:",inline"`
	ColumnListValueType   `json:",inline"`
	DividerValueType      `json:",inline"`
	ContentValueType      `json:",inline"`
	CollectionValueType   `json:",inline"`
	ColumnValueType       `json:",inline"`
	CodeValueType         `json:",inline"`
	CalloutValueType      `json:",inline"`
	ToggleValueType       `json:",inline"`
	BookmarkValueType     `json:",inline"`
	TweetType             `json:",inline"`
	TodoValueType         `json:",inline"`
	ID                    string   `json:"id"`
	Version               number   `json:"version"`
	CreatedTime           number   `json:"created_time"`
	LastEditedTime        number   `json:"last_edited_time"`
	ParentId              string   `json:"parent_id"`
	ParentTable           string   `json:"parent_table"`
	Alive                 boolean  `json:"alive"`
	CreatedByTable        string   `json:"created_by_table"`
	CreatedById           string   `json:"created_by_id"`
	LastEditedByTable     string   `json:"last_edited_by_table"`
	LastEditedById        string   `json:"last_edited_by_id"`
	SpaceId               string   `json:"space_id"`
	properties            any      `json:"properties"`
	Content               []string `json:"content"`
}

type TableGalleryType struct {
	BaseValueType
	Type   BaseValueTypeKind `json:"type"` //type=gallery
	Format struct {
		GalleryCover struct {
			Type string
		}
		GalleryCoverAspect string
		GalleryProperties  []struct {
			Visible  boolean
			Property string
		}
	}
}

type TableCollectionType struct {
	BaseValueType
	Type   BaseValueTypeKind `json:"type"` //type=table
	Format struct {
		TableWrap       boolean
		TableProperties []struct {
			Visible  boolean
			Property string
			Width    number
		}
	}
}

type DecorationType []string

var (
	BoldFormatType   DecorationType = []string{"b"}
	ItalicFormatType                = []string{"i"}
	StrikeFormatType                = []string{"s"}
	CodeFormatType                  = []string{"c"}
	LinkFormatType                  = []string{"a"}
	ColorFormatType                 = []ColorKind{"h"}
	DateFormatType                  = []string{"d"} //todo 格式化类型
	UserFormatType                  = []string{"u"}
	PageFormatType                  = []string{"p"}
)

type BaseTextValueType struct {
	properties struct {
		Title []DecorationType
	}
	format struct {
		BlockColor ColorKind `json:"block_color"`
	}
}

//======================================================================================

type ColorKind string

// 类型枚举
var ColorTypeEnums = []ColorKind{"gray", "brown", "orange", "yellow", "teal", "blue", "purple", "pink", "red",
	"gray_background", "brown_background", "orange_background", "yellow_background", "teal_background",
	"blue_background", "purple_background", "pink_background", "red_background"}

type CollectionViewType struct {
	TableGalleryType
	TableCollectionType
}

type BlockType struct {
	Role       string         `json:"role"`
	Value      BlockValueType `json:"value"`
	Collection struct {
		Title  []DecorationType            `json:"title"`
		Types  []CollectionViewType        `json:"types"`
		Data   map[string][]DecorationType `json:"data"`
		schema map[string]struct {
			Name string `json:"name"`
			Type string `json:"type"`
		} `json:"schema"`
	} `json:"collection"`
}

type BlockMapType map[string]BlockType

type NotionUserType struct {
	Role  string `json:"role"`
	Value struct {
		ID                        string
		Version                   number
		Email                     string
		GivenName                 string
		FamilyName                string
		ProfilePhoto              string
		OnboardingCompleted       boolean
		MobileOnboardingCompleted boolean
	}
}

type RecordMapType struct {
	Block      BlockMapType
	NotionUser map[string]NotionUserType
}

type LoadPageChunkData struct {
	RecordMap RecordMapType
	Cursor    struct {
		Stack []any
	}
}

type MapPageUrl func(pageId string) string
type MapImageUrl func(image string, block BlockType) string
