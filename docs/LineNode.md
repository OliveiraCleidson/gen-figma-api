# LineNode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | A string uniquely identifying this node within the document. | 
**Name** | **string** | The name given to the node by the user in the tool. | 
**Type** | **string** | The type of this node, represented by the string literal \&quot;LINE\&quot; | 
**Visible** | Pointer to **bool** | Whether or not the node is visible on the canvas. | [optional] [default to true]
**Locked** | Pointer to **bool** | If true, layer is locked and cannot be edited | [optional] [default to false]
**IsFixed** | Pointer to **bool** | Whether the layer is fixed while the parent is scrolling | [optional] [default to false]
**ScrollBehavior** | **string** | How layer should be treated when the frame is resized | [default to "SCROLLS"]
**Rotation** | Pointer to **float32** | The rotation of the node, if not 0. | [optional] [default to 0]
**ComponentPropertyReferences** | Pointer to **map[string]string** | A mapping of a layer&#39;s property to component property name of component properties attached to this node. The component property name can be used to look up more information on the corresponding component&#39;s or component set&#39;s componentPropertyDefinitions. | [optional] 
**PluginData** | Pointer to **interface{}** |  | [optional] 
**SharedPluginData** | Pointer to **interface{}** |  | [optional] 
**BoundVariables** | Pointer to [**IsLayerTraitBoundVariables**](IsLayerTraitBoundVariables.md) |  | [optional] 
**ExplicitVariableModes** | Pointer to **map[string]string** | A mapping of variable collection ID to mode ID representing the explicitly set modes for this node. | [optional] 
**BlendMode** | [**BlendMode**](BlendMode.md) | How this node blends with nodes behind it in the scene (see blend mode section for more details) | 
**Opacity** | Pointer to **float32** | Opacity of the node | [optional] [default to 1]
**AbsoluteBoundingBox** | [**Rectangle**](Rectangle.md) |  | 
**AbsoluteRenderBounds** | [**Rectangle**](Rectangle.md) |  | 
**PreserveRatio** | Pointer to **bool** | Keep height and width constrained to same ratio. | [optional] [default to false]
**Constraints** | Pointer to [**LayoutConstraint**](LayoutConstraint.md) | Horizontal and vertical layout constraints for node. | [optional] 
**RelativeTransform** | Pointer to **[][]float32** | A transformation matrix is standard way in computer graphics to represent translation and rotation. These are the top two rows of a 3x3 matrix. The bottom row of the matrix is assumed to be [0, 0, 1]. This is known as an affine transform and is enough to represent translation, rotation, and skew.  The identity transform is [[1, 0, 0], [0, 1, 0]].  A translation matrix will typically look like:  &#x60;&#x60;&#x60; [[1, 0, tx],   [0, 1, ty]] &#x60;&#x60;&#x60;  and a rotation matrix will typically look like:  &#x60;&#x60;&#x60; [[cos(angle), sin(angle), 0],   [-sin(angle), cos(angle), 0]] &#x60;&#x60;&#x60;  Another way to think about this transform is as three vectors:  - The x axis (t[0][0], t[1][0]) - The y axis (t[0][1], t[1][1]) - The translation offset (t[0][2], t[1][2])  The most common usage of the Transform matrix is the &#x60;relativeTransform property&#x60;. This particular usage of the matrix has a few additional restrictions. The translation offset can take on any value but we do enforce that the axis vectors are unit vectors (i.e. have length 1). The axes are not required to be at 90° angles to each other. | [optional] 
**Size** | Pointer to [**Vector**](Vector.md) | Width and height of element. This is different from the width and height of the bounding box in that the absolute bounding box represents the element after scaling and rotation. Only present if &#x60;geometry&#x3D;paths&#x60; is passed. | [optional] 
**LayoutAlign** | Pointer to **string** |  Determines if the layer should stretch along the parent&#39;s counter axis. This property is only provided for direct children of auto-layout frames.  - &#x60;INHERIT&#x60; - &#x60;STRETCH&#x60;  In previous versions of auto layout, determined how the layer is aligned inside an auto-layout frame. This property is only provided for direct children of auto-layout frames.  - &#x60;MIN&#x60; - &#x60;CENTER&#x60; - &#x60;MAX&#x60; - &#x60;STRETCH&#x60;  In horizontal auto-layout frames, \&quot;MIN\&quot; and \&quot;MAX\&quot; correspond to \&quot;TOP\&quot; and \&quot;BOTTOM\&quot;. In vertical auto-layout frames, \&quot;MIN\&quot; and \&quot;MAX\&quot; correspond to \&quot;LEFT\&quot; and \&quot;RIGHT\&quot;. | [optional] 
**LayoutGrow** | Pointer to **float32** | This property is applicable only for direct children of auto-layout frames, ignored otherwise. Determines whether a layer should stretch along the parent&#39;s primary axis. A &#x60;0&#x60; corresponds to a fixed size and &#x60;1&#x60; corresponds to stretch. | [optional] [default to 0]
**LayoutPositioning** | Pointer to **string** | Determines whether a layer&#39;s size and position should be determined by auto-layout settings or manually adjustable. | [optional] [default to "AUTO"]
**MinWidth** | Pointer to **float32** | The minimum width of the frame. This property is only applicable for auto-layout frames or direct children of auto-layout frames. | [optional] [default to 0]
**MaxWidth** | Pointer to **float32** | The maximum width of the frame. This property is only applicable for auto-layout frames or direct children of auto-layout frames. | [optional] [default to 0]
**MinHeight** | Pointer to **float32** | The minimum height of the frame. This property is only applicable for auto-layout frames or direct children of auto-layout frames. | [optional] [default to 0]
**MaxHeight** | Pointer to **float32** | The maximum height of the frame. This property is only applicable for auto-layout frames or direct children of auto-layout frames. | [optional] [default to 0]
**LayoutSizingHorizontal** | Pointer to **string** | The horizontal sizing setting on this auto-layout frame or frame child. - &#x60;FIXED&#x60; - &#x60;HUG&#x60;: only valid on auto-layout frames and text nodes - &#x60;FILL&#x60;: only valid on auto-layout frame children | [optional] 
**LayoutSizingVertical** | Pointer to **string** | The vertical sizing setting on this auto-layout frame or frame child. - &#x60;FIXED&#x60; - &#x60;HUG&#x60;: only valid on auto-layout frames and text nodes - &#x60;FILL&#x60;: only valid on auto-layout frame children | [optional] 
**Fills** | [**[]Paint**](Paint.md) | An array of fill paints applied to the node. | 
**Styles** | Pointer to **map[string]string** | A mapping of a StyleType to style ID (see Style) of styles present on this node. The style ID can be used to look up more information about the style in the top-level styles field. | [optional] 
**Strokes** | Pointer to [**[]Paint**](Paint.md) | An array of stroke paints applied to the node. | [optional] 
**StrokeWeight** | Pointer to **float32** | The weight of strokes on the node. | [optional] [default to 1]
**StrokeAlign** | Pointer to **string** | Position of stroke relative to vector outline, as a string enum  - &#x60;INSIDE&#x60;: stroke drawn inside the shape boundary - &#x60;OUTSIDE&#x60;: stroke drawn outside the shape boundary - &#x60;CENTER&#x60;: stroke drawn centered along the shape boundary | [optional] 
**StrokeJoin** | Pointer to **string** | A string enum with value of \&quot;MITER\&quot;, \&quot;BEVEL\&quot;, or \&quot;ROUND\&quot;, describing how corners in vector paths are rendered. | [optional] [default to "MITER"]
**StrokeDashes** | Pointer to **[]float32** | An array of floating point numbers describing the pattern of dash length and gap lengths that the vector stroke will use when drawn.  For example a value of [1, 2] indicates that the stroke will be drawn with a dash of length 1 followed by a gap of length 2, repeated. | [optional] 
**FillOverrideTable** | Pointer to [**map[string]HasGeometryTraitAllOfFillOverrideTable**](HasGeometryTraitAllOfFillOverrideTable.md) | Map from ID to PaintOverride for looking up fill overrides. To see which regions are overriden, you must use the &#x60;geometry&#x3D;paths&#x60; option. Each path returned may have an &#x60;overrideID&#x60; which maps to this table. | [optional] 
**FillGeometry** | Pointer to [**[]Path**](Path.md) | Only specified if parameter &#x60;geometry&#x3D;paths&#x60; is used. An array of paths representing the object fill. | [optional] 
**StrokeGeometry** | Pointer to [**[]Path**](Path.md) | Only specified if parameter &#x60;geometry&#x3D;paths&#x60; is used. An array of paths representing the object stroke. | [optional] 
**StrokeCap** | Pointer to **string** | A string enum describing the end caps of vector paths. | [optional] [default to "NONE"]
**StrokeMiterAngle** | Pointer to **float32** | Only valid if &#x60;strokeJoin&#x60; is \&quot;MITER\&quot;. The corner angle, in degrees, below which &#x60;strokeJoin&#x60; will be set to \&quot;BEVEL\&quot; to avoid super sharp corners. By default this is 28.96 degrees. | [optional] [default to 28.96]
**ExportSettings** | Pointer to [**[]ExportSetting**](ExportSetting.md) | An array of export settings representing images to export from the node. | [optional] 
**Effects** | [**[]Effect**](Effect.md) | An array of effects attached to this node (see effects section for more details) | 
**IsMask** | Pointer to **bool** | Does this node mask sibling nodes in front of it? | [optional] [default to false]
**MaskType** | Pointer to **string** | If this layer is a mask, this property describes the operation used to mask the layer&#39;s siblings. The value may be one of the following:  - ALPHA: the mask node&#39;s alpha channel will be used to determine the opacity of each pixel in the masked result. - VECTOR: if the mask node has visible fill paints, every pixel inside the node&#39;s fill regions will be fully visible in the masked result. If the mask has visible stroke paints, every pixel inside the node&#39;s stroke regions will be fully visible in the masked result. - LUMINANCE: the luminance value of each pixel of the mask node will be used to determine the opacity of that pixel in the masked result. | [optional] 
**IsMaskOutline** | Pointer to **bool** | True if maskType is VECTOR. This field is deprecated; use maskType instead. | [optional] [default to false]
**TransitionNodeID** | Pointer to **string** | Node ID of node to transition to in prototyping | [optional] 
**TransitionDuration** | Pointer to **float32** | The duration of the prototyping transition on this node (in milliseconds). This will override the default transition duration on the prototype, for this node. | [optional] 
**TransitionEasing** | Pointer to [**EasingType**](EasingType.md) | The easing curve used in the prototyping transition on this node. | [optional] 
**Interactions** | Pointer to [**[]Interaction**](Interaction.md) |  | [optional] 

