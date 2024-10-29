# InstanceNode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | A string uniquely identifying this node within the document. | 
**Name** | **string** | The name given to the node by the user in the tool. | 
**Type** | **string** | The type of this node, represented by the string literal \&quot;INSTANCE\&quot; | 
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
**ComponentId** | **string** | ID of component that this instance came from. | 
**IsExposedInstance** | Pointer to **bool** | If true, this node has been marked as exposed to its containing component or component set. | [optional] [default to false]
**ExposedInstances** | Pointer to **[]string** | IDs of instances that have been exposed to this node&#39;s level. | [optional] 
**ComponentProperties** | Pointer to [**map[string]ComponentProperty**](ComponentProperty.md) | A mapping of name to &#x60;ComponentProperty&#x60; for all component properties on this instance. Each property has a type, value, and other optional values. | [optional] 
**Overrides** | [**[]Overrides**](Overrides.md) | An array of all of the fields directly overridden on this instance. Inherited overrides are not included. | 

## Methods

### NewInstanceNode

`func NewInstanceNode(id string, name string, type_ string, scrollBehavior string, blendMode BlendMode, children []SubcanvasNode, absoluteBoundingBox Rectangle, absoluteRenderBounds Rectangle, clipsContent bool, fills []Paint, effects []Effect, componentId string, overrides []Overrides, ) *InstanceNode`

NewInstanceNode instantiates a new InstanceNode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstanceNodeWithDefaults

`func NewInstanceNodeWithDefaults() *InstanceNode`

