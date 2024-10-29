# TypeStyle

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FontFamily** | Pointer to **string** | Font family of text (standard name). | [optional] 
**FontPostScriptName** | Pointer to **NullableString** | PostScript font name. | [optional] 
**ParagraphSpacing** | Pointer to **float32** | Space between paragraphs in px, 0 if not present. | [optional] [default to 0]
**ParagraphIndent** | Pointer to **float32** | Paragraph indentation in px, 0 if not present. | [optional] [default to 0]
**ListSpacing** | Pointer to **float32** | Space between list items in px, 0 if not present. | [optional] [default to 0]
**Italic** | Pointer to **bool** | Whether or not text is italicized. | [optional] [default to false]
**FontWeight** | Pointer to **float32** | Numeric font weight. | [optional] 
**FontSize** | Pointer to **float32** | Font size in px. | [optional] 
**TextCase** | Pointer to **string** | Text casing applied to the node, default is the original casing. | [optional] 
**TextDecoration** | Pointer to **string** | Text decoration applied to the node, default is none. | [optional] [default to "NONE"]
**TextAutoResize** | Pointer to **string** | Dimensions along which text will auto resize, default is that the text does not auto-resize. TRUNCATE means that the text will be shortened and trailing text will be replaced with \&quot;…\&quot; if the text contents is larger than the bounds. &#x60;TRUNCATE&#x60; as a return value is deprecated and will be removed in a future version. Read from &#x60;textTruncation&#x60; instead. | [optional] [default to "NONE"]
**TextTruncation** | Pointer to **string** | Whether this text node will truncate with an ellipsis when the text contents is larger than the text node. | [optional] [default to "DISABLED"]
**MaxLines** | Pointer to **float32** | When &#x60;textTruncation: \&quot;ENDING\&quot;&#x60; is set, &#x60;maxLines&#x60; determines how many lines a text node can grow to before it truncates. | [optional] 
**TextAlignHorizontal** | Pointer to **string** | Horizontal text alignment as string enum. | [optional] 
**TextAlignVertical** | Pointer to **string** | Vertical text alignment as string enum. | [optional] 
**LetterSpacing** | Pointer to **float32** | Space between characters in px. | [optional] 
**Fills** | Pointer to [**[]Paint**](Paint.md) | An array of fill paints applied to the characters. | [optional] 
**Hyperlink** | Pointer to [**Hyperlink**](Hyperlink.md) | Link to a URL or frame. | [optional] 
**OpentypeFlags** | Pointer to **map[string]float32** | A map of OpenType feature flags to 1 or 0, 1 if it is enabled and 0 if it is disabled. Note that some flags aren&#39;t reflected here. For example, SMCP (small caps) is still represented by the &#x60;textCase&#x60; field. | [optional] 
**LineHeightPx** | Pointer to **float32** | Line height in px. | [optional] 
**LineHeightPercent** | Pointer to **float32** | Line height as a percentage of normal line height. This is deprecated; in a future version of the API only lineHeightPx and lineHeightPercentFontSize will be returned. | [optional] [default to 100]
**LineHeightPercentFontSize** | Pointer to **float32** | Line height as a percentage of the font size. Only returned when &#x60;lineHeightPercent&#x60; (deprecated) is not 100. | [optional] 
**LineHeightUnit** | Pointer to **string** | The unit of the line height value specified by the user. | [optional] 
**BoundVariables** | Pointer to [**TypeStyleBoundVariables**](TypeStyleBoundVariables.md) |  | [optional] 
**IsOverrideOverTextStyle** | Pointer to **bool** |  Whether or not this style has overrides over a text style. The possible fields to override are semanticWeight, semanticItalic, hyperlink, and textDecoration. If this is true, then those fields are overrides if present. | [optional] 
**SemanticWeight** | Pointer to **string** | Indicates how the font weight was overridden when there is a text style override. | [optional] 
**SemanticItalic** | Pointer to **string** | Indicates how the font style was overridden when there is a text style override. | [optional] 

## Methods

### NewTypeStyle

`func NewTypeStyle() *TypeStyle`

NewTypeStyle instantiates a new TypeStyle object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTypeStyleWithDefaults

`func NewTypeStyleWithDefaults() *TypeStyle`

NewTypeStyleWithDefaults instantiates a new TypeStyle object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFontFamily

`func (o *TypeStyle) GetFontFamily() string`

GetFontFamily returns the FontFamily field if non-nil, zero value otherwise.

### GetFontFamilyOk

`func (o *TypeStyle) GetFontFamilyOk() (*string, bool)`

