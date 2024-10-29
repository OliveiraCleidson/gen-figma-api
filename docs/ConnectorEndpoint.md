# ConnectorEndpoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EndpointNodeId** | Pointer to **string** | Node ID that this endpoint attaches to. | [optional] 
**Position** | Pointer to [**Vector**](Vector.md) | The position of the endpoint relative to the node. | [optional] 
**Magnet** | Pointer to **string** | The magnet type is a string enum. | [optional] 

## Methods

### NewConnectorEndpoint

`func NewConnectorEndpoint() *ConnectorEndpoint`

NewConnectorEndpoint instantiates a new ConnectorEndpoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectorEndpointWithDefaults

`func NewConnectorEndpointWithDefaults() *ConnectorEndpoint`

NewConnectorEndpointWithDefaults instantiates a new ConnectorEndpoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndpointNodeId

`func (o *ConnectorEndpoint) GetEndpointNodeId() string`

GetEndpointNodeId returns the EndpointNodeId field if non-nil, zero value otherwise.

### GetEndpointNodeIdOk

`func (o *ConnectorEndpoint) GetEndpointNodeIdOk() (*string, bool)`

GetEndpointNodeIdOk returns a tuple with the EndpointNodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointNodeId

`func (o *ConnectorEndpoint) SetEndpointNodeId(v string)`

SetEndpointNodeId sets EndpointNodeId field to given value.

### HasEndpointNodeId

`func (o *ConnectorEndpoint) HasEndpointNodeId() bool`

HasEndpointNodeId returns a boolean if a field has been set.

### GetPosition

`func (o *ConnectorEndpoint) GetPosition() Vector`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *ConnectorEndpoint) GetPositionOk() (*Vector, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *ConnectorEndpoint) SetPosition(v Vector)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *ConnectorEndpoint) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### GetMagnet

`func (o *ConnectorEndpoint) GetMagnet() string`

GetMagnet returns the Magnet field if non-nil, zero value otherwise.

### GetMagnetOk

`func (o *ConnectorEndpoint) GetMagnetOk() (*string, bool)`

GetMagnetOk returns a tuple with the Magnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMagnet

`func (o *ConnectorEndpoint) SetMagnet(v string)`

SetMagnet sets Magnet field to given value.

### HasMagnet

`func (o *ConnectorEndpoint) HasMagnet() bool`

HasMagnet returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