NewInstanceNodeWithDefaults instantiates a new InstanceNode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InstanceNode) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InstanceNode) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InstanceNode) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *InstanceNode) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InstanceNode) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InstanceNode) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *InstanceNode) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InstanceNode) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InstanceNode) SetType(v string)`

SetType sets Type field to given value.


### GetVisible

`func (o *InstanceNode) GetVisible() bool`

GetVisible returns the Visible field if non-nil, zero value otherwise.

### GetVisibleOk

`func (o *InstanceNode) GetVisibleOk() (*bool, bool)`

GetVisibleOk returns a tuple with the Visible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisible

`func (o *InstanceNode) SetVisible(v bool)`

SetVisible sets Visible field to given value.

### HasVisible

`func (o *InstanceNode) HasVisible() bool`

HasVisible returns a boolean if a field has been set.

### GetLocked

`func (o *InstanceNode) GetLocked() bool`

GetLocked returns the Locked field if non-nil, zero value otherwise.

### GetLockedOk

`func (o *InstanceNode) GetLockedOk() (*bool, bool)`

GetLockedOk returns a tuple with the Locked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocked

`func (o *InstanceNode) SetLocked(v bool)`

SetLocked sets Locked field to given value.

### HasLocked

`func (o *InstanceNode) HasLocked() bool`

HasLocked returns a boolean if a field has been set.

### GetIsFixed

`func (o *InstanceNode) GetIsFixed() bool`

GetIsFixed returns the IsFixed field if non-nil, zero value otherwise.

### GetIsFixedOk

`func (o *InstanceNode) GetIsFixedOk() (*bool, bool)`

GetIsFixedOk returns a tuple with the IsFixed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFixed

`func (o *InstanceNode) SetIsFixed(v bool)`

SetIsFixed sets IsFixed field to given value.

### HasIsFixed

`func (o *InstanceNode) HasIsFixed() bool`

HasIsFixed returns a boolean if a field has been set.

### GetScrollBehavior

`func (o *InstanceNode) GetScrollBehavior() string`

GetScrollBehavior returns the ScrollBehavior field if non-nil, zero value otherwise.

### GetScrollBehaviorOk

`func (o *InstanceNode) GetScrollBehaviorOk() (*string, bool)`

GetScrollBehaviorOk returns a tuple with the ScrollBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScrollBehavior

`func (o *InstanceNode) SetScrollBehavior(v string)`

SetScrollBehavior sets ScrollBehavior field to given value.


### GetRotation

`func (o *InstanceNode) GetRotation() float32`

GetRotation returns the Rotation field if non-nil, zero value otherwise.

### GetRotationOk

`func (o *InstanceNode) GetRotationOk() (*float32, bool)`

GetRotationOk returns a tuple with the Rotation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRotation

`func (o *InstanceNode) SetRotation(v float32)`

SetRotation sets Rotation field to given value.

### HasRotation

`func (o *InstanceNode) HasRotation() bool`

HasRotation returns a boolean if a field has been set.

### GetComponentPropertyReferences

`func (o *InstanceNode) GetComponentPropertyReferences() map[string]string`

GetComponentPropertyReferences returns the ComponentPropertyReferences field if non-nil, zero value otherwise.

### GetComponentPropertyReferencesOk

`func (o *InstanceNode) GetComponentPropertyReferencesOk() (*map[string]string, bool)`

GetComponentPropertyReferencesOk returns a tuple with the ComponentPropertyReferences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentPropertyReferences

`func (o *InstanceNode) SetComponentPropertyReferences(v map[string]string)`

SetComponentPropertyReferences sets ComponentPropertyReferences field to given value.

### HasComponentPropertyReferences

`func (o *InstanceNode) HasComponentPropertyReferences() bool`

HasComponentPropertyReferences returns a boolean if a field has been set.

### GetPluginData

`func (o *InstanceNode) GetPluginData() interface{}`

GetPluginData returns the PluginData field if non-nil, zero value otherwise.

### GetPluginDataOk

`func (o *InstanceNode) GetPluginDataOk() (*interface{}, bool)`

GetPluginDataOk returns a tuple with the PluginData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPluginData

`func (o *InstanceNode) SetPluginData(v interface{})`

SetPluginData sets PluginData field to given value.

### HasPluginData

`func (o *InstanceNode) HasPluginData() bool`

HasPluginData returns a boolean if a field has been set.

### SetPluginDataNil

`func (o *InstanceNode) SetPluginDataNil(b bool)`

 SetPluginDataNil sets the value for PluginData to be an explicit nil

### UnsetPluginData
`func (o *InstanceNode) UnsetPluginData()`

UnsetPluginData ensures that no value is present for PluginData, not even an explicit nil
### GetSharedPluginData

`func (o *InstanceNode) GetSharedPluginData() interface{}`

GetSharedPluginData returns the SharedPluginData field if non-nil, zero value otherwise.

### GetSharedPluginDataOk

`func (o *InstanceNode) GetSharedPluginDataOk() (*interface{}, bool)`

GetSharedPluginDataOk returns a tuple with the SharedPluginData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedPluginData

`func (o *InstanceNode) SetSharedPluginData(v interface{})`

SetSharedPluginData sets SharedPluginData field to given value.

### HasSharedPluginData

`func (o *InstanceNode) HasSharedPluginData() bool`

HasSharedPluginData returns a boolean if a field has been set.

### SetSharedPluginDataNil

`func (o *InstanceNode) SetSharedPluginDataNil(b bool)`

 SetSharedPluginDataNil sets the value for SharedPluginData to be an explicit nil

### UnsetSharedPluginData
`func (o *InstanceNode) UnsetSharedPluginData()`

UnsetSharedPluginData ensures that no value is present for SharedPluginData, not even an explicit nil
### GetBoundVariables

`func (o *InstanceNode) GetBoundVariables() IsLayerTraitBoundVariables`

GetBoundVariables returns the BoundVariables field if non-nil, zero value otherwise.

### GetBoundVariablesOk

`func (o *InstanceNode) GetBoundVariablesOk() (*IsLayerTraitBoundVariables, bool)`

GetBoundVariablesOk returns a tuple with the BoundVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundVariables

`func (o *InstanceNode) SetBoundVariables(v IsLayerTraitBoundVariables)`

SetBoundVariables sets BoundVariables field to given value.

### HasBoundVariables

`func (o *InstanceNode) HasBoundVariables() bool`

HasBoundVariables returns a boolean if a field has been set.

### GetExplicitVariableModes

`func (o *InstanceNode) GetExplicitVariableModes() map[string]string`

GetExplicitVariableModes returns the ExplicitVariableModes field if non-nil, zero value otherwise.

### GetExplicitVariableModesOk

`func (o *InstanceNode) GetExplicitVariableModesOk() (*map[string]string, bool)`

GetExplicitVariableModesOk returns a tuple with the ExplicitVariableModes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExplicitVariableModes

`func (o *InstanceNode) SetExplicitVariableModes(v map[string]string)`

SetExplicitVariableModes sets ExplicitVariableModes field to given value.

### HasExplicitVariableModes

`func (o *InstanceNode) HasExplicitVariableModes() bool`

HasExplicitVariableModes returns a boolean if a field has been set.

### GetBlendMode

`func (o *InstanceNode) GetBlendMode() BlendMode`

GetBlendMode returns the BlendMode field if non-nil, zero value otherwise.

### GetBlendModeOk

`func (o *InstanceNode) GetBlendModeOk() (*BlendMode, bool)`

GetBlendModeOk returns a tuple with the BlendMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlendMode

`func (o *InstanceNode) SetBlendMode(v BlendMode)`

SetBlendMode sets BlendMode field to given value.


### GetOpacity

`func (o *InstanceNode) GetOpacity() float32`

GetOpacity returns the Opacity field if non-nil, zero value otherwise.

### GetOpacityOk

`func (o *InstanceNode) GetOpacityOk() (*float32, bool)`

GetOpacityOk returns a tuple with the Opacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpacity

`func (o *InstanceNode) SetOpacity(v float32)`

SetOpacity sets Opacity field to given value.

### HasOpacity

`func (o *InstanceNode) HasOpacity() bool`

HasOpacity returns a boolean if a field has been set.

### GetChildren

`func (o *InstanceNode) GetChildren() []SubcanvasNode`

GetChildren returns the Children field if non-nil, zero value otherwise.

### GetChildrenOk

`func (o *InstanceNode) GetChildrenOk() (*[]SubcanvasNode, bool)`

GetChildrenOk returns a tuple with the Children field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildren

`func (o *InstanceNode) SetChildren(v []SubcanvasNode)`

SetChildren sets Children field to given value.


### GetAbsoluteBoundingBox

`func (o *InstanceNode) GetAbsoluteBoundingBox() Rectangle`

GetAbsoluteBoundingBox returns the AbsoluteBoundingBox field if non-nil, zero value otherwise.

### GetAbsoluteBoundingBoxOk

`func (o *InstanceNode) GetAbsoluteBoundingBoxOk() (*Rectangle, bool)`

GetAbsoluteBoundingBoxOk returns a tuple with the AbsoluteBoundingBox field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbsoluteBoundingBox

`func (o *InstanceNode) SetAbsoluteBoundingBox(v Rectangle)`

SetAbsoluteBoundingBox sets AbsoluteBoundingBox field to given value.


### GetAbsoluteRenderBounds

`func (o *InstanceNode) GetAbsoluteRenderBounds() Rectangle`

GetAbsoluteRenderBounds returns the AbsoluteRenderBounds field if non-nil, zero value otherwise.

### GetAbsoluteRenderBoundsOk

`func (o *InstanceNode) GetAbsoluteRenderBoundsOk() (*Rectangle, bool)`

GetAbsoluteRenderBoundsOk returns a tuple with the AbsoluteRenderBounds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbsoluteRenderBounds

`func (o *InstanceNode) SetAbsoluteRenderBounds(v Rectangle)`

SetAbsoluteRenderBounds sets AbsoluteRenderBounds field to given value.


### GetPreserveRatio

`func (o *InstanceNode) GetPreserveRatio() bool`

GetPreserveRatio returns the PreserveRatio field if non-nil, zero value otherwise.

### GetPreserveRatioOk

`func (o *InstanceNode) GetPreserveRatioOk() (*bool, bool)`

GetPreserveRatioOk returns a tuple with the PreserveRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreserveRatio

`func (o *InstanceNode) SetPreserveRatio(v bool)`

SetPreserveRatio sets PreserveRatio field to given value.

### HasPreserveRatio

`func (o *InstanceNode) HasPreserveRatio() bool`

HasPreserveRatio returns a boolean if a field has been set.

### GetConstraints

`func (o *InstanceNode) GetConstraints() LayoutConstraint`

GetConstraints returns the Constraints field if non-nil, zero value otherwise.

### GetConstraintsOk

`func (o *InstanceNode) GetConstraintsOk() (*LayoutConstraint, bool)`

GetConstraintsOk returns a tuple with the Constraints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConstraints

`func (o *InstanceNode) SetConstraints(v LayoutConstraint)`

SetConstraints sets Constraints field to given value.

### HasConstraints

`func (o *InstanceNode) HasConstraints() bool`

HasConstraints returns a boolean if a field has been set.

### GetRelativeTransform

`func (o *InstanceNode) GetRelativeTransform() [][]float32`

GetRelativeTransform returns the RelativeTransform field if non-nil, zero value otherwise.

### GetRelativeTransformOk

`func (o *InstanceNode) GetRelativeTransformOk() (*[][]float32, bool)`

GetRelativeTransformOk returns a tuple with the RelativeTransform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelativeTransform

`func (o *InstanceNode) SetRelativeTransform(v [][]float32)`

SetRelativeTransform sets RelativeTransform field to given value.

### HasRelativeTransform

`func (o *InstanceNode) HasRelativeTransform() bool`

HasRelativeTransform returns a boolean if a field has been set.

### GetSize

`func (o *InstanceNode) GetSize() Vector`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *InstanceNode) GetSizeOk() (*Vector, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *InstanceNode) SetSize(v Vector)`