GetFontFamilyOk returns a tuple with the FontFamily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFontFamily

`func (o *TypeStyle) SetFontFamily(v string)`

SetFontFamily sets FontFamily field to given value.

### HasFontFamily

`func (o *TypeStyle) HasFontFamily() bool`

HasFontFamily returns a boolean if a field has been set.

### GetFontPostScriptName

`func (o *TypeStyle) GetFontPostScriptName() string`

GetFontPostScriptName returns the FontPostScriptName field if non-nil, zero value otherwise.

### GetFontPostScriptNameOk

`func (o *TypeStyle) GetFontPostScriptNameOk() (*string, bool)`

GetFontPostScriptNameOk returns a tuple with the FontPostScriptName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFontPostScriptName

`func (o *TypeStyle) SetFontPostScriptName(v string)`

SetFontPostScriptName sets FontPostScriptName field to given value.

### HasFontPostScriptName

`func (o *TypeStyle) HasFontPostScriptName() bool`

HasFontPostScriptName returns a boolean if a field has been set.

### SetFontPostScriptNameNil

`func (o *TypeStyle) SetFontPostScriptNameNil(b bool)`

 SetFontPostScriptNameNil sets the value for FontPostScriptName to be an explicit nil

### UnsetFontPostScriptName
`func (o *TypeStyle) UnsetFontPostScriptName()`

UnsetFontPostScriptName ensures that no value is present for FontPostScriptName, not even an explicit nil
### GetParagraphSpacing

`func (o *TypeStyle) GetParagraphSpacing() float32`

GetParagraphSpacing returns the ParagraphSpacing field if non-nil, zero value otherwise.

### GetParagraphSpacingOk

`func (o *TypeStyle) GetParagraphSpacingOk() (*float32, bool)`

GetParagraphSpacingOk returns a tuple with the ParagraphSpacing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParagraphSpacing

`func (o *TypeStyle) SetParagraphSpacing(v float32)`

SetParagraphSpacing sets ParagraphSpacing field to given value.

### HasParagraphSpacing

`func (o *TypeStyle) HasParagraphSpacing() bool`

HasParagraphSpacing returns a boolean if a field has been set.

### GetParagraphIndent

`func (o *TypeStyle) GetParagraphIndent() float32`

GetParagraphIndent returns the ParagraphIndent field if non-nil, zero value otherwise.

### GetParagraphIndentOk

`func (o *TypeStyle) GetParagraphIndentOk() (*float32, bool)`

GetParagraphIndentOk returns a tuple with the ParagraphIndent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParagraphIndent

`func (o *TypeStyle) SetParagraphIndent(v float32)`

SetParagraphIndent sets ParagraphIndent field to given value.

### HasParagraphIndent

`func (o *TypeStyle) HasParagraphIndent() bool`

HasParagraphIndent returns a boolean if a field has been set.

### GetListSpacing

`func (o *TypeStyle) GetListSpacing() float32`

GetListSpacing returns the ListSpacing field if non-nil, zero value otherwise.

### GetListSpacingOk

`func (o *TypeStyle) GetListSpacingOk() (*float32, bool)`

GetListSpacingOk returns a tuple with the ListSpacing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListSpacing

`func (o *TypeStyle) SetListSpacing(v float32)`

SetListSpacing sets ListSpacing field to given value.

### HasListSpacing

`func (o *TypeStyle) HasListSpacing() bool`

HasListSpacing returns a boolean if a field has been set.

### GetItalic

`func (o *TypeStyle) GetItalic() bool`

GetItalic returns the Italic field if non-nil, zero value otherwise.

### GetItalicOk

`func (o *TypeStyle) GetItalicOk() (*bool, bool)`

GetItalicOk returns a tuple with the Italic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItalic

`func (o *TypeStyle) SetItalic(v bool)`

SetItalic sets Italic field to given value.

### HasItalic

`func (o *TypeStyle) HasItalic() bool`

HasItalic returns a boolean if a field has been set.

### GetFontWeight

`func (o *TypeStyle) GetFontWeight() float32`

GetFontWeight returns the FontWeight field if non-nil, zero value otherwise.

### GetFontWeightOk

`func (o *TypeStyle) GetFontWeightOk() (*float32, bool)`

GetFontWeightOk returns a tuple with the FontWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFontWeight

`func (o *TypeStyle) SetFontWeight(v float32)`

SetFontWeight sets FontWeight field to given value.

### HasFontWeight

`func (o *TypeStyle) HasFontWeight() bool`

