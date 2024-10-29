# ComponentSetNode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | A string uniquely identifying this node within the document. | 
**Name** | **string** | The name given to the node by the user in the tool. | 
**Type** | **string** | The type of this node, represented by the string literal \&quot;COMPONENT_SET\&quot; | 
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
**Children** | [**[]SubcanvasNode**](SubcanvasNode.md) | An array of nodes that are direct children of this node | 
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
**ClipsContent** | **bool** | Whether or not this node clip content outside of its bounds | 
**Background** | Pointer to [**[]Paint**](Paint.md) | Background of the node. This is deprecated, as backgrounds for frames are now in the &#x60;fills&#x60; field. | [optional] 
**BackgroundColor** | Pointer to [**RGBA**](RGBA.md) | Background color of the node. This is deprecated, as frames now support more than a solid color as a background. Please use the &#x60;fills&#x60; field instead. | [optional] 
**LayoutGrids** | Pointer to [**[]LayoutGrid**](LayoutGrid.md) | An array of layout grids attached to this node (see layout grids section for more details). GROUP nodes do not have this attribute | [optional] 
**OverflowDirection** | Pointer to **string** | Whether a node has primary axis scrolling, horizontal or vertical. | [optional] [default to "NONE"]
**LayoutMode** | Pointer to **string** | Whether this layer uses auto-layout to position its children. | [optional] [default to "NONE"]
**PrimaryAxisSizingMode** | Pointer to **string** | Whether the primary axis has a fixed length (determined by the user) or an automatic length (determined by the layout engine). This property is only applicable for auto-layout frames. | [optional] [default to "AUTO"]
**CounterAxisSizingMode** | Pointer to **string** | Whether the counter axis has a fixed length (determined by the user) or an automatic length (determined by the layout engine). This property is only applicable for auto-layout frames. | [optional] [default to "AUTO"]
**PrimaryAxisAlignItems** | Pointer to **string** | Determines how the auto-layout frame&#39;s children should be aligned in the primary axis direction. This property is only applicable for auto-layout frames. | [optional] [default to "MIN"]
**CounterAxisAlignItems** | Pointer to **string** | Determines how the auto-layout frame&#39;s children should be aligned in the counter axis direction. This property is only applicable for auto-layout frames. | [optional] [default to "MIN"]
**PaddingLeft** | Pointer to **float32** | The padding between the left border of the frame and its children. This property is only applicable for auto-layout frames. | [optional] [default to 0]
**PaddingRight** | Pointer to **float32** | The padding between the right border of the frame and its children. This property is only applicable for auto-layout frames. | [optional] [default to 0]
**PaddingTop** | Pointer to **float32** | The padding between the top border of the frame and its children. This property is only applicable for auto-layout frames. | [optional] [default to 0]
**PaddingBottom** | Pointer to **float32** | The padding between the bottom border of the frame and its children. This property is only applicable for auto-layout frames. | [optional] [default to 0]
**ItemSpacing** | Pointer to **float32** | The distance between children of the frame. Can be negative. This property is only applicable for auto-layout frames. | [optional] [default to 0]
**ItemReverseZIndex** | Pointer to **bool** | Determines the canvas stacking order of layers in this frame. When true, the first layer will be draw on top. This property is only applicable for auto-layout frames. | [optional] [default to false]
**StrokesIncludedInLayout** | Pointer to **bool** | Determines whether strokes are included in layout calculations. When true, auto-layout frames behave like css \&quot;box-sizing: border-box\&quot;. This property is only applicable for auto-layout frames. | [optional] [default to false]
**LayoutWrap** | Pointer to **string** | Whether this auto-layout frame has wrapping enabled. | [optional] 
**CounterAxisSpacing** | Pointer to **float32** | The distance between wrapped tracks of an auto-layout frame. This property is only applicable for auto-layout frames with &#x60;layoutWrap: \&quot;WRAP\&quot;&#x60; | [optional] 
**CounterAxisAlignContent** | Pointer to **string** | Determines how the auto-layout frame’s wrapped tracks should be aligned in the counter axis direction. This property is only applicable for auto-layout frames with &#x60;layoutWrap: \&quot;WRAP\&quot;&#x60;. | [optional] [default to "AUTO"]
**CornerRadius** | Pointer to **float32** | Radius of each corner if a single radius is set for all corners | [optional] [default to 0]
**CornerSmoothing** | Pointer to **float32** | A value that lets you control how \&quot;smooth\&quot; the corners are. Ranges from 0 to 1. 0 is the default and means that the corner is perfectly circular. A value of 0.6 means the corner matches the iOS 7 \&quot;squircle\&quot; icon shape. Other values produce various other curves. | [optional] 
**RectangleCornerRadii** | Pointer to **[]float32** | Array of length 4 of the radius of each corner of the frame, starting in the top left and proceeding clockwise.  Values are given in the order top-left, top-right, bottom-right, bottom-left. | [optional] 
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
**IndividualStrokeWeights** | Pointer to [**StrokeWeights**](StrokeWeights.md) | An object including the top, bottom, left, and right stroke weights. Only returned if individual stroke weights are used. | [optional] 
**DevStatus** | Pointer to [**DevStatusTraitDevStatus**](DevStatusTraitDevStatus.md) |  | [optional] 
**ComponentPropertyDefinitions** | Pointer to [**map[string]ComponentPropertyDefinition**](ComponentPropertyDefinition.md) | A mapping of name to &#x60;ComponentPropertyDefinition&#x60; for every component property on this component. Each property has a type, defaultValue, and other optional values. | [optional] 

## Methods

### NewComponentSetNode

`func NewComponentSetNode(id string, name string, type_ string, scrollBehavior string, blendMode BlendMode, children []SubcanvasNode, absoluteBoundingBox Rectangle, absoluteRenderBounds Rectangle, clipsContent bool, fills []Paint, effects []Effect, ) *ComponentSetNode`

NewComponentSetNode instantiates a new ComponentSetNode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComponentSetNodeWithDefaults

`func NewComponentSetNodeWithDefaults() *ComponentSetNode`