SetSize sets Size field to given value.

### HasSize

`func (o *InstanceNode) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetLayoutAlign

`func (o *InstanceNode) GetLayoutAlign() string`

GetLayoutAlign returns the LayoutAlign field if non-nil, zero value otherwise.

### GetLayoutAlignOk

`func (o *InstanceNode) GetLayoutAlignOk() (*string, bool)`

GetLayoutAlignOk returns a tuple with the LayoutAlign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayoutAlign

`func (o *InstanceNode) SetLayoutAlign(v string)`

SetLayoutAlign sets LayoutAlign field to given value.

### HasLayoutAlign

`func (o *InstanceNode) HasLayoutAlign() bool`

HasLayoutAlign returns a boolean if a field has been set.

### GetLayoutGrow

`func (o *InstanceNode) GetLayoutGrow() float32`

GetLayoutGrow returns the LayoutGrow field if non-nil, zero value otherwise.

### GetLayoutGrowOk

`func (o *InstanceNode) GetLayoutGrowOk() (*float32, bool)`

GetLayoutGrowOk returns a tuple with the LayoutGrow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayoutGrow

`func (o *InstanceNode) SetLayoutGrow(v float32)`

SetLayoutGrow sets LayoutGrow field to given value.

### HasLayoutGrow

`func (o *InstanceNode) HasLayoutGrow() bool`

HasLayoutGrow returns a boolean if a field has been set.

### GetLayoutPositioning

`func (o *InstanceNode) GetLayoutPositioning() string`

GetLayoutPositioning returns the LayoutPositioning field if non-nil, zero value otherwise.

### GetLayoutPositioningOk

`func (o *InstanceNode) GetLayoutPositioningOk() (*string, bool)`

GetLayoutPositioningOk returns a tuple with the LayoutPositioning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayoutPositioning

`func (o *InstanceNode) SetLayoutPositioning(v string)`

SetLayoutPositioning sets LayoutPositioning field to given value.

### HasLayoutPositioning

`func (o *InstanceNode) HasLayoutPositioning() bool`

HasLayoutPositioning returns a boolean if a field has been set.

### GetMinWidth

`func (o *InstanceNode) GetMinWidth() float32`

GetMinWidth returns the MinWidth field if non-nil, zero value otherwise.

### GetMinWidthOk

`func (o *InstanceNode) GetMinWidthOk() (*float32, bool)`

GetMinWidthOk returns a tuple with the MinWidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinWidth

`func (o *InstanceNode) SetMinWidth(v float32)`

SetMinWidth sets MinWidth field to given value.

### HasMinWidth

`func (o *InstanceNode) HasMinWidth() bool`

HasMinWidth returns a boolean if a field has been set.

### GetMaxWidth

`func (o *InstanceNode) GetMaxWidth() float32`

GetMaxWidth returns the MaxWidth field if non-nil, zero value otherwise.

### GetMaxWidthOk

`func (o *InstanceNode) GetMaxWidthOk() (*float32, bool)`

GetMaxWidthOk returns a tuple with the MaxWidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxWidth

`func (o *InstanceNode) SetMaxWidth(v float32)`

SetMaxWidth sets MaxWidth field to given value.

### HasMaxWidth

`func (o *InstanceNode) HasMaxWidth() bool`

HasMaxWidth returns a boolean if a field has been set.

### GetMinHeight

`func (o *InstanceNode) GetMinHeight() float32`

GetMinHeight returns the MinHeight field if non-nil, zero value otherwise.

### GetMinHeightOk

`func (o *InstanceNode) GetMinHeightOk() (*float32, bool)`

GetMinHeightOk returns a tuple with the MinHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinHeight

`func (o *InstanceNode) SetMinHeight(v float32)`

SetMinHeight sets MinHeight field to given value.

### HasMinHeight

`func (o *InstanceNode) HasMinHeight() bool`

HasMinHeight returns a boolean if a field has been set.

### GetMaxHeight

`func (o *InstanceNode) GetMaxHeight() float32`

GetMaxHeight returns the MaxHeight field if non-nil, zero value otherwise.

### GetMaxHeightOk

`func (o *InstanceNode) GetMaxHeightOk() (*float32, bool)`

GetMaxHeightOk returns a tuple with the MaxHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxHeight

`func (o *InstanceNode) SetMaxHeight(v float32)`

SetMaxHeight sets MaxHeight field to given value.

### HasMaxHeight

`func (o *InstanceNode) HasMaxHeight() bool`

HasMaxHeight returns a boolean if a field has been set.

### GetLayoutSizingHorizontal

`func (o *InstanceNode) GetLayoutSizingHorizontal() string`

GetLayoutSizingHorizontal returns the LayoutSizingHorizontal field if non-nil, zero value otherwise.

### GetLayoutSizingHorizontalOk

`func (o *InstanceNode) GetLayoutSizingHorizontalOk() (*string, bool)`

GetLayoutSizingHorizontalOk returns a tuple with the LayoutSizingHorizontal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayoutSizingHorizontal

`func (o *InstanceNode) SetLayoutSizingHorizontal(v string)`

SetLayoutSizingHorizontal sets LayoutSizingHorizontal field to given value.

### HasLayoutSizingHorizontal

`func (o *InstanceNode) HasLayoutSizingHorizontal() bool`

HasLayoutSizingHorizontal returns a boolean if a field has been set.

### GetLayoutSizingVertical

`func (o *InstanceNode) GetLayoutSizingVertical() string`

GetLayoutSizingVertical returns the LayoutSizingVertical field if non-nil, zero value otherwise.

### GetLayoutSizingVerticalOk

`func (o *InstanceNode) GetLayoutSizingVerticalOk() (*string, bool)`

GetLayoutSizingVerticalOk returns a tuple with the LayoutSizingVertical field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayoutSizingVertical

`func (o *InstanceNode) SetLayoutSizingVertical(v string)`

SetLayoutSizingVertical sets LayoutSizingVertical field to given value.

### HasLayoutSizingVertical

`func (o *InstanceNode) HasLayoutSizingVertical() bool`

HasLayoutSizingVertical returns a boolean if a field has been set.

### GetClipsContent

`func (o *InstanceNode) GetClipsContent() bool`

GetClipsContent returns the ClipsContent field if non-nil, zero value otherwise.

### GetClipsContentOk

`func (o *InstanceNode) GetClipsContentOk() (*bool, bool)`

GetClipsContentOk returns a tuple with the ClipsContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClipsContent

`func (o *InstanceNode) SetClipsContent(v bool)`

SetClipsContent sets ClipsContent field to given value.


### GetBackground

`func (o *InstanceNode) GetBackground() []Paint`

GetBackground returns the Background field if non-nil, zero value otherwise.

### GetBackgroundOk

`func (o *InstanceNode) GetBackgroundOk() (*[]Paint, bool)`

GetBackgroundOk returns a tuple with the Background field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackground

`func (o *InstanceNode) SetBackground(v []Paint)`

SetBackground sets Background field to given value.

### HasBackground

`func (o *InstanceNode) HasBackground() bool`

HasBackground returns a boolean if a field has been set.

### GetBackgroundColor

`func (o *InstanceNode) GetBackgroundColor() RGBA`

GetBackgroundColor returns the BackgroundColor field if non-nil, zero value otherwise.

### GetBackgroundColorOk

`func (o *InstanceNode) GetBackgroundColorOk() (*RGBA, bool)`

GetBackgroundColorOk returns a tuple with the BackgroundColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackgroundColor

`func (o *InstanceNode) SetBackgroundColor(v RGBA)`

SetBackgroundColor sets BackgroundColor field to given value.

### HasBackgroundColor

`func (o *InstanceNode) HasBackgroundColor() bool`

HasBackgroundColor returns a boolean if a field has been set.

### GetLayoutGrids

`func (o *InstanceNode) GetLayoutGrids() []LayoutGrid`

GetLayoutGrids returns the LayoutGrids field if non-nil, zero value otherwise.

### GetLayoutGridsOk

`func (o *InstanceNode) GetLayoutGridsOk() (*[]LayoutGrid, bool)`

GetLayoutGridsOk returns a tuple with the LayoutGrids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayoutGrids

`func (o *InstanceNode) SetLayoutGrids(v []LayoutGrid)`

SetLayoutGrids sets LayoutGrids field to given value.

### HasLayoutGrids

`func (o *InstanceNode) HasLayoutGrids() bool`

HasLayoutGrids returns a boolean if a field has been set.

### GetOverflowDirection

`func (o *InstanceNode) GetOverflowDirection() string`

GetOverflowDirection returns the OverflowDirection field if non-nil, zero value otherwise.

### GetOverflowDirectionOk

`func (o *InstanceNode) GetOverflowDirectionOk() (*string, bool)`

GetOverflowDirectionOk returns a tuple with the OverflowDirection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverflowDirection

`func (o *InstanceNode) SetOverflowDirection(v string)`

SetOverflowDirection sets OverflowDirection field to given value.

### HasOverflowDirection

`func (o *InstanceNode) HasOverflowDirection() bool`

HasOverflowDirection returns a boolean if a field has been set.

### GetLayoutMode

`func (o *InstanceNode) GetLayoutMode() string`

GetLayoutMode returns the LayoutMode field if non-nil, zero value otherwise.

### GetLayoutModeOk

`func (o *InstanceNode) GetLayoutModeOk() (*string, bool)`

GetLayoutModeOk returns a tuple with the LayoutMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayoutMode

`func (o *InstanceNode) SetLayoutMode(v string)`

SetLayoutMode sets LayoutMode field to given value.

### HasLayoutMode

`func (o *InstanceNode) HasLayoutMode() bool`

HasLayoutMode returns a boolean if a field has been set.

### GetPrimaryAxisSizingMode

`func (o *InstanceNode) GetPrimaryAxisSizingMode() string`

GetPrimaryAxisSizingMode returns the PrimaryAxisSizingMode field if non-nil, zero value otherwise.

### GetPrimaryAxisSizingModeOk

`func (o *InstanceNode) GetPrimaryAxisSizingModeOk() (*string, bool)`

GetPrimaryAxisSizingModeOk returns a tuple with the PrimaryAxisSizingMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryAxisSizingMode

`func (o *InstanceNode) SetPrimaryAxisSizingMode(v string)`

SetPrimaryAxisSizingMode sets PrimaryAxisSizingMode field to given value.

### HasPrimaryAxisSizingMode

`func (o *InstanceNode) HasPrimaryAxisSizingMode() bool`

HasPrimaryAxisSizingMode returns a boolean if a field has been set.

### GetCounterAxisSizingMode

`func (o *InstanceNode) GetCounterAxisSizingMode() string`

GetCounterAxisSizingMode returns the CounterAxisSizingMode field if non-nil, zero value otherwise.

### GetCounterAxisSizingModeOk

`func (o *InstanceNode) GetCounterAxisSizingModeOk() (*string, bool)`

GetCounterAxisSizingModeOk returns a tuple with the CounterAxisSizingMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounterAxisSizingMode

`func (o *InstanceNode) SetCounterAxisSizingMode(v string)`

SetCounterAxisSizingMode sets CounterAxisSizingMode field to given value.

### HasCounterAxisSizingMode

`func (o *InstanceNode) HasCounterAxisSizingMode() bool`

HasCounterAxisSizingMode returns a boolean if a field has been set.

### GetPrimaryAxisAlignItems

`func (o *InstanceNode) GetPrimaryAxisAlignItems() string`

GetPrimaryAxisAlignItems returns the PrimaryAxisAlignItems field if non-nil, zero value otherwise.

### GetPrimaryAxisAlignItemsOk

`func (o *InstanceNode) GetPrimaryAxisAlignItemsOk() (*string, bool)`

GetPrimaryAxisAlignItemsOk returns a tuple with the PrimaryAxisAlignItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryAxisAlignItems

`func (o *InstanceNode) SetPrimaryAxisAlignItems(v string)`

SetPrimaryAxisAlignItems sets PrimaryAxisAlignItems field to given value.

### HasPrimaryAxisAlignItems

`func (o *InstanceNode) HasPrimaryAxisAlignItems() bool`

HasPrimaryAxisAlignItems returns a boolean if a field has been set.

### GetCounterAxisAlignItems

`func (o *InstanceNode) GetCounterAxisAlignItems() string`

GetCounterAxisAlignItems returns the CounterAxisAlignItems field if non-nil, zero value otherwise.

### GetCounterAxisAlignItemsOk

`func (o *InstanceNode) GetCounterAxisAlignItemsOk() (*string, bool)`

GetCounterAxisAlignItemsOk returns a tuple with the CounterAxisAlignItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounterAxisAlignItems

`func (o *InstanceNode) SetCounterAxisAlignItems(v string)`

SetCounterAxisAlignItems sets CounterAxisAlignItems field to given value.

### HasCounterAxisAlignItems

`func (o *InstanceNode) HasCounterAxisAlignItems() bool`

HasCounterAxisAlignItems returns a boolean if a field has been set.

### GetPaddingLeft

`func (o *InstanceNode) GetPaddingLeft() float32`

GetPaddingLeft returns the PaddingLeft field if non-nil, zero value otherwise.

### GetPaddingLeftOk

`func (o *InstanceNode) GetPaddingLeftOk() (*float32, bool)`

GetPaddingLeftOk returns a tuple with the PaddingLeft field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaddingLeft

`func (o *InstanceNode) SetPaddingLeft(v float32)`

SetPaddingLeft sets PaddingLeft field to given value.

### HasPaddingLeft

`func (o *InstanceNode) HasPaddingLeft() bool`

HasPaddingLeft returns a boolean if a field has been set.

### GetPaddingRight

`func (o *InstanceNode) GetPaddingRight() float32`

GetPaddingRight returns the PaddingRight field if non-nil, zero value otherwise.

### GetPaddingRightOk

`func (o *InstanceNode) GetPaddingRightOk() (*float32, bool)`

GetPaddingRightOk returns a tuple with the PaddingRight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaddingRight

`func (o *InstanceNode) SetPaddingRight(v float32)`

SetPaddingRight sets PaddingRight field to given value.

### HasPaddingRight

`func (o *InstanceNode) HasPaddingRight() bool`

HasPaddingRight returns a boolean if a field has been set.

### GetPaddingTop

`func (o *InstanceNode) GetPaddingTop() float32`

GetPaddingTop returns the PaddingTop field if non-nil, zero value otherwise.

### GetPaddingTopOk

`func (o *InstanceNode) GetPaddingTopOk() (*float32, bool)`

GetPaddingTopOk returns a tuple with the PaddingTop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaddingTop

`func (o *InstanceNode) SetPaddingTop(v float32)`

SetPaddingTop sets PaddingTop field to given value.

### HasPaddingTop

`func (o *InstanceNode) HasPaddingTop() bool`

HasPaddingTop returns a boolean if a field has been set.

### GetPaddingBottom

`func (o *InstanceNode) GetPaddingBottom() float32`

GetPaddingBottom returns the PaddingBottom field if non-nil, zero value otherwise.

### GetPaddingBottomOk

`func (o *InstanceNode) GetPaddingBottomOk() (*float32, bool)`

GetPaddingBottomOk returns a tuple with the PaddingBottom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaddingBottom

`func (o *InstanceNode) SetPaddingBottom(v float32)`

SetPaddingBottom sets PaddingBottom field to given value.

### HasPaddingBottom

`func (o *InstanceNode) HasPaddingBottom() bool`

HasPaddingBottom returns a boolean if a field has been set.

### GetItemSpacing

`func (o *InstanceNode) GetItemSpacing() float32`

GetItemSpacing returns the ItemSpacing field if non-nil, zero value otherwise.

### GetItemSpacingOk

`func (o *InstanceNode) GetItemSpacingOk() (*float32, bool)`

GetItemSpacingOk returns a tuple with the ItemSpacing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemSpacing

`func (o *InstanceNode) SetItemSpacing(v float32)`

SetItemSpacing sets ItemSpacing field to given value.

### HasItemSpacing

`func (o *InstanceNode) HasItemSpacing() bool`

HasItemSpacing returns a boolean if a field has been set.

### GetItemReverseZIndex

`func (o *InstanceNode) GetItemReverseZIndex() bool`

GetItemReverseZIndex returns the ItemReverseZIndex field if non-nil, zero value otherwise.

### GetItemReverseZIndexOk

`func (o *InstanceNode) GetItemReverseZIndexOk() (*bool, bool)`

GetItemReverseZIndexOk returns a tuple with the ItemReverseZIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemReverseZIndex

`func (o *InstanceNode) SetItemReverseZIndex(v bool)`

SetItemReverseZIndex sets ItemReverseZIndex field to given value.

### HasItemReverseZIndex

`func (o *InstanceNode) HasItemReverseZIndex() bool`

HasItemReverseZIndex returns a boolean if a field has been set.

### GetStrokesIncludedInLayout

`func (o *InstanceNode) GetStrokesIncludedInLayout() bool`

GetStrokesIncludedInLayout returns the StrokesIncludedInLayout field if non-nil, zero value otherwise.

### GetStrokesIncludedInLayoutOk

`func (o *InstanceNode) GetStrokesIncludedInLayoutOk() (*bool, bool)`

GetStrokesIncludedInLayoutOk returns a tuple with the StrokesIncludedInLayout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrokesIncludedInLayout

`func (o *InstanceNode) SetStrokesIncludedInLayout(v bool)`

SetStrokesIncludedInLayout sets StrokesIncludedInLayout field to given value.

### HasStrokesIncludedInLayout

`func (o *InstanceNode) HasStrokesIncludedInLayout() bool`

HasStrokesIncludedInLayout returns a boolean if a field has been set.

### GetLayoutWrap

`func (o *InstanceNode) GetLayoutWrap() string`

GetLayoutWrap returns the LayoutWrap field if non-nil, zero value otherwise.

### GetLayoutWrapOk

`func (o *InstanceNode) GetLayoutWrapOk() (*string, bool)`

GetLayoutWrapOk returns a tuple with the LayoutWrap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayoutWrap

`func (o *InstanceNode) SetLayoutWrap(v string)`

SetLayoutWrap sets LayoutWrap field to given value.

### HasLayoutWrap

`func (o *InstanceNode) HasLayoutWrap() bool`

HasLayoutWrap returns a boolean if a field has been set.

### GetCounterAxisSpacing

`func (o *InstanceNode) GetCounterAxisSpacing() float32`

GetCounterAxisSpacing returns the CounterAxisSpacing field if non-nil, zero value otherwise.

### GetCounterAxisSpacingOk

`func (o *InstanceNode) GetCounterAxisSpacingOk() (*float32, bool)`

GetCounterAxisSpacingOk returns a tuple with the CounterAxisSpacing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounterAxisSpacing

`func (o *InstanceNode) SetCounterAxisSpacing(v float32)`

SetCounterAxisSpacing sets CounterAxisSpacing field to given value.

### HasCounterAxisSpacing

`func (o *InstanceNode) HasCounterAxisSpacing() bool`

HasCounterAxisSpacing returns a boolean if a field has been set.

### GetCounterAxisAlignContent

`func (o *InstanceNode) GetCounterAxisAlignContent() string`

GetCounterAxisAlignContent returns the CounterAxisAlignContent field if non-nil, zero value otherwise.

### GetCounterAxisAlignContentOk

`func (o *InstanceNode) GetCounterAxisAlignContentOk() (*string, bool)`

GetCounterAxisAlignContentOk returns a tuple with the CounterAxisAlignContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounterAxisAlignContent

`func (o *InstanceNode) SetCounterAxisAlignContent(v string)`

SetCounterAxisAlignContent sets CounterAxisAlignContent field to given value.

### HasCounterAxisAlignContent

`func (o *InstanceNode) HasCounterAxisAlignContent() bool`

HasCounterAxisAlignContent returns a boolean if a field has been set.

### GetCornerRadius

`func (o *InstanceNode) GetCornerRadius() float32`

GetCornerRadius returns the CornerRadius field if non-nil, zero value otherwise.

### GetCornerRadiusOk

`func (o *InstanceNode) GetCornerRadiusOk() (*float32, bool)`

GetCornerRadiusOk returns a tuple with the CornerRadius field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCornerRadius

`func (o *InstanceNode) SetCornerRadius(v float32)`

SetCornerRadius sets CornerRadius field to given value.

### HasCornerRadius

`func (o *InstanceNode) HasCornerRadius() bool`

HasCornerRadius returns a boolean if a field has been set.

### GetCornerSmoothing

`func (o *InstanceNode) GetCornerSmoothing() float32`

GetCornerSmoothing returns the CornerSmoothing field if non-nil, zero value otherwise.

### GetCornerSmoothingOk

`func (o *InstanceNode) GetCornerSmoothingOk() (*float32, bool)`

GetCornerSmoothingOk returns a tuple with the CornerSmoothing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCornerSmoothing

`func (o *InstanceNode) SetCornerSmoothing(v float32)`

SetCornerSmoothing sets CornerSmoothing field to given value.

### HasCornerSmoothing

`func (o *InstanceNode) HasCornerSmoothing() bool`

HasCornerSmoothing returns a boolean if a field has been set.

### GetRectangleCornerRadii

`func (o *InstanceNode) GetRectangleCornerRadii() []float32`

GetRectangleCornerRadii returns the RectangleCornerRadii field if non-nil, zero value otherwise.

### GetRectangleCornerRadiiOk

`func (o *InstanceNode) GetRectangleCornerRadiiOk() (*[]float32, bool)`

GetRectangleCornerRadiiOk returns a tuple with the RectangleCornerRadii field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRectangleCornerRadii

`func (o *InstanceNode) SetRectangleCornerRadii(v []float32)`

SetRectangleCornerRadii sets RectangleCornerRadii field to given value.

### HasRectangleCornerRadii

`func (o *InstanceNode) HasRectangleCornerRadii() bool`

HasRectangleCornerRadii returns a boolean if a field has been set.

### GetFills

`func (o *InstanceNode) GetFills() []Paint`

GetFills returns the Fills field if non-nil, zero value otherwise.

### GetFillsOk

`func (o *InstanceNode) GetFillsOk() (*[]Paint, bool)`

GetFillsOk returns a tuple with the Fills field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFills

`func (o *InstanceNode) SetFills(v []Paint)`

SetFills sets Fills field to given value.


### GetStyles

`func (o *InstanceNode) GetStyles() map[string]string`

GetStyles returns the Styles field if non-nil, zero value otherwise.

### GetStylesOk

`func (o *InstanceNode) GetStylesOk() (*map[string]string, bool)`

GetStylesOk returns a tuple with the Styles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStyles

`func (o *InstanceNode) SetStyles(v map[string]string)`

SetStyles sets Styles field to given value.

### HasStyles

`func (o *InstanceNode) HasStyles() bool`

HasStyles returns a boolean if a field has been set.

### GetStrokes

`func (o *InstanceNode) GetStrokes() []Paint`

GetStrokes returns the Strokes field if non-nil, zero value otherwise.

### GetStrokesOk

`func (o *InstanceNode) GetStrokesOk() (*[]Paint, bool)`

GetStrokesOk returns a tuple with the Strokes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrokes

`func (o *InstanceNode) SetStrokes(v []Paint)`

SetStrokes sets Strokes field to given value.

### HasStrokes

`func (o *InstanceNode) HasStrokes() bool`

HasStrokes returns a boolean if a field has been set.

### GetStrokeWeight

`func (o *InstanceNode) GetStrokeWeight() float32`

GetStrokeWeight returns the StrokeWeight field if non-nil, zero value otherwise.

### GetStrokeWeightOk

`func (o *InstanceNode) GetStrokeWeightOk() (*float32, bool)`

GetStrokeWeightOk returns a tuple with the StrokeWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrokeWeight

`func (o *InstanceNode) SetStrokeWeight(v float32)`

SetStrokeWeight sets StrokeWeight field to given value.

### HasStrokeWeight

`func (o *InstanceNode) HasStrokeWeight() bool`

HasStrokeWeight returns a boolean if a field has been set.

### GetStrokeAlign

`func (o *InstanceNode) GetStrokeAlign() string`

GetStrokeAlign returns the StrokeAlign field if non-nil, zero value otherwise.

### GetStrokeAlignOk

`func (o *InstanceNode) GetStrokeAlignOk() (*string, bool)`

GetStrokeAlignOk returns a tuple with the StrokeAlign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrokeAlign

`func (o *InstanceNode) SetStrokeAlign(v string)`

SetStrokeAlign sets StrokeAlign field to given value.

### HasStrokeAlign

`func (o *InstanceNode) HasStrokeAlign() bool`

HasStrokeAlign returns a boolean if a field has been set.

### GetStrokeJoin

`func (o *InstanceNode) GetStrokeJoin() string`

GetStrokeJoin returns the StrokeJoin field if non-nil, zero value otherwise.

### GetStrokeJoinOk

`func (o *InstanceNode) GetStrokeJoinOk() (*string, bool)`

GetStrokeJoinOk returns a tuple with the StrokeJoin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrokeJoin

`func (o *InstanceNode) SetStrokeJoin(v string)`

SetStrokeJoin sets StrokeJoin field to given value.

### HasStrokeJoin

`func (o *InstanceNode) HasStrokeJoin() bool`

HasStrokeJoin returns a boolean if a field has been set.

### GetStrokeDashes

`func (o *InstanceNode) GetStrokeDashes() []float32`

GetStrokeDashes returns the StrokeDashes field if non-nil, zero value otherwise.

### GetStrokeDashesOk

`func (o *InstanceNode) GetStrokeDashesOk() (*[]float32, bool)`

GetStrokeDashesOk returns a tuple with the StrokeDashes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrokeDashes

`func (o *InstanceNode) SetStrokeDashes(v []float32)`

SetStrokeDashes sets StrokeDashes field to given value.

### HasStrokeDashes

`func (o *InstanceNode) HasStrokeDashes() bool`

HasStrokeDashes returns a boolean if a field has been set.

### GetFillOverrideTable

`func (o *InstanceNode) GetFillOverrideTable() map[string]HasGeometryTraitAllOfFillOverrideTable`

GetFillOverrideTable returns the FillOverrideTable field if non-nil, zero value otherwise.

### GetFillOverrideTableOk

`func (o *InstanceNode) GetFillOverrideTableOk() (*map[string]HasGeometryTraitAllOfFillOverrideTable, bool)`

GetFillOverrideTableOk returns a tuple with the FillOverrideTable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFillOverrideTable

`func (o *InstanceNode) SetFillOverrideTable(v map[string]HasGeometryTraitAllOfFillOverrideTable)`

SetFillOverrideTable sets FillOverrideTable field to given value.

### HasFillOverrideTable

`func (o *InstanceNode) HasFillOverrideTable() bool`

HasFillOverrideTable returns a boolean if a field has been set.

### GetFillGeometry

`func (o *InstanceNode) GetFillGeometry() []Path`

GetFillGeometry returns the FillGeometry field if non-nil, zero value otherwise.

### GetFillGeometryOk

`func (o *InstanceNode) GetFillGeometryOk() (*[]Path, bool)`

GetFillGeometryOk returns a tuple with the FillGeometry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFillGeometry

`func (o *InstanceNode) SetFillGeometry(v []Path)`

SetFillGeometry sets FillGeometry field to given value.

### HasFillGeometry

`func (o *InstanceNode) HasFillGeometry() bool`

HasFillGeometry returns a boolean if a field has been set.

### GetStrokeGeometry

`func (o *InstanceNode) GetStrokeGeometry() []Path`

GetStrokeGeometry returns the StrokeGeometry field if non-nil, zero value otherwise.

### GetStrokeGeometryOk

`func (o *InstanceNode) GetStrokeGeometryOk() (*[]Path, bool)`

GetStrokeGeometryOk returns a tuple with the StrokeGeometry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrokeGeometry

`func (o *InstanceNode) SetStrokeGeometry(v []Path)`

SetStrokeGeometry sets StrokeGeometry field to given value.

### HasStrokeGeometry

`func (o *InstanceNode) HasStrokeGeometry() bool`

HasStrokeGeometry returns a boolean if a field has been set.

### GetStrokeCap

`func (o *InstanceNode) GetStrokeCap() string`

GetStrokeCap returns the StrokeCap field if non-nil, zero value otherwise.

### GetStrokeCapOk

`func (o *InstanceNode) GetStrokeCapOk() (*string, bool)`

GetStrokeCapOk returns a tuple with the StrokeCap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrokeCap

`func (o *InstanceNode) SetStrokeCap(v string)`

SetStrokeCap sets StrokeCap field to given value.

### HasStrokeCap

`func (o *InstanceNode) HasStrokeCap() bool`

HasStrokeCap returns a boolean if a field has been set.

### GetStrokeMiterAngle

`func (o *InstanceNode) GetStrokeMiterAngle() float32`

GetStrokeMiterAngle returns the StrokeMiterAngle field if non-nil, zero value otherwise.

### GetStrokeMiterAngleOk

`func (o *InstanceNode) GetStrokeMiterAngleOk() (*float32, bool)`

GetStrokeMiterAngleOk returns a tuple with the StrokeMiterAngle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrokeMiterAngle

`func (o *InstanceNode) SetStrokeMiterAngle(v float32)`

SetStrokeMiterAngle sets StrokeMiterAngle field to given value.

### HasStrokeMiterAngle

`func (o *InstanceNode) HasStrokeMiterAngle() bool`

HasStrokeMiterAngle returns a boolean if a field has been set.

### GetExportSettings

`func (o *InstanceNode) GetExportSettings() []ExportSetting`

GetExportSettings returns the ExportSettings field if non-nil, zero value otherwise.

### GetExportSettingsOk

`func (o *InstanceNode) GetExportSettingsOk() (*[]ExportSetting, bool)`

GetExportSettingsOk returns a tuple with the ExportSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportSettings

`func (o *InstanceNode) SetExportSettings(v []ExportSetting)`

SetExportSettings sets ExportSettings field to given value.

### HasExportSettings

`func (o *InstanceNode) HasExportSettings() bool`

HasExportSettings returns a boolean if a field has been set.

### GetEffects

`func (o *InstanceNode) GetEffects() []Effect`

GetEffects returns the Effects field if non-nil, zero value otherwise.

### GetEffectsOk

`func (o *InstanceNode) GetEffectsOk() (*[]Effect, bool)`

GetEffectsOk returns a tuple with the Effects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffects

`func (o *InstanceNode) SetEffects(v []Effect)`

SetEffects sets Effects field to given value.


### GetIsMask

`func (o *InstanceNode) GetIsMask() bool`

GetIsMask returns the IsMask field if non-nil, zero value otherwise.

### GetIsMaskOk

`func (o *InstanceNode) GetIsMaskOk() (*bool, bool)`

GetIsMaskOk returns a tuple with the IsMask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMask

`func (o *InstanceNode) SetIsMask(v bool)`

SetIsMask sets IsMask field to given value.

### HasIsMask

`func (o *InstanceNode) HasIsMask() bool`

HasIsMask returns a boolean if a field has been set.

### GetMaskType

`func (o *InstanceNode) GetMaskType() string`

GetMaskType returns the MaskType field if non-nil, zero value otherwise.

### GetMaskTypeOk

`func (o *InstanceNode) GetMaskTypeOk() (*string, bool)`

GetMaskTypeOk returns a tuple with the MaskType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaskType

`func (o *InstanceNode) SetMaskType(v string)`

SetMaskType sets MaskType field to given value.

### HasMaskType

`func (o *InstanceNode) HasMaskType() bool`

HasMaskType returns a boolean if a field has been set.

### GetIsMaskOutline

`func (o *InstanceNode) GetIsMaskOutline() bool`

GetIsMaskOutline returns the IsMaskOutline field if non-nil, zero value otherwise.

### GetIsMaskOutlineOk

`func (o *InstanceNode) GetIsMaskOutlineOk() (*bool, bool)`

GetIsMaskOutlineOk returns a tuple with the IsMaskOutline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMaskOutline

`func (o *InstanceNode) SetIsMaskOutline(v bool)`

SetIsMaskOutline sets IsMaskOutline field to given value.

### HasIsMaskOutline

`func (o *InstanceNode) HasIsMaskOutline() bool`

HasIsMaskOutline returns a boolean if a field has been set.

### GetTransitionNodeID

`func (o *InstanceNode) GetTransitionNodeID() string`

GetTransitionNodeID returns the TransitionNodeID field if non-nil, zero value otherwise.

### GetTransitionNodeIDOk

`func (o *InstanceNode) GetTransitionNodeIDOk() (*string, bool)`

GetTransitionNodeIDOk returns a tuple with the TransitionNodeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransitionNodeID

`func (o *InstanceNode) SetTransitionNodeID(v string)`

SetTransitionNodeID sets TransitionNodeID field to given value.

### HasTransitionNodeID

`func (o *InstanceNode) HasTransitionNodeID() bool`

HasTransitionNodeID returns a boolean if a field has been set.

### GetTransitionDuration

`func (o *InstanceNode) GetTransitionDuration() float32`

GetTransitionDuration returns the TransitionDuration field if non-nil, zero value otherwise.

### GetTransitionDurationOk

`func (o *InstanceNode) GetTransitionDurationOk() (*float32, bool)`

GetTransitionDurationOk returns a tuple with the TransitionDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransitionDuration

`func (o *InstanceNode) SetTransitionDuration(v float32)`

SetTransitionDuration sets TransitionDuration field to given value.

### HasTransitionDuration

`func (o *InstanceNode) HasTransitionDuration() bool`

HasTransitionDuration returns a boolean if a field has been set.

### GetTransitionEasing

`func (o *InstanceNode) GetTransitionEasing() EasingType`

GetTransitionEasing returns the TransitionEasing field if non-nil, zero value otherwise.

### GetTransitionEasingOk

`func (o *InstanceNode) GetTransitionEasingOk() (*EasingType, bool)`

GetTransitionEasingOk returns a tuple with the TransitionEasing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransitionEasing

`func (o *InstanceNode) SetTransitionEasing(v EasingType)`

SetTransitionEasing sets TransitionEasing field to given value.

### HasTransitionEasing

`func (o *InstanceNode) HasTransitionEasing() bool`

HasTransitionEasing returns a boolean if a field has been set.

### GetInteractions

`func (o *InstanceNode) GetInteractions() []Interaction`

GetInteractions returns the Interactions field if non-nil, zero value otherwise.

### GetInteractionsOk

`func (o *InstanceNode) GetInteractionsOk() (*[]Interaction, bool)`

GetInteractionsOk returns a tuple with the Interactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInteractions

`func (o *InstanceNode) SetInteractions(v []Interaction)`

SetInteractions sets Interactions field to given value.

### HasInteractions

`func (o *InstanceNode) HasInteractions() bool`

HasInteractions returns a boolean if a field has been set.

### GetIndividualStrokeWeights

`func (o *InstanceNode) GetIndividualStrokeWeights() StrokeWeights`

GetIndividualStrokeWeights returns the IndividualStrokeWeights field if non-nil, zero value otherwise.

### GetIndividualStrokeWeightsOk

`func (o *InstanceNode) GetIndividualStrokeWeightsOk() (*StrokeWeights, bool)`

GetIndividualStrokeWeightsOk returns a tuple with the IndividualStrokeWeights field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndividualStrokeWeights

`func (o *InstanceNode) SetIndividualStrokeWeights(v StrokeWeights)`

SetIndividualStrokeWeights sets IndividualStrokeWeights field to given value.

### HasIndividualStrokeWeights

`func (o *InstanceNode) HasIndividualStrokeWeights() bool`

HasIndividualStrokeWeights returns a boolean if a field has been set.

### GetDevStatus

`func (o *InstanceNode) GetDevStatus() DevStatusTraitDevStatus`

GetDevStatus returns the DevStatus field if non-nil, zero value otherwise.

### GetDevStatusOk

`func (o *InstanceNode) GetDevStatusOk() (*DevStatusTraitDevStatus, bool)`

GetDevStatusOk returns a tuple with the DevStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevStatus

`func (o *InstanceNode) SetDevStatus(v DevStatusTraitDevStatus)`

SetDevStatus sets DevStatus field to given value.

### HasDevStatus

`func (o *InstanceNode) HasDevStatus() bool`

HasDevStatus returns a boolean if a field has been set.

### GetComponentId

`func (o *InstanceNode) GetComponentId() string`

GetComponentId returns the ComponentId field if non-nil, zero value otherwise.

### GetComponentIdOk

`func (o *InstanceNode) GetComponentIdOk() (*string, bool)`

GetComponentIdOk returns a tuple with the ComponentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentId

`func (o *InstanceNode) SetComponentId(v string)`

SetComponentId sets ComponentId field to given value.


### GetIsExposedInstance

`func (o *InstanceNode) GetIsExposedInstance() bool`

GetIsExposedInstance returns the IsExposedInstance field if non-nil, zero value otherwise.

### GetIsExposedInstanceOk

`func (o *InstanceNode) GetIsExposedInstanceOk() (*bool, bool)`

GetIsExposedInstanceOk returns a tuple with the IsExposedInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsExposedInstance

`func (o *InstanceNode) SetIsExposedInstance(v bool)`

SetIsExposedInstance sets IsExposedInstance field to given value.

### HasIsExposedInstance

`func (o *InstanceNode) HasIsExposedInstance() bool`

HasIsExposedInstance returns a boolean if a field has been set.

### GetExposedInstances

`func (o *InstanceNode) GetExposedInstances() []string`

GetExposedInstances returns the ExposedInstances field if non-nil, zero value otherwise.

### GetExposedInstancesOk

`func (o *InstanceNode) GetExposedInstancesOk() (*[]string, bool)`

GetExposedInstancesOk returns a tuple with the ExposedInstances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExposedInstances

`func (o *InstanceNode) SetExposedInstances(v []string)`

SetExposedInstances sets ExposedInstances field to given value.

### HasExposedInstances

`func (o *InstanceNode) HasExposedInstances() bool`

HasExposedInstances returns a boolean if a field has been set.

### GetComponentProperties

`func (o *InstanceNode) GetComponentProperties() map[string]ComponentProperty`

GetComponentProperties returns the ComponentProperties field if non-nil, zero value otherwise.

### GetComponentPropertiesOk

`func (o *InstanceNode) GetComponentPropertiesOk() (*map[string]ComponentProperty, bool)`

GetComponentPropertiesOk returns a tuple with the ComponentProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentProperties

`func (o *InstanceNode) SetComponentProperties(v map[string]ComponentProperty)`

SetComponentProperties sets ComponentProperties field to given value.

### HasComponentProperties

`func (o *InstanceNode) HasComponentProperties() bool`

HasComponentProperties returns a boolean if a field has been set.

### GetOverrides

`func (o *InstanceNode) GetOverrides() []Overrides`

GetOverrides returns the Overrides field if non-nil, zero value otherwise.

### GetOverridesOk

`func (o *InstanceNode) GetOverridesOk() (*[]Overrides, bool)`

GetOverridesOk returns a tuple with the Overrides field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrides

`func (o *InstanceNode) SetOverrides(v []Overrides)`

SetOverrides sets Overrides field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


