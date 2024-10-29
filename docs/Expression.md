# Expression

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExpressionFunction** | [**ExpressionFunction**](ExpressionFunction.md) |  | 
**ExpressionArguments** | [**[]VariableData**](VariableData.md) |  | 

## Methods

### NewExpression

`func NewExpression(expressionFunction ExpressionFunction, expressionArguments []VariableData, ) *Expression`

NewExpression instantiates a new Expression object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExpressionWithDefaults

`func NewExpressionWithDefaults() *Expression`

NewExpressionWithDefaults instantiates a new Expression object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpressionFunction

`func (o *Expression) GetExpressionFunction() ExpressionFunction`

GetExpressionFunction returns the ExpressionFunction field if non-nil, zero value otherwise.

### GetExpressionFunctionOk

`func (o *Expression) GetExpressionFunctionOk() (*ExpressionFunction, bool)`

GetExpressionFunctionOk returns a tuple with the ExpressionFunction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpressionFunction

`func (o *Expression) SetExpressionFunction(v ExpressionFunction)`

SetExpressionFunction sets ExpressionFunction field to given value.


### GetExpressionArguments

`func (o *Expression) GetExpressionArguments() []VariableData`

GetExpressionArguments returns the ExpressionArguments field if non-nil, zero value otherwise.

### GetExpressionArgumentsOk

`func (o *Expression) GetExpressionArgumentsOk() (*[]VariableData, bool)`

GetExpressionArgumentsOk returns a tuple with the ExpressionArguments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpressionArguments

`func (o *Expression) SetExpressionArguments(v []VariableData)`

SetExpressionArguments sets ExpressionArguments field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