NewComponentSetNodeWithDefaults instantiates a new ComponentSetNode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ComponentSetNode) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ComponentSetNode) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ComponentSetNode) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *ComponentSetNode) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ComponentSetNode) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ComponentSetNode) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *ComponentSetNode) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ComponentSetNode) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ComponentSetNode) SetType(v string)`

SetType sets Type field to given value.


### GetVisible

`func (o *ComponentSetNode) GetVisible() bool`

GetVisible returns the Visible field if non-nil, zero value otherwise.

### GetVisibleOk

`func (o *ComponentSetNode) GetVisibleOk() (*bool, bool)`

GetVisibleOk returns a tuple with the Visible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisible

`func (o *ComponentSetNode) SetVisible(v bool)`

SetVisible sets Visible field to given value.

### HasVisible

`func (o *ComponentSetNode) HasVisible() bool`

HasVisible returns a boolean if a field has been set.

### GetLocked

`func (o *ComponentSetNode) GetLocked() bool`

GetLocked returns the Locked field if non-nil, zero value otherwise.

### GetLockedOk

`func (o *ComponentSetNode) GetLockedOk() (*bool, bool)`

GetLockedOk returns a tuple with the Locked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocked

`func (o *ComponentSetNode) SetLocked(v bool)`

SetLocked sets Locked field to given value.

### HasLocked

`func (o *ComponentSetNode) HasLocked() bool`

HasLocked returns a boolean if a field has been set.

### GetIsFixed

`func (o *ComponentSetNode) GetIsFixed() bool`

GetIsFixed returns the IsFixed field if non-nil, zero value otherwise.

### GetIsFixedOk

`func (o *ComponentSetNode) GetIsFixedOk() (*bool, bool)`

GetIsFixedOk returns a tuple with the IsFixed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFixed

`func (o *ComponentSetNode) SetIsFixed(v bool)`

SetIsFixed sets IsFixed field to given value.

### HasIsFixed

`func (o *ComponentSetNode) HasIsFixed() bool`

HasIsFixed returns a boolean if a field has been set.

### GetScrollBehavior

`func (o *ComponentSetNode) GetScrollBehavior() string`

GetScrollBehavior returns the ScrollBehavior field if non-nil, zero value otherwise.

### GetScrollBehaviorOk

`func (o *ComponentSetNode) GetScrollBehaviorOk() (*string, bool)`

GetScrollBehaviorOk returns a tuple with the ScrollBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScrollBehavior

`func (o *ComponentSetNode) SetScrollBehavior(v string)`

SetScrollBehavior sets ScrollBehavior field to given value.


### GetRotation

`func (o *ComponentSetNode) GetRotation() float32`

GetRotation returns the Rotation field if non-nil, zero value otherwise.

### GetRotationOk

`func (o *ComponentSetNode) GetRotationOk() (*float32, bool)`

GetRotationOk returns a tuple with the Rotation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRotation

`func (o *ComponentSetNode) SetRotation(v float32)`

SetRotation sets Rotation field to given value.

### HasRotation

`func (o *ComponentSetNode) HasRotation() bool`

HasRotation returns a boolean if a field has been set.

### GetComponentPropertyReferences

`func (o *ComponentSetNode) GetComponentPropertyReferences() map[string]string`

GetComponentPropertyReferences returns the ComponentPropertyReferences field if non-nil, zero value otherwise.

### GetComponentPropertyReferencesOk

`func (o *ComponentSetNode) GetComponentPropertyReferencesOk() (*map[string]string, bool)`

GetComponentPropertyReferencesOk returns a tuple with the ComponentPropertyReferences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentPropertyReferences

`func (o *ComponentSetNode) SetComponentPropertyReferences(v map[string]string)`

SetComponentPropertyReferences sets ComponentPropertyReferences field to given value.

### HasComponentPropertyReferences

`func (o *ComponentSetNode) HasComponentPropertyReferences() bool`

HasComponentPropertyReferences returns a boolean if a field has been set.

### GetPluginData

`func (o *ComponentSetNode) GetPluginData() interface{}`

GetPluginData returns the PluginData field if non-nil, zero value otherwise.

### GetPluginDataOk

`func (o *ComponentSetNode) GetPluginDataOk() (*interface{}, bool)`

GetPluginDataOk returns a tuple with the PluginData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPluginData

`func (o *ComponentSetNode) SetPluginData(v interface{})`

SetPluginData sets PluginData field to given value.

### HasPluginData

`func (o *ComponentSetNode) HasPluginData() bool`

HasPluginData returns a boolean if a field has been set.

### SetPluginDataNil

`func (o *ComponentSetNode) SetPluginDataNil(b bool)`

 SetPluginDataNil sets the value for PluginData to be an explicit nil

### UnsetPluginData
`func (o *ComponentSetNode) UnsetPluginData()`

UnsetPluginData ensures that no value is present for PluginData, not even an explicit nil
### GetSharedPluginData

`func (o *ComponentSetNode) GetSharedPluginData() interface{}`

GetSharedPluginData returns the SharedPluginData field if non-nil, zero value otherwise.

### GetSharedPluginDataOk

`func (o *ComponentSetNode) GetSharedPluginDataOk() (*interface{}, bool)`

GetSharedPluginDataOk returns a tuple with the SharedPluginData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedPluginData

`func (o *ComponentSetNode) SetSharedPluginData(v interface{})`

SetSharedPluginData sets SharedPluginData field to given value.

### HasSharedPluginData

`func (o *ComponentSetNode) HasSharedPluginData() bool`

HasSharedPluginData returns a boolean if a field has been set.

### SetSharedPluginDataNil

`func (o *ComponentSetNode) SetSharedPluginDataNil(b bool)`

 SetSharedPluginDataNil sets the value for SharedPluginData to be an explicit nil

### UnsetSharedPluginData
`func (o *ComponentSetNode) UnsetSharedPluginData()`

UnsetSharedPluginData ensures that no value is present for SharedPluginData, not even an explicit nil
### GetBoundVariables

`func (o *ComponentSetNode) GetBoundVariables() IsLayerTraitBoundVariables`

GetBoundVariables returns the BoundVariables field if non-nil, zero value otherwise.

### GetBoundVariablesOk

`func (o *ComponentSetNode) GetBoundVariablesOk() (*IsLayerTraitBoundVariables, bool)`

GetBoundVariablesOk returns a tuple with the BoundVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundVariables

`func (o *ComponentSetNode) SetBoundVariables(v IsLayerTraitBoundVariables)`

SetBoundVariables sets BoundVariables field to given value.

### HasBoundVariables

`func (o *ComponentSetNode) HasBoundVariables() bool`

HasBoundVariables returns a boolean if a field has been set.

### GetExplicitVariableModes

`func (o *ComponentSetNode) GetExplicitVariableModes() map[string]string`

GetExplicitVariableModes returns the ExplicitVariableModes field if non-nil, zero value otherwise.

### GetExplicitVariableModesOk

`func (o *ComponentSetNode) GetExplicitVariableModesOk() (*map[string]string, bool)`

GetExplicitVariableModesOk returns a tuple with the ExplicitVariableModes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExplicitVariableModes

`func (o *ComponentSetNode) SetExplicitVariableModes(v map[string]string)`

SetExplicitVariableModes sets ExplicitVariableModes field to given value.

### HasExplicitVariableModes

`func (o *ComponentSetNode) HasExplicitVariableModes() bool`

HasExplicitVariableModes returns a boolean if a field has been set.

### GetBlendMode

`func (o *ComponentSetNode) GetBlendMode() BlendMode`

GetBlendMode returns the BlendMode field if non-nil, zero value otherwise.

### GetBlendModeOk

`func (o *ComponentSetNode) GetBlendModeOk() (*BlendMode, bool)`

GetBlendModeOk returns a tuple with the BlendMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlendMode

`func (o *ComponentSetNode) SetBlendMode(v BlendMode)`

SetBlendMode sets BlendMode field to given value.


### GetOpacity

`func (o *ComponentSetNode) GetOpacity() float32`

GetOpacity returns the Opacity field if non-nil, zero value otherwise.

### GetOpacityOk

`func (o *ComponentSetNode) GetOpacityOk() (*float32, bool)`

GetOpacityOk returns a tuple with the Opacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpacity

`func (o *ComponentSetNode) SetOpacity(v float32)`

SetOpacity sets Opacity field to given value.

### HasOpacity

`func (o *ComponentSetNode) HasOpacity() bool`

HasOpacity returns a boolean if a field has been set.

### GetChildren

`func (o *ComponentSetNode) GetChildren() []SubcanvasNode`

GetChildren returns the Children field if non-nil, zero value otherwise.

### GetChildrenOk

`func (o *ComponentSetNode) GetChildrenOk() (*[]SubcanvasNode, bool)`

GetChildrenOk returns a tuple with the Children field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildren

`func (o *ComponentSetNode) SetChildren(v []SubcanvasNode)`

SetChildren sets Children field to given value.


### GetAbsoluteBoundingBox

`func (o *ComponentSetNode) GetAbsoluteBoundingBox() Rectangle`

GetAbsoluteBoundingBox returns the AbsoluteBoundingBox field if non-nil, zero value otherwise.

### GetAbsoluteBoundingBoxOk

`func (o *ComponentSetNode) GetAbsoluteBoundingBoxOk() (*Rectangle, bool)`

GetAbsoluteBoundingBoxOk returns a tuple with the AbsoluteBoundingBox field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbsoluteBoundingBox

`func (o *ComponentSetNode) SetAbsoluteBoundingBox(v Rectangle)`

SetAbsoluteBoundingBox sets AbsoluteBoundingBox field to given value.


### GetAbsoluteRenderBounds

`func (o *ComponentSetNode) GetAbsoluteRenderBounds() Rectangle`

GetAbsoluteRenderBounds returns the AbsoluteRenderBounds field if non-nil, zero value otherwise.

### GetAbsoluteRenderBoundsOk

`func (o *ComponentSetNode) GetAbsoluteRenderBoundsOk() (*Rectangle, bool)`

GetAbsoluteRenderBoundsOk returns a tuple with the AbsoluteRenderBounds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbsoluteRenderBounds

`func (o *ComponentSetNode) SetAbsoluteRenderBounds(v Rectangle)`

SetAbsoluteRenderBounds sets AbsoluteRenderBounds field to given value.


### GetPreserveRatio

`func (o *ComponentSetNode) GetPreserveRatio() bool`

GetPreserveRatio returns the PreserveRatio field if non-nil, zero value otherwise.

### GetPreserveRatioOk

`func (o *ComponentSetNode) GetPreserveRatioOk() (*bool, bool)`

GetPreserveRatioOk returns a tuple with the PreserveRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreserveRatio

`func (o *ComponentSetNode) SetPreserveRatio(v bool)`

SetPreserveRatio sets PreserveRatio field to given value.

### HasPreserveRatio

`func (o *ComponentSetNode) HasPreserveRatio() bool`

HasPreserveRatio returns a boolean if a field has been set.

### GetConstraints

`func (o *ComponentSetNode) GetConstraints() LayoutConstraint`

GetConstraints returns the Constraints field if non-nil, zero value otherwise.

### GetConstraintsOk

`func (o *ComponentSetNode) GetConstraintsOk() (*LayoutConstraint, bool)`

GetConstraintsOk returns a tuple with the Constraints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConstraints

`func (o *ComponentSetNode) SetConstraints(v LayoutConstraint)`

SetConstraints sets Constraints field to given value.

### HasConstraints

`func (o *ComponentSetNode) HasConstraints() bool`

HasConstraints returns a boolean if a field has been set.

### GetRelativeTransform

`func (o *ComponentSetNode) GetRelativeTransform() [][]float32`

GetRelativeTransform returns the RelativeTransform field if non-nil, zero value otherwise.

### GetRelativeTransformOk

`func (o *ComponentSetNode) GetRelativeTransformOk() (*[][]float32, bool)`

GetRelativeTransformOk returns a tuple with the RelativeTransform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelativeTransform

`func (o *ComponentSetNode) SetRelativeTransform(v [][]float32)`

SetRelativeTransform sets RelativeTransform field to given value.

### HasRelativeTransform

`func (o *ComponentSetNode) HasRelativeTransform() bool`

HasRelativeTransform returns a boolean if a field has been set.

### GetSize

`func (o *ComponentSetNode) GetSize() Vector`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *ComponentSetNode) GetSizeOk() (*Vector, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *ComponentSetNode) SetSize(v Vector)`

