# ConnectorEndpointOneOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EndpointNodeId** | Pointer to **string** | Node ID that this endpoint attaches to. | [optional] 
**Position** | Pointer to [**Vector**](Vector.md) | The position of the endpoint relative to the node. | [optional] 

## Methods

### NewConnectorEndpointOneOf

`func NewConnectorEndpointOneOf() *ConnectorEndpointOneOf`

NewConnectorEndpointOneOf instantiates a new ConnectorEndpointOneOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectorEndpointOneOfWithDefaults

`func NewConnectorEndpointOneOfWithDefaults() *ConnectorEndpointOneOf`

NewConnectorEndpointOneOfWithDefaults instantiates a new ConnectorEndpointOneOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndpointNodeId

`func (o *ConnectorEndpointOneOf) GetEndpointNodeId() string`

GetEndpointNodeId returns the EndpointNodeId field if non-nil, zero value otherwise.

### GetEndpointNodeIdOk

`func (o *ConnectorEndpointOneOf) GetEndpointNodeIdOk() (*string, bool)`

GetEndpointNodeIdOk returns a tuple with the EndpointNodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointNodeId

`func (o *ConnectorEndpointOneOf) SetEndpointNodeId(v string)`

SetEndpointNodeId sets EndpointNodeId field to given value.

### HasEndpointNodeId

`func (o *ConnectorEndpointOneOf) HasEndpointNodeId() bool`

HasEndpointNodeId returns a boolean if a field has been set.

### GetPosition

`func (o *ConnectorEndpointOneOf) GetPosition() Vector`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *ConnectorEndpointOneOf) GetPositionOk() (*Vector, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *ConnectorEndpointOneOf) SetPosition(v Vector)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *ConnectorEndpointOneOf) HasPosition() bool`

HasPosition returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