## Methods

### NewLineNode

`func NewLineNode(id string, name string, type_ string, scrollBehavior string, blendMode BlendMode, absoluteBoundingBox Rectangle, absoluteRenderBounds Rectangle, fills []Paint, effects []Effect, ) *LineNode`

NewLineNode instantiates a new LineNode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLineNodeWithDefaults

`func NewLineNodeWithDefaults() *LineNode`

NewLineNodeWithDefaults instantiates a new LineNode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LineNode) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LineNode) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LineNode) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *LineNode) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LineNode) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LineNode) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *LineNode) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LineNode) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *LineNode) SetType(v string)`

SetType sets Type field to given value.


### GetVisible

`func (o *LineNode) GetVisible() bool`

GetVisible returns the Visible field if non-nil, zero value otherwise.

### GetVisibleOk

`func (o *LineNode) GetVisibleOk() (*bool, bool)`

GetVisibleOk returns a tuple with the Visible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisible

`func (o *LineNode) SetVisible(v bool)`

SetVisible sets Visible field to given value.

### HasVisible

`func (o *LineNode) HasVisible() bool`

HasVisible returns a boolean if a field has been set.

### GetLocked

`func (o *LineNode) GetLocked() bool`

GetLocked returns the Locked field if non-nil, zero value otherwise.

### GetLockedOk

`func (o *LineNode) GetLockedOk() (*bool, bool)`

GetLockedOk returns a tuple with the Locked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocked

`func (o *LineNode) SetLocked(v bool)`

SetLocked sets Locked field to given value.

### HasLocked

`func (o *LineNode) HasLocked() bool`

HasLocked returns a boolean if a field has been set.

### GetIsFixed

`func (o *LineNode) GetIsFixed() bool`

GetIsFixed returns the IsFixed field if non-nil, zero value otherwise.

### GetIsFixedOk

`func (o *LineNode) GetIsFixedOk() (*bool, bool)`

GetIsFixedOk returns a tuple with the IsFixed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFixed

`func (o *LineNode) SetIsFixed(v bool)`

SetIsFixed sets IsFixed field to given value.

### HasIsFixed

`func (o *LineNode) HasIsFixed() bool`

HasIsFixed returns a boolean if a field has been set.

### GetScrollBehavior

`func (o *LineNode) GetScrollBehavior() string`

GetScrollBehavior returns the ScrollBehavior field if non-nil, zero value otherwise.

### GetScrollBehaviorOk

`func (o *LineNode) GetScrollBehaviorOk() (*string, bool)`

GetScrollBehaviorOk returns a tuple with the ScrollBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScrollBehavior

`func (o *LineNode) SetScrollBehavior(v string)`

SetScrollBehavior sets ScrollBehavior field to given value.


### GetRotation

`func (o *LineNode) GetRotation() float32`

GetRotation returns the Rotation field if non-nil, zero value otherwise.

### GetRotationOk

`func (o *LineNode) GetRotationOk() (*float32, bool)`

GetRotationOk returns a tuple with the Rotation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRotation

`func (o *LineNode) SetRotation(v float32)`

SetRotation sets Rotation field to given value.

### HasRotation

`func (o *LineNode) HasRotation() bool`

HasRotation returns a boolean if a field has been set.

### GetComponentPropertyReferences

`func (o *LineNode) GetComponentPropertyReferences() map[string]string`

GetComponentPropertyReferences returns the ComponentPropertyReferences field if non-nil, zero value otherwise.

### GetComponentPropertyReferencesOk

`func (o *LineNode) GetComponentPropertyReferencesOk() (*map[string]string, bool)`

GetComponentPropertyReferencesOk returns a tuple with the ComponentPropertyReferences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentPropertyReferences

`func (o *LineNode) SetComponentPropertyReferences(v map[string]string)`

SetComponentPropertyReferences sets ComponentPropertyReferences field to given value.

### HasComponentPropertyReferences

`func (o *LineNode) HasComponentPropertyReferences() bool`

HasComponentPropertyReferences returns a boolean if a field has been set.

### GetPluginData

`func (o *LineNode) GetPluginData() interface{}`

GetPluginData returns the PluginData field if non-nil, zero value otherwise.

### GetPluginDataOk

`func (o *LineNode) GetPluginDataOk() (*interface{}, bool)`

GetPluginDataOk returns a tuple with the PluginData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPluginData

`func (o *LineNode) SetPluginData(v interface{})`

SetPluginData sets PluginData field to given value.

### HasPluginData

`func (o *LineNode) HasPluginData() bool`

HasPluginData returns a boolean if a field has been set.

### SetPluginDataNil

`func (o *LineNode) SetPluginDataNil(b bool)`

 SetPluginDataNil sets the value for PluginData to be an explicit nil

### UnsetPluginData
`func (o *LineNode) UnsetPluginData()`

UnsetPluginData ensures that no value is present for PluginData, not even an explicit nil
### GetSharedPluginData

`func (o *LineNode) GetSharedPluginData() interface{}`

GetSharedPluginData returns the SharedPluginData field if non-nil, zero value otherwise.

### GetSharedPluginDataOk

`func (o *LineNode) GetSharedPluginDataOk() (*interface{}, bool)`

GetSharedPluginDataOk returns a tuple with the SharedPluginData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedPluginData

`func (o *LineNode) SetSharedPluginData(v interface{})`

SetSharedPluginData sets SharedPluginData field to given value.

### HasSharedPluginData

`func (o *LineNode) HasSharedPluginData() bool`

HasSharedPluginData returns a boolean if a field has been set.

### SetSharedPluginDataNil

`func (o *LineNode) SetSharedPluginDataNil(b bool)`

 SetSharedPluginDataNil sets the value for SharedPluginData to be an explicit nil

### UnsetSharedPluginData
`func (o *LineNode) UnsetSharedPluginData()`

UnsetSharedPluginData ensures that no value is present for SharedPluginData, not even an explicit nil
### GetBoundVariables

`func (o *LineNode) GetBoundVariables() IsLayerTraitBoundVariables`

GetBoundVariables returns the BoundVariables field if non-nil, zero value otherwise.

### GetBoundVariablesOk

`func (o *LineNode) GetBoundVariablesOk() (*IsLayerTraitBoundVariables, bool)`

GetBoundVariablesOk returns a tuple with the BoundVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundVariables

`func (o *LineNode) SetBoundVariables(v IsLayerTraitBoundVariables)`

SetBoundVariables sets BoundVariables field to given value.

### HasBoundVariables

`func (o *LineNode) HasBoundVariables() bool`

HasBoundVariables returns a boolean if a field has been set.

### GetExplicitVariableModes

`func (o *LineNode) GetExplicitVariableModes() map[string]string`

GetExplicitVariableModes returns the ExplicitVariableModes field if non-nil, zero value otherwise.

### GetExplicitVariableModesOk

`func (o *LineNode) GetExplicitVariableModesOk() (*map[string]string, bool)`

GetExplicitVariableModesOk returns a tuple with the ExplicitVariableModes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExplicitVariableModes

`func (o *LineNode) SetExplicitVariableModes(v map[string]string)`

SetExplicitVariableModes sets ExplicitVariableModes field to given value.

### HasExplicitVariableModes

`func (o *LineNode) HasExplicitVariableModes() bool`

HasExplicitVariableModes returns a boolean if a field has been set.

### GetBlendMode

`func (o *LineNode) GetBlendMode() BlendMode`

GetBlendMode returns the BlendMode field if non-nil, zero value otherwise.

### GetBlendModeOk

`func (o *LineNode) GetBlendModeOk() (*BlendMode, bool)`

GetBlendModeOk returns a tuple with the BlendMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlendMode

`func (o *LineNode) SetBlendMode(v BlendMode)`

SetBlendMode sets BlendMode field to given value.


### GetOpacity

`func (o *LineNode) GetOpacity() float32`

GetOpacity returns the Opacity field if non-nil, zero value otherwise.

### GetOpacityOk

`func (o *LineNode) GetOpacityOk() (*float32, bool)`

GetOpacityOk returns a tuple with the Opacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpacity

`func (o *LineNode) SetOpacity(v float32)`

SetOpacity sets Opacity field to given value.

### HasOpacity

`func (o *LineNode) HasOpacity() bool`

HasOpacity returns a boolean if a field has been set.

### GetAbsoluteBoundingBox

`func (o *LineNode) GetAbsoluteBoundingBox() Rectangle`

GetAbsoluteBoundingBox returns the AbsoluteBoundingBox field if non-nil, zero value otherwise.

### GetAbsoluteBoundingBoxOk

`func (o *LineNode) GetAbsoluteBoundingBoxOk() (*Rectangle, bool)`

GetAbsoluteBoundingBoxOk returns a tuple with the AbsoluteBoundingBox field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbsoluteBoundingBox

`func (o *LineNode) SetAbsoluteBoundingBox(v Rectangle)`

SetAbsoluteBoundingBox sets AbsoluteBoundingBox field to given value.


### GetAbsoluteRenderBounds

`func (o *LineNode) GetAbsoluteRenderBounds() Rectangle`

GetAbsoluteRenderBounds returns the AbsoluteRenderBounds field if non-nil, zero value otherwise.

### GetAbsoluteRenderBoundsOk

`func (o *LineNode) GetAbsoluteRenderBoundsOk() (*Rectangle, bool)`

GetAbsoluteRenderBoundsOk returns a tuple with the AbsoluteRenderBounds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbsoluteRenderBounds

`func (o *LineNode) SetAbsoluteRenderBounds(v Rectangle)`

SetAbsoluteRenderBounds sets AbsoluteRenderBounds field to given value.


### GetPreserveRatio

`func (o *LineNode) GetPreserveRatio() bool`

GetPreserveRatio returns the PreserveRatio field if non-nil, zero value otherwise.

### GetPreserveRatioOk

`func (o *LineNode) GetPreserveRatioOk() (*bool, bool)`

GetPreserveRatioOk returns a tuple with the PreserveRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreserveRatio

`func (o *LineNode) SetPreserveRatio(v bool)`

SetPreserveRatio sets PreserveRatio field to given value.

### HasPreserveRatio

`func (o *LineNode) HasPreserveRatio() bool`

HasPreserveRatio returns a boolean if a field has been set.

### GetConstraints

`func (o *LineNode) GetConstraints() LayoutConstraint`

GetConstraints returns the Constraints field if non-nil, zero value otherwise.

### GetConstraintsOk

`func (o *LineNode) GetConstraintsOk() (*LayoutConstraint, bool)`

GetConstraintsOk returns a tuple with the Constraints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConstraints

`func (o *LineNode) SetConstraints(v LayoutConstraint)`

SetConstraints sets Constraints field to given value.

### HasConstraints

`func (o *LineNode) HasConstraints() bool`

HasConstraints returns a boolean if a field has been set.

### GetRelativeTransform

`func (o *LineNode) GetRelativeTransform() [][]float32`

GetRelativeTransform returns the RelativeTransform field if non-nil, zero value otherwise.

### GetRelativeTransformOk

`func (o *LineNode) GetRelativeTransformOk() (*[][]float32, bool)`

GetRelativeTransformOk returns a tuple with the RelativeTransform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelativeTransform

`func (o *LineNode) SetRelativeTransform(v [][]float32)`

SetRelativeTransform sets RelativeTransform field to given value.

### HasRelativeTransform

`func (o *LineNode) HasRelativeTransform() bool`

HasRelativeTransform returns a boolean if a field has been set.

### GetSize

`func (o *LineNode) GetSize() Vector`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *LineNode) GetSizeOk() (*Vector, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *LineNode) SetSize(v Vector)`