HasFontWeight returns a boolean if a field has been set.

### GetFontSize

`func (o *TypeStyle) GetFontSize() float32`

GetFontSize returns the FontSize field if non-nil, zero value otherwise.

### GetFontSizeOk

`func (o *TypeStyle) GetFontSizeOk() (*float32, bool)`

GetFontSizeOk returns a tuple with the FontSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFontSize

`func (o *TypeStyle) SetFontSize(v float32)`

SetFontSize sets FontSize field to given value.

### HasFontSize

`func (o *TypeStyle) HasFontSize() bool`

HasFontSize returns a boolean if a field has been set.

### GetTextCase

`func (o *TypeStyle) GetTextCase() string`

GetTextCase returns the TextCase field if non-nil, zero value otherwise.

### GetTextCaseOk

`func (o *TypeStyle) GetTextCaseOk() (*string, bool)`

GetTextCaseOk returns a tuple with the TextCase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTextCase

`func (o *TypeStyle) SetTextCase(v string)`

SetTextCase sets TextCase field to given value.

### HasTextCase

`func (o *TypeStyle) HasTextCase() bool`

HasTextCase returns a boolean if a field has been set.

### GetTextDecoration

`func (o *TypeStyle) GetTextDecoration() string`

GetTextDecoration returns the TextDecoration field if non-nil, zero value otherwise.

### GetTextDecorationOk

`func (o *TypeStyle) GetTextDecorationOk() (*string, bool)`

GetTextDecorationOk returns a tuple with the TextDecoration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTextDecoration

`func (o *TypeStyle) SetTextDecoration(v string)`

SetTextDecoration sets TextDecoration field to given value.

### HasTextDecoration

`func (o *TypeStyle) HasTextDecoration() bool`

HasTextDecoration returns a boolean if a field has been set.

### GetTextAutoResize

`func (o *TypeStyle) GetTextAutoResize() string`

GetTextAutoResize returns the TextAutoResize field if non-nil, zero value otherwise.

### GetTextAutoResizeOk

`func (o *TypeStyle) GetTextAutoResizeOk() (*string, bool)`

GetTextAutoResizeOk returns a tuple with the TextAutoResize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTextAutoResize

`func (o *TypeStyle) SetTextAutoResize(v string)`

SetTextAutoResize sets TextAutoResize field to given value.

### HasTextAutoResize

`func (o *TypeStyle) HasTextAutoResize() bool`

HasTextAutoResize returns a boolean if a field has been set.

### GetTextTruncation

`func (o *TypeStyle) GetTextTruncation() string`

GetTextTruncation returns the TextTruncation field if non-nil, zero value otherwise.

### GetTextTruncationOk

`func (o *TypeStyle) GetTextTruncationOk() (*string, bool)`

GetTextTruncationOk returns a tuple with the TextTruncation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTextTruncation

`func (o *TypeStyle) SetTextTruncation(v string)`

SetTextTruncation sets TextTruncation field to given value.

### HasTextTruncation

`func (o *TypeStyle) HasTextTruncation() bool`

HasTextTruncation returns a boolean if a field has been set.

### GetMaxLines

`func (o *TypeStyle) GetMaxLines() float32`

GetMaxLines returns the MaxLines field if non-nil, zero value otherwise.

### GetMaxLinesOk

`func (o *TypeStyle) GetMaxLinesOk() (*float32, bool)`

GetMaxLinesOk returns a tuple with the MaxLines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxLines

`func (o *TypeStyle) SetMaxLines(v float32)`

SetMaxLines sets MaxLines field to given value.

### HasMaxLines

`func (o *TypeStyle) HasMaxLines() bool`

HasMaxLines returns a boolean if a field has been set.

### GetTextAlignHorizontal

`func (o *TypeStyle) GetTextAlignHorizontal() string`

GetTextAlignHorizontal returns the TextAlignHorizontal field if non-nil, zero value otherwise.

### GetTextAlignHorizontalOk

`func (o *TypeStyle) GetTextAlignHorizontalOk() (*string, bool)`

GetTextAlignHorizontalOk returns a tuple with the TextAlignHorizontal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTextAlignHorizontal

`func (o *TypeStyle) SetTextAlignHorizontal(v string)`

SetTextAlignHorizontal sets TextAlignHorizontal field to given value.

### HasTextAlignHorizontal

`func (o *TypeStyle) HasTextAlignHorizontal() bool`

HasTextAlignHorizontal returns a boolean if a field has been set.

### GetTextAlignVertical

