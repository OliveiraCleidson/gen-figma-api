# TypePropertiesTrait

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Characters** | **string** | The raw characters in the text node. | 
**Style** | [**TypeStyle**](TypeStyle.md) | Style of text including font family and weight. | 
**CharacterStyleOverrides** | **[]float32** | The array corresponds to characters in the text box, where each element references the &#39;styleOverrideTable&#39; to apply specific styles to each character. The array&#39;s length can be less than or equal to the number of characters due to the removal of trailing zeros. Elements with a value of 0 indicate characters that use the default type style. If the array is shorter than the total number of characters, the characters beyond the array&#39;s length also use the default style. | 
**LayoutVersion** | Pointer to **float32** | Internal property, preserved for backward compatibility. Avoid using this value. | [optional] 
**StyleOverrideTable** | [**map[string]TypeStyle**](TypeStyle.md) | Map from ID to TypeStyle for looking up style overrides. | 
**LineTypes** | **[]string** | An array with the same number of elements as lines in the text node, where lines are delimited by newline or paragraph separator characters. Each element in the array corresponds to the list type of a specific line. List types are represented as string enums with one of these possible values:  - &#x60;NONE&#x60;: Not a list item. - &#x60;ORDERED&#x60;: Text is an ordered list (numbered). - &#x60;UNORDERED&#x60;: Text is an unordered list (bulleted). | 
**LineIndentations** | **[]float32** | An array with the same number of elements as lines in the text node, where lines are delimited by newline or paragraph separator characters. Each element in the array corresponds to the indentation level of a specific line. | 

## Methods

### NewTypePropertiesTrait

`func NewTypePropertiesTrait(characters string, style TypeStyle, characterStyleOverrides []float32, styleOverrideTable map[string]TypeStyle, lineTypes []string, lineIndentations []float32, ) *TypePropertiesTrait`

NewTypePropertiesTrait instantiates a new TypePropertiesTrait object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTypePropertiesTraitWithDefaults

`func NewTypePropertiesTraitWithDefaults() *TypePropertiesTrait`

NewTypePropertiesTraitWithDefaults instantiates a new TypePropertiesTrait object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCharacters

`func (o *TypePropertiesTrait) GetCharacters() string`

GetCharacters returns the Characters field if non-nil, zero value otherwise.

### GetCharactersOk

`func (o *TypePropertiesTrait) GetCharactersOk() (*string, bool)`

GetCharactersOk returns a tuple with the Characters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCharacters

`func (o *TypePropertiesTrait) SetCharacters(v string)`

SetCharacters sets Characters field to given value.


### GetStyle

`func (o *TypePropertiesTrait) GetStyle() TypeStyle`

GetStyle returns the Style field if non-nil, zero value otherwise.

### GetStyleOk

`func (o *TypePropertiesTrait) GetStyleOk() (*TypeStyle, bool)`

GetStyleOk returns a tuple with the Style field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStyle

`func (o *TypePropertiesTrait) SetStyle(v TypeStyle)`

SetStyle sets Style field to given value.


### GetCharacterStyleOverrides

`func (o *TypePropertiesTrait) GetCharacterStyleOverrides() []float32`

GetCharacterStyleOverrides returns the CharacterStyleOverrides field if non-nil, zero value otherwise.

### GetCharacterStyleOverridesOk

`func (o *TypePropertiesTrait) GetCharacterStyleOverridesOk() (*[]float32, bool)`

GetCharacterStyleOverridesOk returns a tuple with the CharacterStyleOverrides field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCharacterStyleOverrides

`func (o *TypePropertiesTrait) SetCharacterStyleOverrides(v []float32)`

SetCharacterStyleOverrides sets CharacterStyleOverrides field to given value.


### GetLayoutVersion

`func (o *TypePropertiesTrait) GetLayoutVersion() float32`

GetLayoutVersion returns the LayoutVersion field if non-nil, zero value otherwise.

### GetLayoutVersionOk

`func (o *TypePropertiesTrait) GetLayoutVersionOk() (*float32, bool)`

GetLayoutVersionOk returns a tuple with the LayoutVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayoutVersion

`func (o *TypePropertiesTrait) SetLayoutVersion(v float32)`

SetLayoutVersion sets LayoutVersion field to given value.

### HasLayoutVersion

`func (o *TypePropertiesTrait) HasLayoutVersion() bool`

HasLayoutVersion returns a boolean if a field has been set.

### GetStyleOverrideTable

`func (o *TypePropertiesTrait) GetStyleOverrideTable() map[string]TypeStyle`

GetStyleOverrideTable returns the StyleOverrideTable field if non-nil, zero value otherwise.

### GetStyleOverrideTableOk

`func (o *TypePropertiesTrait) GetStyleOverrideTableOk() (*map[string]TypeStyle, bool)`

GetStyleOverrideTableOk returns a tuple with the StyleOverrideTable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStyleOverrideTable

`func (o *TypePropertiesTrait) SetStyleOverrideTable(v map[string]TypeStyle)`

SetStyleOverrideTable sets StyleOverrideTable field to given value.


### GetLineTypes

`func (o *TypePropertiesTrait) GetLineTypes() []string`

GetLineTypes returns the LineTypes field if non-nil, zero value otherwise.

### GetLineTypesOk

`func (o *TypePropertiesTrait) GetLineTypesOk() (*[]string, bool)`

GetLineTypesOk returns a tuple with the LineTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineTypes

`func (o *TypePropertiesTrait) SetLineTypes(v []string)`

SetLineTypes sets LineTypes field to given value.


### GetLineIndentations

`func (o *TypePropertiesTrait) GetLineIndentations() []float32`

GetLineIndentations returns the LineIndentations field if non-nil, zero value otherwise.

### GetLineIndentationsOk

`func (o *TypePropertiesTrait) GetLineIndentationsOk() (*[]float32, bool)`

GetLineIndentationsOk returns a tuple with the LineIndentations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineIndentations

`func (o *TypePropertiesTrait) SetLineIndentations(v []float32)`

SetLineIndentations sets LineIndentations field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