SetSize sets Size field to given value.

### HasSize

`func (o *LineNode) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetLayoutAlign

`func (o *LineNode) GetLayoutAlign() string`

GetLayoutAlign returns the LayoutAlign field if non-nil, zero value otherwise.

### GetLayoutAlignOk

`func (o *LineNode) GetLayoutAlignOk() (*string, bool)`

GetLayoutAlignOk returns a tuple with the LayoutAlign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayoutAlign

`func (o *LineNode) SetLayoutAlign(v string)`

SetLayoutAlign sets LayoutAlign field to given value.

### HasLayoutAlign

`func (o *LineNode) HasLayoutAlign() bool`

HasLayoutAlign returns a boolean if a field has been set.

### GetLayoutGrow

`func (o *LineNode) GetLayoutGrow() float32`

GetLayoutGrow returns the LayoutGrow field if non-nil, zero value otherwise.

### GetLayoutGrowOk

`func (o *LineNode) GetLayoutGrowOk() (*float32, bool)`

GetLayoutGrowOk returns a tuple with the LayoutGrow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayoutGrow

`func (o *LineNode) SetLayoutGrow(v float32)`

SetLayoutGrow sets LayoutGrow field to given value.

### HasLayoutGrow

`func (o *LineNode) HasLayoutGrow() bool`

HasLayoutGrow returns a boolean if a field has been set.

### GetLayoutPositioning

`func (o *LineNode) GetLayoutPositioning() string`

GetLayoutPositioning returns the LayoutPositioning field if non-nil, zero value otherwise.

### GetLayoutPositioningOk

`func (o *LineNode) GetLayoutPositioningOk() (*string, bool)`

GetLayoutPositioningOk returns a tuple with the LayoutPositioning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayoutPositioning

`func (o *LineNode) SetLayoutPositioning(v string)`

SetLayoutPositioning sets LayoutPositioning field to given value.

### HasLayoutPositioning

`func (o *LineNode) HasLayoutPositioning() bool`

HasLayoutPositioning returns a boolean if a field has been set.

### GetMinWidth

`func (o *LineNode) GetMinWidth() float32`

GetMinWidth returns the MinWidth field if non-nil, zero value otherwise.

### GetMinWidthOk

`func (o *LineNode) GetMinWidthOk() (*float32, bool)`

GetMinWidthOk returns a tuple with the MinWidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinWidth

`func (o *LineNode) SetMinWidth(v float32)`

SetMinWidth sets MinWidth field to given value.

### HasMinWidth

`func (o *LineNode) HasMinWidth() bool`

HasMinWidth returns a boolean if a field has been set.

### GetMaxWidth

`func (o *LineNode) GetMaxWidth() float32`

GetMaxWidth returns the MaxWidth field if non-nil, zero value otherwise.

### GetMaxWidthOk

`func (o *LineNode) GetMaxWidthOk() (*float32, bool)`

GetMaxWidthOk returns a tuple with the MaxWidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxWidth

`func (o *LineNode) SetMaxWidth(v float32)`

SetMaxWidth sets MaxWidth field to given value.

### HasMaxWidth

`func (o *LineNode) HasMaxWidth() bool`

HasMaxWidth returns a boolean if a field has been set.

### GetMinHeight

`func (o *LineNode) GetMinHeight() float32`

GetMinHeight returns the MinHeight field if non-nil, zero value otherwise.

### GetMinHeightOk

`func (o *LineNode) GetMinHeightOk() (*float32, bool)`

GetMinHeightOk returns a tuple with the MinHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinHeight

`func (o *LineNode) SetMinHeight(v float32)`

SetMinHeight sets MinHeight field to given value.

### HasMinHeight

`func (o *LineNode) HasMinHeight() bool`

HasMinHeight returns a boolean if a field has been set.

### GetMaxHeight

`func (o *LineNode) GetMaxHeight() float32`

GetMaxHeight returns the MaxHeight field if non-nil, zero value otherwise.

### GetMaxHeightOk

`func (o *LineNode) GetMaxHeightOk() (*float32, bool)`

GetMaxHeightOk returns a tuple with the MaxHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxHeight

`func (o *LineNode) SetMaxHeight(v float32)`

SetMaxHeight sets MaxHeight field to given value.

### HasMaxHeight

`func (o *LineNode) HasMaxHeight() bool`

HasMaxHeight returns a boolean if a field has been set.

### GetLayoutSizingHorizontal

`func (o *LineNode) GetLayoutSizingHorizontal() string`

GetLayoutSizingHorizontal returns the LayoutSizingHorizontal field if non-nil, zero value otherwise.

### GetLayoutSizingHorizontalOk

`func (o *LineNode) GetLayoutSizingHorizontalOk() (*string, bool)`

GetLayoutSizingHorizontalOk returns a tuple with the LayoutSizingHorizontal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayoutSizingHorizontal

`func (o *LineNode) SetLayoutSizingHorizontal(v string)`

SetLayoutSizingHorizontal sets LayoutSizingHorizontal field to given value.

### HasLayoutSizingHorizontal

`func (o *LineNode) HasLayoutSizingHorizontal() bool`

HasLayoutSizingHorizontal returns a boolean if a field has been set.

### GetLayoutSizingVertical

`func (o *LineNode) GetLayoutSizingVertical() string`

GetLayoutSizingVertical returns the LayoutSizingVertical field if non-nil, zero value otherwise.

### GetLayoutSizingVerticalOk

`func (o *LineNode) GetLayoutSizingVerticalOk() (*string, bool)`

GetLayoutSizingVerticalOk returns a tuple with the LayoutSizingVertical field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayoutSizingVertical

`func (o *LineNode) SetLayoutSizingVertical(v string)`

SetLayoutSizingVertical sets LayoutSizingVertical field to given value.

### HasLayoutSizingVertical

`func (o *LineNode) HasLayoutSizingVertical() bool`

HasLayoutSizingVertical returns a boolean if a field has been set.

### GetFills

`func (o *LineNode) GetFills() []Paint`

GetFills returns the Fills field if non-nil, zero value otherwise.

### GetFillsOk

`func (o *LineNode) GetFillsOk() (*[]Paint, bool)`

GetFillsOk returns a tuple with the Fills field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFills

`func (o *LineNode) SetFills(v []Paint)`

SetFills sets Fills field to given value.


### GetStyles

`func (o *LineNode) GetStyles() map[string]string`

GetStyles returns the Styles field if non-nil, zero value otherwise.

### GetStylesOk

`func (o *LineNode) GetStylesOk() (*map[string]string, bool)`

GetStylesOk returns a tuple with the Styles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStyles

`func (o *LineNode) SetStyles(v map[string]string)`

SetStyles sets Styles field to given value.

### HasStyles

`func (o *LineNode) HasStyles() bool`

HasStyles returns a boolean if a field has been set.

### GetStrokes

`func (o *LineNode) GetStrokes() []Paint`

GetStrokes returns the Strokes field if non-nil, zero value otherwise.

### GetStrokesOk

`func (o *LineNode) GetStrokesOk() (*[]Paint, bool)`

GetStrokesOk returns a tuple with the Strokes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrokes

`func (o *LineNode) SetStrokes(v []Paint)`

SetStrokes sets Strokes field to given value.

### HasStrokes

`func (o *LineNode) HasStrokes() bool`

HasStrokes returns a boolean if a field has been set.

### GetStrokeWeight

`func (o *LineNode) GetStrokeWeight() float32`

GetStrokeWeight returns the StrokeWeight field if non-nil, zero value otherwise.

### GetStrokeWeightOk

`func (o *LineNode) GetStrokeWeightOk() (*float32, bool)`

GetStrokeWeightOk returns a tuple with the StrokeWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrokeWeight

`func (o *LineNode) SetStrokeWeight(v float32)`

SetStrokeWeight sets StrokeWeight field to given value.

### HasStrokeWeight

`func (o *LineNode) HasStrokeWeight() bool`

HasStrokeWeight returns a boolean if a field has been set.

### GetStrokeAlign

`func (o *LineNode) GetStrokeAlign() string`

GetStrokeAlign returns the StrokeAlign field if non-nil, zero value otherwise.

### GetStrokeAlignOk

`func (o *LineNode) GetStrokeAlignOk() (*string, bool)`

GetStrokeAlignOk returns a tuple with the StrokeAlign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrokeAlign

`func (o *LineNode) SetStrokeAlign(v string)`

SetStrokeAlign sets StrokeAlign field to given value.

### HasStrokeAlign

`func (o *LineNode) HasStrokeAlign() bool`

HasStrokeAlign returns a boolean if a field has been set.

### GetStrokeJoin

`func (o *LineNode) GetStrokeJoin() string`

GetStrokeJoin returns the StrokeJoin field if non-nil, zero value otherwise.

### GetStrokeJoinOk

`func (o *LineNode) GetStrokeJoinOk() (*string, bool)`

GetStrokeJoinOk returns a tuple with the StrokeJoin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrokeJoin

`func (o *LineNode) SetStrokeJoin(v string)`

SetStrokeJoin sets StrokeJoin field to given value.

### HasStrokeJoin

`func (o *LineNode) HasStrokeJoin() bool`

HasStrokeJoin returns a boolean if a field has been set.

### GetStrokeDashes

`func (o *LineNode) GetStrokeDashes() []float32`

GetStrokeDashes returns the StrokeDashes field if non-nil, zero value otherwise.

### GetStrokeDashesOk

`func (o *LineNode) GetStrokeDashesOk() (*[]float32, bool)`

GetStrokeDashesOk returns a tuple with the StrokeDashes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrokeDashes

`func (o *LineNode) SetStrokeDashes(v []float32)`

SetStrokeDashes sets StrokeDashes field to given value.

### HasStrokeDashes

`func (o *LineNode) HasStrokeDashes() bool`

HasStrokeDashes returns a boolean if a field has been set.

### GetFillOverrideTable

`func (o *LineNode) GetFillOverrideTable() map[string]HasGeometryTraitAllOfFillOverrideTable`

GetFillOverrideTable returns the FillOverrideTable field if non-nil, zero value otherwise.

### GetFillOverrideTableOk

`func (o *LineNode) GetFillOverrideTableOk() (*map[string]HasGeometryTraitAllOfFillOverrideTable, bool)`

GetFillOverrideTableOk returns a tuple with the FillOverrideTable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFillOverrideTable

`func (o *LineNode) SetFillOverrideTable(v map[string]HasGeometryTraitAllOfFillOverrideTable)`

SetFillOverrideTable sets FillOverrideTable field to given value.

### HasFillOverrideTable

`func (o *LineNode) HasFillOverrideTable() bool`

HasFillOverrideTable returns a boolean if a field has been set.

### GetFillGeometry

`func (o *LineNode) GetFillGeometry() []Path`

GetFillGeometry returns the FillGeometry field if non-nil, zero value otherwise.

### GetFillGeometryOk

`func (o *LineNode) GetFillGeometryOk() (*[]Path, bool)`

GetFillGeometryOk returns a tuple with the FillGeometry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFillGeometry

`func (o *LineNode) SetFillGeometry(v []Path)`

SetFillGeometry sets FillGeometry field to given value.

### HasFillGeometry

`func (o *LineNode) HasFillGeometry() bool`

HasFillGeometry returns a boolean if a field has been set.

### GetStrokeGeometry

`func (o *LineNode) GetStrokeGeometry() []Path`

GetStrokeGeometry returns the StrokeGeometry field if non-nil, zero value otherwise.

### GetStrokeGeometryOk

`func (o *LineNode) GetStrokeGeometryOk() (*[]Path, bool)`

GetStrokeGeometryOk returns a tuple with the StrokeGeometry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrokeGeometry

`func (o *LineNode) SetStrokeGeometry(v []Path)`

SetStrokeGeometry sets StrokeGeometry field to given value.

### HasStrokeGeometry

`func (o *LineNode) HasStrokeGeometry() bool`

HasStrokeGeometry returns a boolean if a field has been set.

### GetStrokeCap

`func (o *LineNode) GetStrokeCap() string`

GetStrokeCap returns the StrokeCap field if non-nil, zero value otherwise.

### GetStrokeCapOk

`func (o *LineNode) GetStrokeCapOk() (*string, bool)`

GetStrokeCapOk returns a tuple with the StrokeCap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrokeCap

`func (o *LineNode) SetStrokeCap(v string)`

SetStrokeCap sets StrokeCap field to given value.

### HasStrokeCap

`func (o *LineNode) HasStrokeCap() bool`

HasStrokeCap returns a boolean if a field has been set.

### GetStrokeMiterAngle

`func (o *LineNode) GetStrokeMiterAngle() float32`

GetStrokeMiterAngle returns the StrokeMiterAngle field if non-nil, zero value otherwise.

### GetStrokeMiterAngleOk

`func (o *LineNode) GetStrokeMiterAngleOk() (*float32, bool)`

GetStrokeMiterAngleOk returns a tuple with the StrokeMiterAngle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrokeMiterAngle

`func (o *LineNode) SetStrokeMiterAngle(v float32)`

SetStrokeMiterAngle sets StrokeMiterAngle field to given value.

### HasStrokeMiterAngle

`func (o *LineNode) HasStrokeMiterAngle() bool`

HasStrokeMiterAngle returns a boolean if a field has been set.

### GetExportSettings

`func (o *LineNode) GetExportSettings() []ExportSetting`

GetExportSettings returns the ExportSettings field if non-nil, zero value otherwise.

### GetExportSettingsOk

`func (o *LineNode) GetExportSettingsOk() (*[]ExportSetting, bool)`

GetExportSettingsOk returns a tuple with the ExportSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportSettings

`func (o *LineNode) SetExportSettings(v []ExportSetting)`

SetExportSettings sets ExportSettings field to given value.

### HasExportSettings

`func (o *LineNode) HasExportSettings() bool`

HasExportSettings returns a boolean if a field has been set.

### GetEffects

`func (o *LineNode) GetEffects() []Effect`

GetEffects returns the Effects field if non-nil, zero value otherwise.

### GetEffectsOk

`func (o *LineNode) GetEffectsOk() (*[]Effect, bool)`

GetEffectsOk returns a tuple with the Effects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffects

`func (o *LineNode) SetEffects(v []Effect)`

SetEffects sets Effects field to given value.


### GetIsMask

`func (o *LineNode) GetIsMask() bool`

GetIsMask returns the IsMask field if non-nil, zero value otherwise.

### GetIsMaskOk

`func (o *LineNode) GetIsMaskOk() (*bool, bool)`

GetIsMaskOk returns a tuple with the IsMask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMask

`func (o *LineNode) SetIsMask(v bool)`

SetIsMask sets IsMask field to given value.

### HasIsMask

`func (o *LineNode) HasIsMask() bool`

HasIsMask returns a boolean if a field has been set.

### GetMaskType

`func (o *LineNode) GetMaskType() string`

GetMaskType returns the MaskType field if non-nil, zero value otherwise.

### GetMaskTypeOk

`func (o *LineNode) GetMaskTypeOk() (*string, bool)`

GetMaskTypeOk returns a tuple with the MaskType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaskType

`func (o *LineNode) SetMaskType(v string)`

SetMaskType sets MaskType field to given value.

### HasMaskType

`func (o *LineNode) HasMaskType() bool`

HasMaskType returns a boolean if a field has been set.

### GetIsMaskOutline

`func (o *LineNode) GetIsMaskOutline() bool`

GetIsMaskOutline returns the IsMaskOutline field if non-nil, zero value otherwise.

### GetIsMaskOutlineOk

`func (o *LineNode) GetIsMaskOutlineOk() (*bool, bool)`

GetIsMaskOutlineOk returns a tuple with the IsMaskOutline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMaskOutline

`func (o *LineNode) SetIsMaskOutline(v bool)`

SetIsMaskOutline sets IsMaskOutline field to given value.

### HasIsMaskOutline

`func (o *LineNode) HasIsMaskOutline() bool`

HasIsMaskOutline returns a boolean if a field has been set.

### GetTransitionNodeID

`func (o *LineNode) GetTransitionNodeID() string`

GetTransitionNodeID returns the TransitionNodeID field if non-nil, zero value otherwise.

### GetTransitionNodeIDOk

`func (o *LineNode) GetTransitionNodeIDOk() (*string, bool)`

GetTransitionNodeIDOk returns a tuple with the TransitionNodeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransitionNodeID

`func (o *LineNode) SetTransitionNodeID(v string)`

SetTransitionNodeID sets TransitionNodeID field to given value.

### HasTransitionNodeID

`func (o *LineNode) HasTransitionNodeID() bool`

HasTransitionNodeID returns a boolean if a field has been set.

### GetTransitionDuration

`func (o *LineNode) GetTransitionDuration() float32`

GetTransitionDuration returns the TransitionDuration field if non-nil, zero value otherwise.

### GetTransitionDurationOk

`func (o *LineNode) GetTransitionDurationOk() (*float32, bool)`

GetTransitionDurationOk returns a tuple with the TransitionDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransitionDuration

`func (o *LineNode) SetTransitionDuration(v float32)`

SetTransitionDuration sets TransitionDuration field to given value.

### HasTransitionDuration

`func (o *LineNode) HasTransitionDuration() bool`

HasTransitionDuration returns a boolean if a field has been set.

### GetTransitionEasing

`func (o *LineNode) GetTransitionEasing() EasingType`

GetTransitionEasing returns the TransitionEasing field if non-nil, zero value otherwise.

### GetTransitionEasingOk

`func (o *LineNode) GetTransitionEasingOk() (*EasingType, bool)`

GetTransitionEasingOk returns a tuple with the TransitionEasing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransitionEasing

`func (o *LineNode) SetTransitionEasing(v EasingType)`

SetTransitionEasing sets TransitionEasing field to given value.

### HasTransitionEasing

`func (o *LineNode) HasTransitionEasing() bool`

HasTransitionEasing returns a boolean if a field has been set.

### GetInteractions

`func (o *LineNode) GetInteractions() []Interaction`

GetInteractions returns the Interactions field if non-nil, zero value otherwise.

### GetInteractionsOk

`func (o *LineNode) GetInteractionsOk() (*[]Interaction, bool)`

GetInteractionsOk returns a tuple with the Interactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInteractions

`func (o *LineNode) SetInteractions(v []Interaction)`

SetInteractions sets Interactions field to given value.

### HasInteractions

`func (o *LineNode) HasInteractions() bool`

HasInteractions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