`func (o *TypeStyle) GetTextAlignVertical() string`

GetTextAlignVertical returns the TextAlignVertical field if non-nil, zero value otherwise.

### GetTextAlignVerticalOk

`func (o *TypeStyle) GetTextAlignVerticalOk() (*string, bool)`

GetTextAlignVerticalOk returns a tuple with the TextAlignVertical field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTextAlignVertical

`func (o *TypeStyle) SetTextAlignVertical(v string)`

SetTextAlignVertical sets TextAlignVertical field to given value.

### HasTextAlignVertical

`func (o *TypeStyle) HasTextAlignVertical() bool`

HasTextAlignVertical returns a boolean if a field has been set.

### GetLetterSpacing

`func (o *TypeStyle) GetLetterSpacing() float32`

GetLetterSpacing returns the LetterSpacing field if non-nil, zero value otherwise.

### GetLetterSpacingOk

`func (o *TypeStyle) GetLetterSpacingOk() (*float32, bool)`

GetLetterSpacingOk returns a tuple with the LetterSpacing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLetterSpacing

`func (o *TypeStyle) SetLetterSpacing(v float32)`

SetLetterSpacing sets LetterSpacing field to given value.

### HasLetterSpacing

`func (o *TypeStyle) HasLetterSpacing() bool`

HasLetterSpacing returns a boolean if a field has been set.

### GetFills

`func (o *TypeStyle) GetFills() []Paint`

GetFills returns the Fills field if non-nil, zero value otherwise.

### GetFillsOk

`func (o *TypeStyle) GetFillsOk() (*[]Paint, bool)`

GetFillsOk returns a tuple with the Fills field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFills

`func (o *TypeStyle) SetFills(v []Paint)`

SetFills sets Fills field to given value.

### HasFills

`func (o *TypeStyle) HasFills() bool`

HasFills returns a boolean if a field has been set.

### GetHyperlink

`func (o *TypeStyle) GetHyperlink() Hyperlink`

GetHyperlink returns the Hyperlink field if non-nil, zero value otherwise.

### GetHyperlinkOk

`func (o *TypeStyle) GetHyperlinkOk() (*Hyperlink, bool)`

GetHyperlinkOk returns a tuple with the Hyperlink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHyperlink

`func (o *TypeStyle) SetHyperlink(v Hyperlink)`

SetHyperlink sets Hyperlink field to given value.

### HasHyperlink

`func (o *TypeStyle) HasHyperlink() bool`

HasHyperlink returns a boolean if a field has been set.

### GetOpentypeFlags

`func (o *TypeStyle) GetOpentypeFlags() map[string]float32`

GetOpentypeFlags returns the OpentypeFlags field if non-nil, zero value otherwise.

### GetOpentypeFlagsOk

`func (o *TypeStyle) GetOpentypeFlagsOk() (*map[string]float32, bool)`

GetOpentypeFlagsOk returns a tuple with the OpentypeFlags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpentypeFlags

`func (o *TypeStyle) SetOpentypeFlags(v map[string]float32)`

SetOpentypeFlags sets OpentypeFlags field to given value.

### HasOpentypeFlags

`func (o *TypeStyle) HasOpentypeFlags() bool`

HasOpentypeFlags returns a boolean if a field has been set.

### GetLineHeightPx

`func (o *TypeStyle) GetLineHeightPx() float32`

GetLineHeightPx returns the LineHeightPx field if non-nil, zero value otherwise.

### GetLineHeightPxOk

`func (o *TypeStyle) GetLineHeightPxOk() (*float32, bool)`

GetLineHeightPxOk returns a tuple with the LineHeightPx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineHeightPx

`func (o *TypeStyle) SetLineHeightPx(v float32)`

SetLineHeightPx sets LineHeightPx field to given value.

### HasLineHeightPx

`func (o *TypeStyle) HasLineHeightPx() bool`

HasLineHeightPx returns a boolean if a field has been set.

### GetLineHeightPercent

`func (o *TypeStyle) GetLineHeightPercent() float32`

GetLineHeightPercent returns the LineHeightPercent field if non-nil, zero value otherwise.

### GetLineHeightPercentOk

`func (o *TypeStyle) GetLineHeightPercentOk() (*float32, bool)`

GetLineHeightPercentOk returns a tuple with the LineHeightPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineHeightPercent

`func (o *TypeStyle) SetLineHeightPercent(v float32)`

SetLineHeightPercent sets LineHeightPercent field to given value.

### HasLineHeightPercent

`func (o *TypeStyle) HasLineHeightPercent() bool`