SetSize sets Size field to given value.

### HasSize

`func (o *ComponentSetNode) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetLayoutAlign

`func (o *ComponentSetNode) GetLayoutAlign() string`

GetLayoutAlign returns the LayoutAlign field if non-nil, zero value otherwise.

### GetLayoutAlignOk

`func (o *ComponentSetNode) GetLayoutAlignOk() (*string, bool)`

GetLayoutAlignOk returns a tuple with the LayoutAlign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayoutAlign

`func (o *ComponentSetNode) SetLayoutAlign(v string)`

SetLayoutAlign sets LayoutAlign field to given value.

### HasLayoutAlign

`func (o *ComponentSetNode) HasLayoutAlign() bool`

HasLayoutAlign returns a boolean if a field has been set.

### GetLayoutGrow

`func (o *ComponentSetNode) GetLayoutGrow() float32`

GetLayoutGrow returns the LayoutGrow field if non-nil, zero value otherwise.

### GetLayoutGrowOk

`func (o *ComponentSetNode) GetLayoutGrowOk() (*float32, bool)`

GetLayoutGrowOk returns a tuple with the LayoutGrow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayoutGrow

`func (o *ComponentSetNode) SetLayoutGrow(v float32)`

SetLayoutGrow sets LayoutGrow field to given value.

### HasLayoutGrow

`func (o *ComponentSetNode) HasLayoutGrow() bool`

HasLayoutGrow returns a boolean if a field has been set.

### GetLayoutPositioning

`func (o *ComponentSetNode) GetLayoutPositioning() string`

GetLayoutPositioning returns the LayoutPositioning field if non-nil, zero value otherwise.

### GetLayoutPositioningOk

`func (o *ComponentSetNode) GetLayoutPositioningOk() (*string, bool)`

GetLayoutPositioningOk returns a tuple with the LayoutPositioning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayoutPositioning

`func (o *ComponentSetNode) SetLayoutPositioning(v string)`

SetLayoutPositioning sets LayoutPositioning field to given value.

### HasLayoutPositioning

`func (o *ComponentSetNode) HasLayoutPositioning() bool`

HasLayoutPositioning returns a boolean if a field has been set.

### GetMinWidth

`func (o *ComponentSetNode) GetMinWidth() float32`

GetMinWidth returns the MinWidth field if non-nil, zero value otherwise.

### GetMinWidthOk

`func (o *ComponentSetNode) GetMinWidthOk() (*float32, bool)`

GetMinWidthOk returns a tuple with the MinWidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinWidth

`func (o *ComponentSetNode) SetMinWidth(v float32)`

SetMinWidth sets MinWidth field to given value.

### HasMinWidth

`func (o *ComponentSetNode) HasMinWidth() bool`

HasMinWidth returns a boolean if a field has been set.

### GetMaxWidth

`func (o *ComponentSetNode) GetMaxWidth() float32`

GetMaxWidth returns the MaxWidth field if non-nil, zero value otherwise.

### GetMaxWidthOk

`func (o *ComponentSetNode) GetMaxWidthOk() (*float32, bool)`

GetMaxWidthOk returns a tuple with the MaxWidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxWidth

`func (o *ComponentSetNode) SetMaxWidth(v float32)`

SetMaxWidth sets MaxWidth field to given value.

### HasMaxWidth

`func (o *ComponentSetNode) HasMaxWidth() bool`

HasMaxWidth returns a boolean if a field has been set.

### GetMinHeight

`func (o *ComponentSetNode) GetMinHeight() float32`

GetMinHeight returns the MinHeight field if non-nil, zero value otherwise.

### GetMinHeightOk

`func (o *ComponentSetNode) GetMinHeightOk() (*float32, bool)`

GetMinHeightOk returns a tuple with the MinHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinHeight

`func (o *ComponentSetNode) SetMinHeight(v float32)`

SetMinHeight sets MinHeight field to given value.

### HasMinHeight

`func (o *ComponentSetNode) HasMinHeight() bool`

HasMinHeight returns a boolean if a field has been set.

### GetMaxHeight

`func (o *ComponentSetNode) GetMaxHeight() float32`

GetMaxHeight returns the MaxHeight field if non-nil, zero value otherwise.

### GetMaxHeightOk

`func (o *ComponentSetNode) GetMaxHeightOk() (*float32, bool)`

GetMaxHeightOk returns a tuple with the MaxHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxHeight

`func (o *ComponentSetNode) SetMaxHeight(v float32)`

SetMaxHeight sets MaxHeight field to given value.

### HasMaxHeight

`func (o *ComponentSetNode) HasMaxHeight() bool`

HasMaxHeight returns a boolean if a field has been set.

### GetLayoutSizingHorizontal

`func (o *ComponentSetNode) GetLayoutSizingHorizontal() string`

GetLayoutSizingHorizontal returns the LayoutSizingHorizontal field if non-nil, zero value otherwise.

### GetLayoutSizingHorizontalOk

`func (o *ComponentSetNode) GetLayoutSizingHorizontalOk() (*string, bool)`

GetLayoutSizingHorizontalOk returns a tuple with the LayoutSizingHorizontal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayoutSizingHorizontal

`func (o *ComponentSetNode) SetLayoutSizingHorizontal(v string)`

SetLayoutSizingHorizontal sets LayoutSizingHorizontal field to given value.

### HasLayoutSizingHorizontal

`func (o *ComponentSetNode) HasLayoutSizingHorizontal() bool`

HasLayoutSizingHorizontal returns a boolean if a field has been set.

### GetLayoutSizingVertical

`func (o *ComponentSetNode) GetLayoutSizingVertical() string`

GetLayoutSizingVertical returns the LayoutSizingVertical field if non-nil, zero value otherwise.

### GetLayoutSizingVerticalOk

`func (o *ComponentSetNode) GetLayoutSizingVerticalOk() (*string, bool)`

GetLayoutSizingVerticalOk returns a tuple with the LayoutSizingVertical field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayoutSizingVertical

`func (o *ComponentSetNode) SetLayoutSizingVertical(v string)`

SetLayoutSizingVertical sets LayoutSizingVertical field to given value.

### HasLayoutSizingVertical

`func (o *ComponentSetNode) HasLayoutSizingVertical() bool`

HasLayoutSizingVertical returns a boolean if a field has been set.

### GetClipsContent

`func (o *ComponentSetNode) GetClipsContent() bool`

GetClipsContent returns the ClipsContent field if non-nil, zero value otherwise.

### GetClipsContentOk

`func (o *ComponentSetNode) GetClipsContentOk() (*bool, bool)`

GetClipsContentOk returns a tuple with the ClipsContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClipsContent

`func (o *ComponentSetNode) SetClipsContent(v bool)`

SetClipsContent sets ClipsContent field to given value.


### GetBackground

`func (o *ComponentSetNode) GetBackground() []Paint`

GetBackground returns the Background field if non-nil, zero value otherwise.

### GetBackgroundOk

`func (o *ComponentSetNode) GetBackgroundOk() (*[]Paint, bool)`

GetBackgroundOk returns a tuple with the Background field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackground

`func (o *ComponentSetNode) SetBackground(v []Paint)`

SetBackground sets Background field to given value.

### HasBackground

`func (o *ComponentSetNode) HasBackground() bool`

HasBackground returns a boolean if a field has been set.

### GetBackgroundColor

`func (o *ComponentSetNode) GetBackgroundColor() RGBA`

GetBackgroundColor returns the BackgroundColor field if non-nil, zero value otherwise.

### GetBackgroundColorOk

`func (o *ComponentSetNode) GetBackgroundColorOk() (*RGBA, bool)`

GetBackgroundColorOk returns a tuple with the BackgroundColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackgroundColor

`func (o *ComponentSetNode) SetBackgroundColor(v RGBA)`

SetBackgroundColor sets BackgroundColor field to given value.

### HasBackgroundColor

`func (o *ComponentSetNode) HasBackgroundColor() bool`

HasBackgroundColor returns a boolean if a field has been set.

### GetLayoutGrids

`func (o *ComponentSetNode) GetLayoutGrids() []LayoutGrid`

GetLayoutGrids returns the LayoutGrids field if non-nil, zero value otherwise.

### GetLayoutGridsOk

`func (o *ComponentSetNode) GetLayoutGridsOk() (*[]LayoutGrid, bool)`

GetLayoutGridsOk returns a tuple with the LayoutGrids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayoutGrids

`func (o *ComponentSetNode) SetLayoutGrids(v []LayoutGrid)`

SetLayoutGrids sets LayoutGrids field to given value.

### HasLayoutGrids

`func (o *ComponentSetNode) HasLayoutGrids() bool`

HasLayoutGrids returns a boolean if a field has been set.

### GetOverflowDirection

`func (o *ComponentSetNode) GetOverflowDirection() string`

GetOverflowDirection returns the OverflowDirection field if non-nil, zero value otherwise.

### GetOverflowDirectionOk

`func (o *ComponentSetNode) GetOverflowDirectionOk() (*string, bool)`

GetOverflowDirectionOk returns a tuple with the OverflowDirection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverflowDirection

`func (o *ComponentSetNode) SetOverflowDirection(v string)`

SetOverflowDirection sets OverflowDirection field to given value.

### HasOverflowDirection

`func (o *ComponentSetNode) HasOverflowDirection() bool`

HasOverflowDirection returns a boolean if a field has been set.

### GetLayoutMode

`func (o *ComponentSetNode) GetLayoutMode() string`

GetLayoutMode returns the LayoutMode field if non-nil, zero value otherwise.

### GetLayoutModeOk

`func (o *ComponentSetNode) GetLayoutModeOk() (*string, bool)`

GetLayoutModeOk returns a tuple with the LayoutMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayoutMode

`func (o *ComponentSetNode) SetLayoutMode(v string)`

SetLayoutMode sets LayoutMode field to given value.

### HasLayoutMode

`func (o *ComponentSetNode) HasLayoutMode() bool`

HasLayoutMode returns a boolean if a field has been set.

### GetPrimaryAxisSizingMode

`func (o *ComponentSetNode) GetPrimaryAxisSizingMode() string`

GetPrimaryAxisSizingMode returns the PrimaryAxisSizingMode field if non-nil, zero value otherwise.

### GetPrimaryAxisSizingModeOk

`func (o *ComponentSetNode) GetPrimaryAxisSizingModeOk() (*string, bool)`

GetPrimaryAxisSizingModeOk returns a tuple with the PrimaryAxisSizingMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryAxisSizingMode

`func (o *ComponentSetNode) SetPrimaryAxisSizingMode(v string)`

SetPrimaryAxisSizingMode sets PrimaryAxisSizingMode field to given value.

### HasPrimaryAxisSizingMode

`func (o *ComponentSetNode) HasPrimaryAxisSizingMode() bool`

HasPrimaryAxisSizingMode returns a boolean if a field has been set.

### GetCounterAxisSizingMode

`func (o *ComponentSetNode) GetCounterAxisSizingMode() string`

GetCounterAxisSizingMode returns the CounterAxisSizingMode field if non-nil, zero value otherwise.

### GetCounterAxisSizingModeOk

`func (o *ComponentSetNode) GetCounterAxisSizingModeOk() (*string, bool)`

GetCounterAxisSizingModeOk returns a tuple with the CounterAxisSizingMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounterAxisSizingMode

`func (o *ComponentSetNode) SetCounterAxisSizingMode(v string)`

SetCounterAxisSizingMode sets CounterAxisSizingMode field to given value.

### HasCounterAxisSizingMode

`func (o *ComponentSetNode) HasCounterAxisSizingMode() bool`

HasCounterAxisSizingMode returns a boolean if a field has been set.

### GetPrimaryAxisAlignItems

`func (o *ComponentSetNode) GetPrimaryAxisAlignItems() string`

GetPrimaryAxisAlignItems returns the PrimaryAxisAlignItems field if non-nil, zero value otherwise.

### GetPrimaryAxisAlignItemsOk

`func (o *ComponentSetNode) GetPrimaryAxisAlignItemsOk() (*string, bool)`

GetPrimaryAxisAlignItemsOk returns a tuple with the PrimaryAxisAlignItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryAxisAlignItems

`func (o *ComponentSetNode) SetPrimaryAxisAlignItems(v string)`

SetPrimaryAxisAlignItems sets PrimaryAxisAlignItems field to given value.

### HasPrimaryAxisAlignItems

`func (o *ComponentSetNode) HasPrimaryAxisAlignItems() bool`

HasPrimaryAxisAlignItems returns a boolean if a field has been set.

### GetCounterAxisAlignItems

`func (o *ComponentSetNode) GetCounterAxisAlignItems() string`

GetCounterAxisAlignItems returns the CounterAxisAlignItems field if non-nil, zero value otherwise.

### GetCounterAxisAlignItemsOk

`func (o *ComponentSetNode) GetCounterAxisAlignItemsOk() (*string, bool)`

GetCounterAxisAlignItemsOk returns a tuple with the CounterAxisAlignItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounterAxisAlignItems

`func (o *ComponentSetNode) SetCounterAxisAlignItems(v string)`

SetCounterAxisAlignItems sets CounterAxisAlignItems field to given value.

### HasCounterAxisAlignItems

`func (o *ComponentSetNode) HasCounterAxisAlignItems() bool`

HasCounterAxisAlignItems returns a boolean if a field has been set.

### GetPaddingLeft

`func (o *ComponentSetNode) GetPaddingLeft() float32`

GetPaddingLeft returns the PaddingLeft field if non-nil, zero value otherwise.

### GetPaddingLeftOk

`func (o *ComponentSetNode) GetPaddingLeftOk() (*float32, bool)`

GetPaddingLeftOk returns a tuple with the PaddingLeft field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaddingLeft

`func (o *ComponentSetNode) SetPaddingLeft(v float32)`

SetPaddingLeft sets PaddingLeft field to given value.

### HasPaddingLeft

`func (o *ComponentSetNode) HasPaddingLeft() bool`

HasPaddingLeft returns a boolean if a field has been set.

### GetPaddingRight

`func (o *ComponentSetNode) GetPaddingRight() float32`

GetPaddingRight returns the PaddingRight field if non-nil, zero value otherwise.

### GetPaddingRightOk

`func (o *ComponentSetNode) GetPaddingRightOk() (*float32, bool)`

GetPaddingRightOk returns a tuple with the PaddingRight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaddingRight

`func (o *ComponentSetNode) SetPaddingRight(v float32)`

SetPaddingRight sets PaddingRight field to given value.

### HasPaddingRight

`func (o *ComponentSetNode) HasPaddingRight() bool`

HasPaddingRight returns a boolean if a field has been set.

### GetPaddingTop

`func (o *ComponentSetNode) GetPaddingTop() float32`

GetPaddingTop returns the PaddingTop field if non-nil, zero value otherwise.

### GetPaddingTopOk

`func (o *ComponentSetNode) GetPaddingTopOk() (*float32, bool)`

GetPaddingTopOk returns a tuple with the PaddingTop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaddingTop

`func (o *ComponentSetNode) SetPaddingTop(v float32)`

SetPaddingTop sets PaddingTop field to given value.

### HasPaddingTop

`func (o *ComponentSetNode) HasPaddingTop() bool`

HasPaddingTop returns a boolean if a field has been set.

### GetPaddingBottom

`func (o *ComponentSetNode) GetPaddingBottom() float32`

GetPaddingBottom returns the PaddingBottom field if non-nil, zero value otherwise.

### GetPaddingBottomOk

`func (o *ComponentSetNode) GetPaddingBottomOk() (*float32, bool)`

GetPaddingBottomOk returns a tuple with the PaddingBottom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaddingBottom

`func (o *ComponentSetNode) SetPaddingBottom(v float32)`

SetPaddingBottom sets PaddingBottom field to given value.

### HasPaddingBottom

`func (o *ComponentSetNode) HasPaddingBottom() bool`

HasPaddingBottom returns a boolean if a field has been set.

### GetItemSpacing

`func (o *ComponentSetNode) GetItemSpacing() float32`

GetItemSpacing returns the ItemSpacing field if non-nil, zero value otherwise.

### GetItemSpacingOk

`func (o *ComponentSetNode) GetItemSpacingOk() (*float32, bool)`

GetItemSpacingOk returns a tuple with the ItemSpacing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemSpacing

`func (o *ComponentSetNode) SetItemSpacing(v float32)`

SetItemSpacing sets ItemSpacing field to given value.

### HasItemSpacing

`func (o *ComponentSetNode) HasItemSpacing() bool`

HasItemSpacing returns a boolean if a field has been set.

### GetItemReverseZIndex

`func (o *ComponentSetNode) GetItemReverseZIndex() bool`

GetItemReverseZIndex returns the ItemReverseZIndex field if non-nil, zero value otherwise.

### GetItemReverseZIndexOk

`func (o *ComponentSetNode) GetItemReverseZIndexOk() (*bool, bool)`

GetItemReverseZIndexOk returns a tuple with the ItemReverseZIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemReverseZIndex

`func (o *ComponentSetNode) SetItemReverseZIndex(v bool)`

SetItemReverseZIndex sets ItemReverseZIndex field to given value.

### HasItemReverseZIndex

`func (o *ComponentSetNode) HasItemReverseZIndex() bool`

HasItemReverseZIndex returns a boolean if a field has been set.

### GetStrokesIncludedInLayout

`func (o *ComponentSetNode) GetStrokesIncludedInLayout() bool`

GetStrokesIncludedInLayout returns the StrokesIncludedInLayout field if non-nil, zero value otherwise.

### GetStrokesIncludedInLayoutOk

`func (o *ComponentSetNode) GetStrokesIncludedInLayoutOk() (*bool, bool)`

GetStrokesIncludedInLayoutOk returns a tuple with the StrokesIncludedInLayout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrokesIncludedInLayout

`func (o *ComponentSetNode) SetStrokesIncludedInLayout(v bool)`

SetStrokesIncludedInLayout sets StrokesIncludedInLayout field to given value.

### HasStrokesIncludedInLayout

`func (o *ComponentSetNode) HasStrokesIncludedInLayout() bool`

HasStrokesIncludedInLayout returns a boolean if a field has been set.

### GetLayoutWrap

`func (o *ComponentSetNode) GetLayoutWrap() string`

GetLayoutWrap returns the LayoutWrap field if non-nil, zero value otherwise.

### GetLayoutWrapOk

`func (o *ComponentSetNode) GetLayoutWrapOk() (*string, bool)`

GetLayoutWrapOk returns a tuple with the LayoutWrap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayoutWrap

`func (o *ComponentSetNode) SetLayoutWrap(v string)`

SetLayoutWrap sets LayoutWrap field to given value.

### HasLayoutWrap

`func (o *ComponentSetNode) HasLayoutWrap() bool`

HasLayoutWrap returns a boolean if a field has been set.

### GetCounterAxisSpacing

`func (o *ComponentSetNode) GetCounterAxisSpacing() float32`

GetCounterAxisSpacing returns the CounterAxisSpacing field if non-nil, zero value otherwise.

### GetCounterAxisSpacingOk

`func (o *ComponentSetNode) GetCounterAxisSpacingOk() (*float32, bool)`

GetCounterAxisSpacingOk returns a tuple with the CounterAxisSpacing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounterAxisSpacing

`func (o *ComponentSetNode) SetCounterAxisSpacing(v float32)`

SetCounterAxisSpacing sets CounterAxisSpacing field to given value.

### HasCounterAxisSpacing

`func (o *ComponentSetNode) HasCounterAxisSpacing() bool`

HasCounterAxisSpacing returns a boolean if a field has been set.

### GetCounterAxisAlignContent

`func (o *ComponentSetNode) GetCounterAxisAlignContent() string`

GetCounterAxisAlignContent returns the CounterAxisAlignContent field if non-nil, zero value otherwise.

### GetCounterAxisAlignContentOk

`func (o *ComponentSetNode) GetCounterAxisAlignContentOk() (*string, bool)`

GetCounterAxisAlignContentOk returns a tuple with the CounterAxisAlignContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounterAxisAlignContent

`func (o *ComponentSetNode) SetCounterAxisAlignContent(v string)`

SetCounterAxisAlignContent sets CounterAxisAlignContent field to given value.

### HasCounterAxisAlignContent

`func (o *ComponentSetNode) HasCounterAxisAlignContent() bool`

HasCounterAxisAlignContent returns a boolean if a field has been set.

### GetCornerRadius

`func (o *ComponentSetNode) GetCornerRadius() float32`

GetCornerRadius returns the CornerRadius field if non-nil, zero value otherwise.

### GetCornerRadiusOk

`func (o *ComponentSetNode) GetCornerRadiusOk() (*float32, bool)`

GetCornerRadiusOk returns a tuple with the CornerRadius field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCornerRadius

`func (o *ComponentSetNode) SetCornerRadius(v float32)`

SetCornerRadius sets CornerRadius field to given value.

### HasCornerRadius

`func (o *ComponentSetNode) HasCornerRadius() bool`

HasCornerRadius returns a boolean if a field has been set.

### GetCornerSmoothing

`func (o *ComponentSetNode) GetCornerSmoothing() float32`

GetCornerSmoothing returns the CornerSmoothing field if non-nil, zero value otherwise.

### GetCornerSmoothingOk

`func (o *ComponentSetNode) GetCornerSmoothingOk() (*float32, bool)`

GetCornerSmoothingOk returns a tuple with the CornerSmoothing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCornerSmoothing

`func (o *ComponentSetNode) SetCornerSmoothing(v float32)`

SetCornerSmoothing sets CornerSmoothing field to given value.

### HasCornerSmoothing

`func (o *ComponentSetNode) HasCornerSmoothing() bool`

HasCornerSmoothing returns a boolean if a field has been set.

### GetRectangleCornerRadii

`func (o *ComponentSetNode) GetRectangleCornerRadii() []float32`

GetRectangleCornerRadii returns the RectangleCornerRadii field if non-nil, zero value otherwise.

### GetRectangleCornerRadiiOk

`func (o *ComponentSetNode) GetRectangleCornerRadiiOk() (*[]float32, bool)`

GetRectangleCornerRadiiOk returns a tuple with the RectangleCornerRadii field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRectangleCornerRadii

`func (o *ComponentSetNode) SetRectangleCornerRadii(v []float32)`

SetRectangleCornerRadii sets RectangleCornerRadii field to given value.

### HasRectangleCornerRadii

`func (o *ComponentSetNode) HasRectangleCornerRadii() bool`

HasRectangleCornerRadii returns a boolean if a field has been set.

### GetFills

`func (o *ComponentSetNode) GetFills() []Paint`

GetFills returns the Fills field if non-nil, zero value otherwise.

### GetFillsOk

`func (o *ComponentSetNode) GetFillsOk() (*[]Paint, bool)`

GetFillsOk returns a tuple with the Fills field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFills

`func (o *ComponentSetNode) SetFills(v []Paint)`

SetFills sets Fills field to given value.


### GetStyles

`func (o *ComponentSetNode) GetStyles() map[string]string`

GetStyles returns the Styles field if non-nil, zero value otherwise.

### GetStylesOk

`func (o *ComponentSetNode) GetStylesOk() (*map[string]string, bool)`

GetStylesOk returns a tuple with the Styles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStyles

`func (o *ComponentSetNode) SetStyles(v map[string]string)`

SetStyles sets Styles field to given value.

### HasStyles

`func (o *ComponentSetNode) HasStyles() bool`

HasStyles returns a boolean if a field has been set.

### GetStrokes

`func (o *ComponentSetNode) GetStrokes() []Paint`

GetStrokes returns the Strokes field if non-nil, zero value otherwise.

### GetStrokesOk

`func (o *ComponentSetNode) GetStrokesOk() (*[]Paint, bool)`

GetStrokesOk returns a tuple with the Strokes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrokes

`func (o *ComponentSetNode) SetStrokes(v []Paint)`

SetStrokes sets Strokes field to given value.

### HasStrokes

`func (o *ComponentSetNode) HasStrokes() bool`

HasStrokes returns a boolean if a field has been set.

### GetStrokeWeight

`func (o *ComponentSetNode) GetStrokeWeight() float32`

GetStrokeWeight returns the StrokeWeight field if non-nil, zero value otherwise.

### GetStrokeWeightOk

`func (o *ComponentSetNode) GetStrokeWeightOk() (*float32, bool)`

GetStrokeWeightOk returns a tuple with the StrokeWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrokeWeight

`func (o *ComponentSetNode) SetStrokeWeight(v float32)`

SetStrokeWeight sets StrokeWeight field to given value.

### HasStrokeWeight

`func (o *ComponentSetNode) HasStrokeWeight() bool`

HasStrokeWeight returns a boolean if a field has been set.

### GetStrokeAlign

`func (o *ComponentSetNode) GetStrokeAlign() string`

GetStrokeAlign returns the StrokeAlign field if non-nil, zero value otherwise.

### GetStrokeAlignOk

`func (o *ComponentSetNode) GetStrokeAlignOk() (*string, bool)`

GetStrokeAlignOk returns a tuple with the StrokeAlign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrokeAlign

`func (o *ComponentSetNode) SetStrokeAlign(v string)`

SetStrokeAlign sets StrokeAlign field to given value.

### HasStrokeAlign

`func (o *ComponentSetNode) HasStrokeAlign() bool`

HasStrokeAlign returns a boolean if a field has been set.

### GetStrokeJoin

`func (o *ComponentSetNode) GetStrokeJoin() string`

GetStrokeJoin returns the StrokeJoin field if non-nil, zero value otherwise.

### GetStrokeJoinOk

`func (o *ComponentSetNode) GetStrokeJoinOk() (*string, bool)`

GetStrokeJoinOk returns a tuple with the StrokeJoin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrokeJoin

`func (o *ComponentSetNode) SetStrokeJoin(v string)`

SetStrokeJoin sets StrokeJoin field to given value.

### HasStrokeJoin

`func (o *ComponentSetNode) HasStrokeJoin() bool`

HasStrokeJoin returns a boolean if a field has been set.

### GetStrokeDashes

`func (o *ComponentSetNode) GetStrokeDashes() []float32`

GetStrokeDashes returns the StrokeDashes field if non-nil, zero value otherwise.

### GetStrokeDashesOk

`func (o *ComponentSetNode) GetStrokeDashesOk() (*[]float32, bool)`

GetStrokeDashesOk returns a tuple with the StrokeDashes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrokeDashes

`func (o *ComponentSetNode) SetStrokeDashes(v []float32)`

SetStrokeDashes sets StrokeDashes field to given value.

### HasStrokeDashes

`func (o *ComponentSetNode) HasStrokeDashes() bool`

HasStrokeDashes returns a boolean if a field has been set.

### GetFillOverrideTable

`func (o *ComponentSetNode) GetFillOverrideTable() map[string]HasGeometryTraitAllOfFillOverrideTable`

GetFillOverrideTable returns the FillOverrideTable field if non-nil, zero value otherwise.

### GetFillOverrideTableOk

`func (o *ComponentSetNode) GetFillOverrideTableOk() (*map[string]HasGeometryTraitAllOfFillOverrideTable, bool)`

GetFillOverrideTableOk returns a tuple with the FillOverrideTable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFillOverrideTable

`func (o *ComponentSetNode) SetFillOverrideTable(v map[string]HasGeometryTraitAllOfFillOverrideTable)`

SetFillOverrideTable sets FillOverrideTable field to given value.

### HasFillOverrideTable

`func (o *ComponentSetNode) HasFillOverrideTable() bool`

HasFillOverrideTable returns a boolean if a field has been set.

### GetFillGeometry

`func (o *ComponentSetNode) GetFillGeometry() []Path`

GetFillGeometry returns the FillGeometry field if non-nil, zero value otherwise.

### GetFillGeometryOk

`func (o *ComponentSetNode) GetFillGeometryOk() (*[]Path, bool)`

GetFillGeometryOk returns a tuple with the FillGeometry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFillGeometry

`func (o *ComponentSetNode) SetFillGeometry(v []Path)`

SetFillGeometry sets FillGeometry field to given value.

### HasFillGeometry

`func (o *ComponentSetNode) HasFillGeometry() bool`

HasFillGeometry returns a boolean if a field has been set.

### GetStrokeGeometry

`func (o *ComponentSetNode) GetStrokeGeometry() []Path`

GetStrokeGeometry returns the StrokeGeometry field if non-nil, zero value otherwise.

### GetStrokeGeometryOk

`func (o *ComponentSetNode) GetStrokeGeometryOk() (*[]Path, bool)`

GetStrokeGeometryOk returns a tuple with the StrokeGeometry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrokeGeometry

`func (o *ComponentSetNode) SetStrokeGeometry(v []Path)`

SetStrokeGeometry sets StrokeGeometry field to given value.

### HasStrokeGeometry

`func (o *ComponentSetNode) HasStrokeGeometry() bool`

HasStrokeGeometry returns a boolean if a field has been set.

### GetStrokeCap

`func (o *ComponentSetNode) GetStrokeCap() string`

GetStrokeCap returns the StrokeCap field if non-nil, zero value otherwise.

### GetStrokeCapOk

`func (o *ComponentSetNode) GetStrokeCapOk() (*string, bool)`

GetStrokeCapOk returns a tuple with the StrokeCap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrokeCap

`func (o *ComponentSetNode) SetStrokeCap(v string)`

SetStrokeCap sets StrokeCap field to given value.

### HasStrokeCap

`func (o *ComponentSetNode) HasStrokeCap() bool`

HasStrokeCap returns a boolean if a field has been set.

### GetStrokeMiterAngle

`func (o *ComponentSetNode) GetStrokeMiterAngle() float32`

GetStrokeMiterAngle returns the StrokeMiterAngle field if non-nil, zero value otherwise.

### GetStrokeMiterAngleOk

`func (o *ComponentSetNode) GetStrokeMiterAngleOk() (*float32, bool)`

GetStrokeMiterAngleOk returns a tuple with the StrokeMiterAngle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrokeMiterAngle

`func (o *ComponentSetNode) SetStrokeMiterAngle(v float32)`

SetStrokeMiterAngle sets StrokeMiterAngle field to given value.

### HasStrokeMiterAngle

`func (o *ComponentSetNode) HasStrokeMiterAngle() bool`

HasStrokeMiterAngle returns a boolean if a field has been set.

### GetExportSettings

`func (o *ComponentSetNode) GetExportSettings() []ExportSetting`

GetExportSettings returns the ExportSettings field if non-nil, zero value otherwise.

### GetExportSettingsOk

`func (o *ComponentSetNode) GetExportSettingsOk() (*[]ExportSetting, bool)`

GetExportSettingsOk returns a tuple with the ExportSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportSettings

`func (o *ComponentSetNode) SetExportSettings(v []ExportSetting)`

SetExportSettings sets ExportSettings field to given value.

### HasExportSettings

`func (o *ComponentSetNode) HasExportSettings() bool`

HasExportSettings returns a boolean if a field has been set.

### GetEffects

`func (o *ComponentSetNode) GetEffects() []Effect`

GetEffects returns the Effects field if non-nil, zero value otherwise.

### GetEffectsOk

`func (o *ComponentSetNode) GetEffectsOk() (*[]Effect, bool)`

GetEffectsOk returns a tuple with the Effects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffects

`func (o *ComponentSetNode) SetEffects(v []Effect)`

SetEffects sets Effects field to given value.


### GetIsMask

`func (o *ComponentSetNode) GetIsMask() bool`

GetIsMask returns the IsMask field if non-nil, zero value otherwise.

### GetIsMaskOk

`func (o *ComponentSetNode) GetIsMaskOk() (*bool, bool)`

GetIsMaskOk returns a tuple with the IsMask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMask

`func (o *ComponentSetNode) SetIsMask(v bool)`

SetIsMask sets IsMask field to given value.

### HasIsMask

`func (o *ComponentSetNode) HasIsMask() bool`

HasIsMask returns a boolean if a field has been set.

### GetMaskType

`func (o *ComponentSetNode) GetMaskType() string`

GetMaskType returns the MaskType field if non-nil, zero value otherwise.

### GetMaskTypeOk

`func (o *ComponentSetNode) GetMaskTypeOk() (*string, bool)`

GetMaskTypeOk returns a tuple with the MaskType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaskType

`func (o *ComponentSetNode) SetMaskType(v string)`

SetMaskType sets MaskType field to given value.

### HasMaskType

`func (o *ComponentSetNode) HasMaskType() bool`

HasMaskType returns a boolean if a field has been set.

### GetIsMaskOutline

`func (o *ComponentSetNode) GetIsMaskOutline() bool`

GetIsMaskOutline returns the IsMaskOutline field if non-nil, zero value otherwise.

### GetIsMaskOutlineOk

`func (o *ComponentSetNode) GetIsMaskOutlineOk() (*bool, bool)`

GetIsMaskOutlineOk returns a tuple with the IsMaskOutline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMaskOutline

`func (o *ComponentSetNode) SetIsMaskOutline(v bool)`

SetIsMaskOutline sets IsMaskOutline field to given value.

### HasIsMaskOutline

`func (o *ComponentSetNode) HasIsMaskOutline() bool`

HasIsMaskOutline returns a boolean if a field has been set.

### GetTransitionNodeID

`func (o *ComponentSetNode) GetTransitionNodeID() string`

GetTransitionNodeID returns the TransitionNodeID field if non-nil, zero value otherwise.

### GetTransitionNodeIDOk

`func (o *ComponentSetNode) GetTransitionNodeIDOk() (*string, bool)`

GetTransitionNodeIDOk returns a tuple with the TransitionNodeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransitionNodeID

`func (o *ComponentSetNode) SetTransitionNodeID(v string)`

SetTransitionNodeID sets TransitionNodeID field to given value.

### HasTransitionNodeID

`func (o *ComponentSetNode) HasTransitionNodeID() bool`

HasTransitionNodeID returns a boolean if a field has been set.

### GetTransitionDuration

`func (o *ComponentSetNode) GetTransitionDuration() float32`

GetTransitionDuration returns the TransitionDuration field if non-nil, zero value otherwise.

### GetTransitionDurationOk

`func (o *ComponentSetNode) GetTransitionDurationOk() (*float32, bool)`

GetTransitionDurationOk returns a tuple with the TransitionDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransitionDuration

`func (o *ComponentSetNode) SetTransitionDuration(v float32)`

SetTransitionDuration sets TransitionDuration field to given value.

### HasTransitionDuration

`func (o *ComponentSetNode) HasTransitionDuration() bool`

HasTransitionDuration returns a boolean if a field has been set.

### GetTransitionEasing

`func (o *ComponentSetNode) GetTransitionEasing() EasingType`

GetTransitionEasing returns the TransitionEasing field if non-nil, zero value otherwise.

### GetTransitionEasingOk

`func (o *ComponentSetNode) GetTransitionEasingOk() (*EasingType, bool)`

GetTransitionEasingOk returns a tuple with the TransitionEasing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransitionEasing

`func (o *ComponentSetNode) SetTransitionEasing(v EasingType)`

SetTransitionEasing sets TransitionEasing field to given value.

### HasTransitionEasing

`func (o *ComponentSetNode) HasTransitionEasing() bool`

HasTransitionEasing returns a boolean if a field has been set.

### GetInteractions

`func (o *ComponentSetNode) GetInteractions() []Interaction`

GetInteractions returns the Interactions field if non-nil, zero value otherwise.

### GetInteractionsOk

`func (o *ComponentSetNode) GetInteractionsOk() (*[]Interaction, bool)`

GetInteractionsOk returns a tuple with the Interactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInteractions

`func (o *ComponentSetNode) SetInteractions(v []Interaction)`

SetInteractions sets Interactions field to given value.

### HasInteractions

`func (o *ComponentSetNode) HasInteractions() bool`

HasInteractions returns a boolean if a field has been set.

### GetIndividualStrokeWeights

`func (o *ComponentSetNode) GetIndividualStrokeWeights() StrokeWeights`

GetIndividualStrokeWeights returns the IndividualStrokeWeights field if non-nil, zero value otherwise.

### GetIndividualStrokeWeightsOk

`func (o *ComponentSetNode) GetIndividualStrokeWeightsOk() (*StrokeWeights, bool)`

GetIndividualStrokeWeightsOk returns a tuple with the IndividualStrokeWeights field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndividualStrokeWeights

`func (o *ComponentSetNode) SetIndividualStrokeWeights(v StrokeWeights)`

SetIndividualStrokeWeights sets IndividualStrokeWeights field to given value.

### HasIndividualStrokeWeights

`func (o *ComponentSetNode) HasIndividualStrokeWeights() bool`

HasIndividualStrokeWeights returns a boolean if a field has been set.

### GetDevStatus

`func (o *ComponentSetNode) GetDevStatus() DevStatusTraitDevStatus`

GetDevStatus returns the DevStatus field if non-nil, zero value otherwise.

### GetDevStatusOk

`func (o *ComponentSetNode) GetDevStatusOk() (*DevStatusTraitDevStatus, bool)`

GetDevStatusOk returns a tuple with the DevStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevStatus

`func (o *ComponentSetNode) SetDevStatus(v DevStatusTraitDevStatus)`

SetDevStatus sets DevStatus field to given value.

### HasDevStatus

`func (o *ComponentSetNode) HasDevStatus() bool`

HasDevStatus returns a boolean if a field has been set.

### GetComponentPropertyDefinitions

`func (o *ComponentSetNode) GetComponentPropertyDefinitions() map[string]ComponentPropertyDefinition`

GetComponentPropertyDefinitions returns the ComponentPropertyDefinitions field if non-nil, zero value otherwise.

### GetComponentPropertyDefinitionsOk

`func (o *ComponentSetNode) GetComponentPropertyDefinitionsOk() (*map[string]ComponentPropertyDefinition, bool)`

GetComponentPropertyDefinitionsOk returns a tuple with the ComponentPropertyDefinitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentPropertyDefinitions

`func (o *ComponentSetNode) SetComponentPropertyDefinitions(v map[string]ComponentPropertyDefinition)`

SetComponentPropertyDefinitions sets ComponentPropertyDefinitions field to given value.

### HasComponentPropertyDefinitions

`func (o *ComponentSetNode) HasComponentPropertyDefinitions() bool`

HasComponentPropertyDefinitions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

