package model

//======================================================================================
type BaseValueKind string

const (
	Page         BaseValueKind = "page"
	Text         BaseValueKind = "text"
	BulletedList BaseValueKind = "bulleted_list"
	NumberedList BaseValueKind = "numbered_list"
	Header       BaseValueKind = "header"
	SubHeader    BaseValueKind = "sub_header"
	SubSubHeader BaseValueKind = "sub_sub_header"
	Quote        BaseValueKind = "quote"
	ColumnList   BaseValueKind = "column_list"
	Divider      BaseValueKind = "divider"

	Column   BaseValueKind = "column"
	Code     BaseValueKind = "code"
	Callout  BaseValueKind = "callout"
	Toggle   BaseValueKind = "toggle"
	Bookmark BaseValueKind = "bookmark"
	Tweet    BaseValueKind = "tweet"
	ToDo     BaseValueKind = "to_do"
	Gallery  BaseValueKind = "gallery"
	Table    BaseValueKind = "table"
)

type ContentValueType BaseValueKind

const (
	Image          BaseValueKind = "image"
	Embed          BaseValueKind = "embed"
	Figma          BaseValueKind = "figma"
	Video          BaseValueKind = "video"
	CollectionView BaseValueKind = "collection_view"
)

type Block struct {
	tableName  struct{} `pg:"block"`
	Role       string   `json:"role"`
	Value      Value    `json:"value"`
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

type Value struct {
	tableName         struct{}                 `pg:"value"`
	ID                string                   `json:"id" pg:"id,type:uuid,pk"`
	Type              BaseValueKind            `json:"type" pg:"type"`
	Version           int                      `json:"version" pg:"version"`
	CreatedTime       int64                    `json:"created_time" pg:"created_time,type:int8,default:floor(extract(epoch from now())),notnull"`
	LastEditedTime    int64                    `json:"last_edited_time" pg:"last_edited_time,type:int8,default:floor(extract(epoch from now())),notnull"`
	ParentId          string                   `json:"parent_id" pg:"parent_id,type:uuid,notnull"`
	ParentTable       string                   `json:"parent_table" pg:"parent_table"`
	Alive             bool                     `json:"alive" pg:"alive"`
	CreatedByTable    string                   `json:"created_by_table" pg:"created_by_table"`
	CreatedById       string                   `json:"created_by_id" pg:"created_by_id,type:uuid,notnull"`
	LastEditedByTable string                   `json:"last_edited_by_table" pg:"last_edited_by_table"`
	LastEditedById    string                   `json:"last_edited_by_id" pg:"last_edited_by_id,type:uuid,notnull"`
	SpaceId           string                   `json:"space_id" pg:"space_id,type:uuid"`
	Properties        map[string]interface{}   `json:"properties,omitempty" pg:"properties,type:jsonb"`
	Content           []string                 `json:"content,omitempty" pg:"content,array"`
	Format            map[string]interface{}   `json:"format,omitempty" pg:"format,type:jsonb"`
	Permissions       []map[string]interface{} `json:"permissions,omitempty" pg:"permissions,array"`
	FileIds           []string                 `json:"file_ids,omitempty" pg:"file_ids,array"`
}

type PageValue struct {
	Value      `json:",inline"`
	Properties struct {
		Title []DecorationType `json:"title,omitempty"`
	} `json:"properties"`
	Format struct {
		PageFullWidth     bool   `json:"page_full_width"`
		PageSmallText     bool   `json:"page_small_text"`
		PageCoverPosition int    `json:"page_cover_position"`
		BlockLocked       bool   `json:"block_locked"`
		BlockLockedBy     string `json:"block_locked_by"`
		PageCover         string `json:"page_cover"`
		PageIcon          string `json:"page_icon"`
	} `json:"format"`
	Permissions []struct {
		Role string `json:"role"`
		Type string `json:"type"`
	} `json:"permissions"`
	FileIds []string `json:"file_ids"`
}

type ContentValue struct {
	Value      `json:",inline"`
	Type       ContentValueType `json:"type"`
	Properties struct {
		Source  [][]string       `json:"source"`
		Caption []DecorationType `json:"caption"`
	} `json:"properties"`
	Format struct {
		BlockWidth         int    `json:"block_width"`
		BlockHeight        int    `json:"block_height"`
		DisplaySource      string `json:"display_source"`
		BlockFullWidth     bool   `json:"block_full_width"`
		BlockPageWidth     bool   `json:"block_page_width"`
		BlockAspectRatio   int    `json:"block_aspect_ratio"`
		BlockPreserveScale bool   `json:"block_preserve_scale"`
	} `json:"format"`
	FileIds []string `json:"file_ids"`
}

type ColumnValue struct {
	Value  `json:",inline"`
	Type   BaseValueKind `json:"type"` // type=column
	Format struct {
		ColumnRatio int `json:"column_ratio"`
	} `json:"format"`
}

type CodeValue struct {
	Value      `json:",inline"`
	Type       BaseValueKind `json:"type"` // type=code
	Properties struct {
		Title    []DecorationType `json:"title"`
		Language []DecorationType `json:"language"`
	} `json:"properties"`
}

type CalloutValue struct {
	Value      `json:",inline"`
	Type       BaseValueKind `json:"type"` // type=callout
	Properties struct {
		Title []DecorationType `json:"title"`
	} `json:"properties"`
	Format struct {
		PageIcon   string    `json:"page_icon"`
		BlockColor ColorType `json:"block_color"`
	} `json:"format"`
}

type ToggleValue struct {
	Value      `json:",inline"`
	Type       BaseValueKind `json:"type"` // type=toggle
	Properties struct {
		Title []DecorationType `json:"title"`
	} `json:"properties"`
}

type BookmarkValue struct {
	Value      `json:",inline"`
	Type       BaseValueKind `json:"type"` // type=bookmark
	Properties struct {
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
	Value      `json:",inline"`
	Type       BaseValueKind `json:"type"` // type=tweet
	Properties struct {
		Source [][]string `json:"source"`
	} `json:"properties"`
}

type CheckedType string

const (
	Yes CheckedType = "Yes"
	No              = "No"
)

type TodoValue struct {
	Value      `json:",inline"`
	Type       BaseValueKind `json:"type"` // type=to_do
	Properties struct {
		title   []DecorationType `json:"title"`
		Checked []CheckedType    `json:"checked"` //"Yes" "No"
	} `json:"properties"`
}

type TableGalleryValue struct {
	Value  `json:",inline"`
	Type   BaseValueKind `json:"type"` //type=gallery
	Format struct {
		GalleryCover struct {
			Type string
		}
		GalleryCoverAspect string
		GalleryProperties  []struct {
			Visible  bool
			Property string
		}
	}
}

type TableCollectionValue struct {
	Value  `json:",inline"`
	Type   BaseValueKind `json:"type"` //type=table
	Format struct {
		TableWrap       bool
		TableProperties []struct {
			Visible  bool
			Property string
			Width    int
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
	ColorFormatType                 = []ColorType{"h"}
	DateFormatType                  = []string{"d"} //todo 格式化类型
	UserFormatType                  = []string{"u"}
	PageFormatType                  = []string{"p"}
)

type TextValue struct {
	properties struct {
		Title []DecorationType
	}
	format struct {
		BlockColor ColorType `json:"block_color"`
	}
}

//======================================================================================

type ColorType string

// 类型枚举
var ColorTypeEnums = []ColorType{"gray", "brown", "orange", "yellow", "teal", "blue", "purple", "pink", "red",
	"gray_background", "brown_background", "orange_background", "yellow_background", "teal_background",
	"blue_background", "purple_background", "pink_background", "red_background"}

type CollectionViewType struct {
	TableGalleryValue
	TableCollectionValue
}

type NotionUser struct {
	Role  string `json:"role"`
	Value struct {
		ID                        string
		Version                   int
		Email                     string
		GivenName                 string
		FamilyName                string
		ProfilePhoto              string
		OnboardingCompleted       bool
		MobileOnboardingCompleted bool
	} `json:"value"`
}

type RecordMap struct {
	Block      map[string]Block      `json:"block"`
	NotionUser map[string]NotionUser `json:"notion_user"`
}

type LoadPageChunkData struct {
	RecordMap RecordMap `json:"record_map"`
	Cursor    struct {
		Stack []interface{}
	} `json:"cursor"`
}

type MapPageUrl func(pageId string) string

type MapImageUrl func(image string, block Block) string