HasLineHeightPercent returns a boolean if a field has been set.

### GetLineHeightPercentFontSize

`func (o *TypeStyle) GetLineHeightPercentFontSize() float32`

GetLineHeightPercentFontSize returns the LineHeightPercentFontSize field if non-nil, zero value otherwise.

### GetLineHeightPercentFontSizeOk

`func (o *TypeStyle) GetLineHeightPercentFontSizeOk() (*float32, bool)`

GetLineHeightPercentFontSizeOk returns a tuple with the LineHeightPercentFontSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineHeightPercentFontSize

`func (o *TypeStyle) SetLineHeightPercentFontSize(v float32)`

SetLineHeightPercentFontSize sets LineHeightPercentFontSize field to given value.

### HasLineHeightPercentFontSize

`func (o *TypeStyle) HasLineHeightPercentFontSize() bool`

HasLineHeightPercentFontSize returns a boolean if a field has been set.

### GetLineHeightUnit

`func (o *TypeStyle) GetLineHeightUnit() string`

GetLineHeightUnit returns the LineHeightUnit field if non-nil, zero value otherwise.

### GetLineHeightUnitOk

`func (o *TypeStyle) GetLineHeightUnitOk() (*string, bool)`

GetLineHeightUnitOk returns a tuple with the LineHeightUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineHeightUnit

`func (o *TypeStyle) SetLineHeightUnit(v string)`

SetLineHeightUnit sets LineHeightUnit field to given value.

### HasLineHeightUnit

`func (o *TypeStyle) HasLineHeightUnit() bool`

HasLineHeightUnit returns a boolean if a field has been set.

### GetBoundVariables

`func (o *TypeStyle) GetBoundVariables() TypeStyleBoundVariables`

GetBoundVariables returns the BoundVariables field if non-nil, zero value otherwise.

### GetBoundVariablesOk

`func (o *TypeStyle) GetBoundVariablesOk() (*TypeStyleBoundVariables, bool)`

GetBoundVariablesOk returns a tuple with the BoundVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundVariables

`func (o *TypeStyle) SetBoundVariables(v TypeStyleBoundVariables)`

SetBoundVariables sets BoundVariables field to given value.

### HasBoundVariables

`func (o *TypeStyle) HasBoundVariables() bool`

HasBoundVariables returns a boolean if a field has been set.

### GetIsOverrideOverTextStyle

`func (o *TypeStyle) GetIsOverrideOverTextStyle() bool`

GetIsOverrideOverTextStyle returns the IsOverrideOverTextStyle field if non-nil, zero value otherwise.

### GetIsOverrideOverTextStyleOk

`func (o *TypeStyle) GetIsOverrideOverTextStyleOk() (*bool, bool)`

GetIsOverrideOverTextStyleOk returns a tuple with the IsOverrideOverTextStyle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOverrideOverTextStyle

`func (o *TypeStyle) SetIsOverrideOverTextStyle(v bool)`

SetIsOverrideOverTextStyle sets IsOverrideOverTextStyle field to given value.

### HasIsOverrideOverTextStyle

`func (o *TypeStyle) HasIsOverrideOverTextStyle() bool`

HasIsOverrideOverTextStyle returns a boolean if a field has been set.

### GetSemanticWeight

`func (o *TypeStyle) GetSemanticWeight() string`

GetSemanticWeight returns the SemanticWeight field if non-nil, zero value otherwise.

### GetSemanticWeightOk

`func (o *TypeStyle) GetSemanticWeightOk() (*string, bool)`

GetSemanticWeightOk returns a tuple with the SemanticWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSemanticWeight

`func (o *TypeStyle) SetSemanticWeight(v string)`

SetSemanticWeight sets SemanticWeight field to given value.

### HasSemanticWeight

`func (o *TypeStyle) HasSemanticWeight() bool`

HasSemanticWeight returns a boolean if a field has been set.

### GetSemanticItalic

`func (o *TypeStyle) GetSemanticItalic() string`

GetSemanticItalic returns the SemanticItalic field if non-nil, zero value otherwise.

### GetSemanticItalicOk

`func (o *TypeStyle) GetSemanticItalicOk() (*string, bool)`

GetSemanticItalicOk returns a tuple with the SemanticItalic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSemanticItalic

`func (o *TypeStyle) SetSemanticItalic(v string)`

SetSemanticItalic sets SemanticItalic field to given value.

### HasSemanticItalic

`func (o *TypeStyle) HasSemanticItalic() bool`

HasSemanticItalic returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


