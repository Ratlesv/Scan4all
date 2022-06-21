// This file is generated by "./lib/proto/generate"

package proto

import (
	"github.com/ysmood/gson"
)

/*

Overlay

This domain provides various functionality related to drawing atop the inspected page.

*/

// OverlaySourceOrderConfig Configuration data for drawing the source order of an elements children.
type OverlaySourceOrderConfig struct {

	// ParentOutlineColor the color to outline the givent element in.
	ParentOutlineColor *DOMRGBA `json:"parentOutlineColor"`

	// ChildOutlineColor the color to outline the child elements in.
	ChildOutlineColor *DOMRGBA `json:"childOutlineColor"`
}

// OverlayGridHighlightConfig Configuration data for the highlighting of Grid elements.
type OverlayGridHighlightConfig struct {

	// ShowGridExtensionLines (optional) Whether the extension lines from grid cells to the rulers should be shown (default: false).
	ShowGridExtensionLines bool `json:"showGridExtensionLines,omitempty"`

	// ShowPositiveLineNumbers (optional) Show Positive line number labels (default: false).
	ShowPositiveLineNumbers bool `json:"showPositiveLineNumbers,omitempty"`

	// ShowNegativeLineNumbers (optional) Show Negative line number labels (default: false).
	ShowNegativeLineNumbers bool `json:"showNegativeLineNumbers,omitempty"`

	// ShowAreaNames (optional) Show area name labels (default: false).
	ShowAreaNames bool `json:"showAreaNames,omitempty"`

	// ShowLineNames (optional) Show line name labels (default: false).
	ShowLineNames bool `json:"showLineNames,omitempty"`

	// ShowTrackSizes (optional) Show track size labels (default: false).
	ShowTrackSizes bool `json:"showTrackSizes,omitempty"`

	// GridBorderColor (optional) The grid container border highlight color (default: transparent).
	GridBorderColor *DOMRGBA `json:"gridBorderColor,omitempty"`

	// CellBorderColor (deprecated) (optional) The cell border color (default: transparent). Deprecated, please use rowLineColor and columnLineColor instead.
	CellBorderColor *DOMRGBA `json:"cellBorderColor,omitempty"`

	// RowLineColor (optional) The row line color (default: transparent).
	RowLineColor *DOMRGBA `json:"rowLineColor,omitempty"`

	// ColumnLineColor (optional) The column line color (default: transparent).
	ColumnLineColor *DOMRGBA `json:"columnLineColor,omitempty"`

	// GridBorderDash (optional) Whether the grid border is dashed (default: false).
	GridBorderDash bool `json:"gridBorderDash,omitempty"`

	// CellBorderDash (deprecated) (optional) Whether the cell border is dashed (default: false). Deprecated, please us rowLineDash and columnLineDash instead.
	CellBorderDash bool `json:"cellBorderDash,omitempty"`

	// RowLineDash (optional) Whether row lines are dashed (default: false).
	RowLineDash bool `json:"rowLineDash,omitempty"`

	// ColumnLineDash (optional) Whether column lines are dashed (default: false).
	ColumnLineDash bool `json:"columnLineDash,omitempty"`

	// RowGapColor (optional) The row gap highlight fill color (default: transparent).
	RowGapColor *DOMRGBA `json:"rowGapColor,omitempty"`

	// RowHatchColor (optional) The row gap hatching fill color (default: transparent).
	RowHatchColor *DOMRGBA `json:"rowHatchColor,omitempty"`

	// ColumnGapColor (optional) The column gap highlight fill color (default: transparent).
	ColumnGapColor *DOMRGBA `json:"columnGapColor,omitempty"`

	// ColumnHatchColor (optional) The column gap hatching fill color (default: transparent).
	ColumnHatchColor *DOMRGBA `json:"columnHatchColor,omitempty"`

	// AreaBorderColor (optional) The named grid areas border color (Default: transparent).
	AreaBorderColor *DOMRGBA `json:"areaBorderColor,omitempty"`

	// GridBackgroundColor (optional) The grid container background color (Default: transparent).
	GridBackgroundColor *DOMRGBA `json:"gridBackgroundColor,omitempty"`
}

// OverlayFlexContainerHighlightConfig Configuration data for the highlighting of Flex container elements.
type OverlayFlexContainerHighlightConfig struct {

	// ContainerBorder (optional) The style of the container border
	ContainerBorder *OverlayLineStyle `json:"containerBorder,omitempty"`

	// LineSeparator (optional) The style of the separator between lines
	LineSeparator *OverlayLineStyle `json:"lineSeparator,omitempty"`

	// ItemSeparator (optional) The style of the separator between items
	ItemSeparator *OverlayLineStyle `json:"itemSeparator,omitempty"`

	// MainDistributedSpace (optional) Style of content-distribution space on the main axis (justify-content).
	MainDistributedSpace *OverlayBoxStyle `json:"mainDistributedSpace,omitempty"`

	// CrossDistributedSpace (optional) Style of content-distribution space on the cross axis (align-content).
	CrossDistributedSpace *OverlayBoxStyle `json:"crossDistributedSpace,omitempty"`

	// RowGapSpace (optional) Style of empty space caused by row gaps (gap/row-gap).
	RowGapSpace *OverlayBoxStyle `json:"rowGapSpace,omitempty"`

	// ColumnGapSpace (optional) Style of empty space caused by columns gaps (gap/column-gap).
	ColumnGapSpace *OverlayBoxStyle `json:"columnGapSpace,omitempty"`

	// CrossAlignment (optional) Style of the self-alignment line (align-items).
	CrossAlignment *OverlayLineStyle `json:"crossAlignment,omitempty"`
}

// OverlayFlexItemHighlightConfig Configuration data for the highlighting of Flex item elements.
type OverlayFlexItemHighlightConfig struct {

	// BaseSizeBox (optional) Style of the box representing the item's base size
	BaseSizeBox *OverlayBoxStyle `json:"baseSizeBox,omitempty"`

	// BaseSizeBorder (optional) Style of the border around the box representing the item's base size
	BaseSizeBorder *OverlayLineStyle `json:"baseSizeBorder,omitempty"`

	// FlexibilityArrow (optional) Style of the arrow representing if the item grew or shrank
	FlexibilityArrow *OverlayLineStyle `json:"flexibilityArrow,omitempty"`
}

// OverlayLineStylePattern enum
type OverlayLineStylePattern string

const (
	// OverlayLineStylePatternDashed enum const
	OverlayLineStylePatternDashed OverlayLineStylePattern = "dashed"

	// OverlayLineStylePatternDotted enum const
	OverlayLineStylePatternDotted OverlayLineStylePattern = "dotted"
)

// OverlayLineStyle Style information for drawing a line.
type OverlayLineStyle struct {

	// Color (optional) The color of the line (default: transparent)
	Color *DOMRGBA `json:"color,omitempty"`

	// Pattern (optional) The line pattern (default: solid)
	Pattern OverlayLineStylePattern `json:"pattern,omitempty"`
}

// OverlayBoxStyle Style information for drawing a box.
type OverlayBoxStyle struct {

	// FillColor (optional) The background color for the box (default: transparent)
	FillColor *DOMRGBA `json:"fillColor,omitempty"`

	// HatchColor (optional) The hatching color for the box (default: transparent)
	HatchColor *DOMRGBA `json:"hatchColor,omitempty"`
}

// OverlayContrastAlgorithm ...
type OverlayContrastAlgorithm string

const (
	// OverlayContrastAlgorithmAa enum const
	OverlayContrastAlgorithmAa OverlayContrastAlgorithm = "aa"

	// OverlayContrastAlgorithmAaa enum const
	OverlayContrastAlgorithmAaa OverlayContrastAlgorithm = "aaa"

	// OverlayContrastAlgorithmApca enum const
	OverlayContrastAlgorithmApca OverlayContrastAlgorithm = "apca"
)

// OverlayHighlightConfig Configuration data for the highlighting of page elements.
type OverlayHighlightConfig struct {

	// ShowInfo (optional) Whether the node info tooltip should be shown (default: false).
	ShowInfo bool `json:"showInfo,omitempty"`

	// ShowStyles (optional) Whether the node styles in the tooltip (default: false).
	ShowStyles bool `json:"showStyles,omitempty"`

	// ShowRulers (optional) Whether the rulers should be shown (default: false).
	ShowRulers bool `json:"showRulers,omitempty"`

	// ShowAccessibilityInfo (optional) Whether the a11y info should be shown (default: true).
	ShowAccessibilityInfo bool `json:"showAccessibilityInfo,omitempty"`

	// ShowExtensionLines (optional) Whether the extension lines from node to the rulers should be shown (default: false).
	ShowExtensionLines bool `json:"showExtensionLines,omitempty"`

	// ContentColor (optional) The content box highlight fill color (default: transparent).
	ContentColor *DOMRGBA `json:"contentColor,omitempty"`

	// PaddingColor (optional) The padding highlight fill color (default: transparent).
	PaddingColor *DOMRGBA `json:"paddingColor,omitempty"`

	// BorderColor (optional) The border highlight fill color (default: transparent).
	BorderColor *DOMRGBA `json:"borderColor,omitempty"`

	// MarginColor (optional) The margin highlight fill color (default: transparent).
	MarginColor *DOMRGBA `json:"marginColor,omitempty"`

	// EventTargetColor (optional) The event target element highlight fill color (default: transparent).
	EventTargetColor *DOMRGBA `json:"eventTargetColor,omitempty"`

	// ShapeColor (optional) The shape outside fill color (default: transparent).
	ShapeColor *DOMRGBA `json:"shapeColor,omitempty"`

	// ShapeMarginColor (optional) The shape margin fill color (default: transparent).
	ShapeMarginColor *DOMRGBA `json:"shapeMarginColor,omitempty"`

	// CSSGridColor (optional) The grid layout color (default: transparent).
	CSSGridColor *DOMRGBA `json:"cssGridColor,omitempty"`

	// ColorFormat (optional) The color format used to format color styles (default: hex).
	ColorFormat OverlayColorFormat `json:"colorFormat,omitempty"`

	// GridHighlightConfig (optional) The grid layout highlight configuration (default: all transparent).
	GridHighlightConfig *OverlayGridHighlightConfig `json:"gridHighlightConfig,omitempty"`

	// FlexContainerHighlightConfig (optional) The flex container highlight configuration (default: all transparent).
	FlexContainerHighlightConfig *OverlayFlexContainerHighlightConfig `json:"flexContainerHighlightConfig,omitempty"`

	// FlexItemHighlightConfig (optional) The flex item highlight configuration (default: all transparent).
	FlexItemHighlightConfig *OverlayFlexItemHighlightConfig `json:"flexItemHighlightConfig,omitempty"`

	// ContrastAlgorithm (optional) The contrast algorithm to use for the contrast ratio (default: aa).
	ContrastAlgorithm OverlayContrastAlgorithm `json:"contrastAlgorithm,omitempty"`

	// ContainerQueryContainerHighlightConfig (optional) The container query container highlight configuration (default: all transparent).
	ContainerQueryContainerHighlightConfig *OverlayContainerQueryContainerHighlightConfig `json:"containerQueryContainerHighlightConfig,omitempty"`
}

// OverlayColorFormat ...
type OverlayColorFormat string

const (
	// OverlayColorFormatRgb enum const
	OverlayColorFormatRgb OverlayColorFormat = "rgb"

	// OverlayColorFormatHsl enum const
	OverlayColorFormatHsl OverlayColorFormat = "hsl"

	// OverlayColorFormatHwb enum const
	OverlayColorFormatHwb OverlayColorFormat = "hwb"

	// OverlayColorFormatHex enum const
	OverlayColorFormatHex OverlayColorFormat = "hex"
)

// OverlayGridNodeHighlightConfig Configurations for Persistent Grid Highlight
type OverlayGridNodeHighlightConfig struct {

	// GridHighlightConfig A descriptor for the highlight appearance.
	GridHighlightConfig *OverlayGridHighlightConfig `json:"gridHighlightConfig"`

	// NodeID Identifier of the node to highlight.
	NodeID DOMNodeID `json:"nodeId"`
}

// OverlayFlexNodeHighlightConfig ...
type OverlayFlexNodeHighlightConfig struct {

	// FlexContainerHighlightConfig A descriptor for the highlight appearance of flex containers.
	FlexContainerHighlightConfig *OverlayFlexContainerHighlightConfig `json:"flexContainerHighlightConfig"`

	// NodeID Identifier of the node to highlight.
	NodeID DOMNodeID `json:"nodeId"`
}

// OverlayScrollSnapContainerHighlightConfig ...
type OverlayScrollSnapContainerHighlightConfig struct {

	// SnapportBorder (optional) The style of the snapport border (default: transparent)
	SnapportBorder *OverlayLineStyle `json:"snapportBorder,omitempty"`

	// SnapAreaBorder (optional) The style of the snap area border (default: transparent)
	SnapAreaBorder *OverlayLineStyle `json:"snapAreaBorder,omitempty"`

	// ScrollMarginColor (optional) The margin highlight fill color (default: transparent).
	ScrollMarginColor *DOMRGBA `json:"scrollMarginColor,omitempty"`

	// ScrollPaddingColor (optional) The padding highlight fill color (default: transparent).
	ScrollPaddingColor *DOMRGBA `json:"scrollPaddingColor,omitempty"`
}

// OverlayScrollSnapHighlightConfig ...
type OverlayScrollSnapHighlightConfig struct {

	// ScrollSnapContainerHighlightConfig A descriptor for the highlight appearance of scroll snap containers.
	ScrollSnapContainerHighlightConfig *OverlayScrollSnapContainerHighlightConfig `json:"scrollSnapContainerHighlightConfig"`

	// NodeID Identifier of the node to highlight.
	NodeID DOMNodeID `json:"nodeId"`
}

// OverlayHingeConfig Configuration for dual screen hinge
type OverlayHingeConfig struct {

	// Rect A rectangle represent hinge
	Rect *DOMRect `json:"rect"`

	// ContentColor (optional) The content box highlight fill color (default: a dark color).
	ContentColor *DOMRGBA `json:"contentColor,omitempty"`

	// OutlineColor (optional) The content box highlight outline color (default: transparent).
	OutlineColor *DOMRGBA `json:"outlineColor,omitempty"`
}

// OverlayContainerQueryHighlightConfig ...
type OverlayContainerQueryHighlightConfig struct {

	// ContainerQueryContainerHighlightConfig A descriptor for the highlight appearance of container query containers.
	ContainerQueryContainerHighlightConfig *OverlayContainerQueryContainerHighlightConfig `json:"containerQueryContainerHighlightConfig"`

	// NodeID Identifier of the container node to highlight.
	NodeID DOMNodeID `json:"nodeId"`
}

// OverlayContainerQueryContainerHighlightConfig ...
type OverlayContainerQueryContainerHighlightConfig struct {

	// ContainerBorder (optional) The style of the container border.
	ContainerBorder *OverlayLineStyle `json:"containerBorder,omitempty"`

	// DescendantBorder (optional) The style of the descendants' borders.
	DescendantBorder *OverlayLineStyle `json:"descendantBorder,omitempty"`
}

// OverlayIsolatedElementHighlightConfig ...
type OverlayIsolatedElementHighlightConfig struct {

	// IsolationModeHighlightConfig A descriptor for the highlight appearance of an element in isolation mode.
	IsolationModeHighlightConfig *OverlayIsolationModeHighlightConfig `json:"isolationModeHighlightConfig"`

	// NodeID Identifier of the isolated element to highlight.
	NodeID DOMNodeID `json:"nodeId"`
}

// OverlayIsolationModeHighlightConfig ...
type OverlayIsolationModeHighlightConfig struct {

	// ResizerColor (optional) The fill color of the resizers (default: transparent).
	ResizerColor *DOMRGBA `json:"resizerColor,omitempty"`

	// ResizerHandleColor (optional) The fill color for resizer handles (default: transparent).
	ResizerHandleColor *DOMRGBA `json:"resizerHandleColor,omitempty"`

	// MaskColor (optional) The fill color for the mask covering non-isolated elements (default: transparent).
	MaskColor *DOMRGBA `json:"maskColor,omitempty"`
}

// OverlayInspectMode ...
type OverlayInspectMode string

const (
	// OverlayInspectModeSearchForNode enum const
	OverlayInspectModeSearchForNode OverlayInspectMode = "searchForNode"

	// OverlayInspectModeSearchForUAShadowDOM enum const
	OverlayInspectModeSearchForUAShadowDOM OverlayInspectMode = "searchForUAShadowDOM"

	// OverlayInspectModeCaptureAreaScreenshot enum const
	OverlayInspectModeCaptureAreaScreenshot OverlayInspectMode = "captureAreaScreenshot"

	// OverlayInspectModeShowDistances enum const
	OverlayInspectModeShowDistances OverlayInspectMode = "showDistances"

	// OverlayInspectModeNone enum const
	OverlayInspectModeNone OverlayInspectMode = "none"
)

// OverlayDisable Disables domain notifications.
type OverlayDisable struct {
}

// ProtoReq name
func (m OverlayDisable) ProtoReq() string { return "Overlay.disable" }

// Call sends the request
func (m OverlayDisable) Call(c Client) error {
	return call(m.ProtoReq(), m, nil, c)
}

// OverlayEnable Enables domain notifications.
type OverlayEnable struct {
}

// ProtoReq name
func (m OverlayEnable) ProtoReq() string { return "Overlay.enable" }

// Call sends the request
func (m OverlayEnable) Call(c Client) error {
	return call(m.ProtoReq(), m, nil, c)
}

// OverlayGetHighlightObjectForTest For testing.
type OverlayGetHighlightObjectForTest struct {

	// NodeID Id of the node to get highlight object for.
	NodeID DOMNodeID `json:"nodeId"`

	// IncludeDistance (optional) Whether to include distance info.
	IncludeDistance bool `json:"includeDistance,omitempty"`

	// IncludeStyle (optional) Whether to include style info.
	IncludeStyle bool `json:"includeStyle,omitempty"`

	// ColorFormat (optional) The color format to get config with (default: hex).
	ColorFormat OverlayColorFormat `json:"colorFormat,omitempty"`

	// ShowAccessibilityInfo (optional) Whether to show accessibility info (default: true).
	ShowAccessibilityInfo bool `json:"showAccessibilityInfo,omitempty"`
}

// ProtoReq name
func (m OverlayGetHighlightObjectForTest) ProtoReq() string {
	return "Overlay.getHighlightObjectForTest"
}

// Call the request
func (m OverlayGetHighlightObjectForTest) Call(c Client) (*OverlayGetHighlightObjectForTestResult, error) {
	var res OverlayGetHighlightObjectForTestResult
	return &res, call(m.ProtoReq(), m, &res, c)
}

// OverlayGetHighlightObjectForTestResult ...
type OverlayGetHighlightObjectForTestResult struct {

	// Highlight Highlight data for the node.
	Highlight map[string]gson.JSON `json:"highlight"`
}

// OverlayGetGridHighlightObjectsForTest For Persistent Grid testing.
type OverlayGetGridHighlightObjectsForTest struct {

	// NodeIds Ids of the node to get highlight object for.
	NodeIds []DOMNodeID `json:"nodeIds"`
}

// ProtoReq name
func (m OverlayGetGridHighlightObjectsForTest) ProtoReq() string {
	return "Overlay.getGridHighlightObjectsForTest"
}

// Call the request
func (m OverlayGetGridHighlightObjectsForTest) Call(c Client) (*OverlayGetGridHighlightObjectsForTestResult, error) {
	var res OverlayGetGridHighlightObjectsForTestResult
	return &res, call(m.ProtoReq(), m, &res, c)
}

// OverlayGetGridHighlightObjectsForTestResult ...
type OverlayGetGridHighlightObjectsForTestResult struct {

	// Highlights Grid Highlight data for the node ids provided.
	Highlights map[string]gson.JSON `json:"highlights"`
}

// OverlayGetSourceOrderHighlightObjectForTest For Source Order Viewer testing.
type OverlayGetSourceOrderHighlightObjectForTest struct {

	// NodeID Id of the node to highlight.
	NodeID DOMNodeID `json:"nodeId"`
}

// ProtoReq name
func (m OverlayGetSourceOrderHighlightObjectForTest) ProtoReq() string {
	return "Overlay.getSourceOrderHighlightObjectForTest"
}

// Call the request
func (m OverlayGetSourceOrderHighlightObjectForTest) Call(c Client) (*OverlayGetSourceOrderHighlightObjectForTestResult, error) {
	var res OverlayGetSourceOrderHighlightObjectForTestResult
	return &res, call(m.ProtoReq(), m, &res, c)
}

// OverlayGetSourceOrderHighlightObjectForTestResult ...
type OverlayGetSourceOrderHighlightObjectForTestResult struct {

	// Highlight Source order highlight data for the node id provided.
	Highlight map[string]gson.JSON `json:"highlight"`
}

// OverlayHideHighlight Hides any highlight.
type OverlayHideHighlight struct {
}

// ProtoReq name
func (m OverlayHideHighlight) ProtoReq() string { return "Overlay.hideHighlight" }

// Call sends the request
func (m OverlayHideHighlight) Call(c Client) error {
	return call(m.ProtoReq(), m, nil, c)
}

// OverlayHighlightFrame (deprecated) Highlights owner element of the frame with given id.
// Deprecated: Doesn't work reliablity and cannot be fixed due to process
// separatation (the owner node might be in a different process). Determine
// the owner node in the client and use highlightNode.
type OverlayHighlightFrame struct {

	// FrameID Identifier of the frame to highlight.
	FrameID PageFrameID `json:"frameId"`

	// ContentColor (optional) The content box highlight fill color (default: transparent).
	ContentColor *DOMRGBA `json:"contentColor,omitempty"`

	// ContentOutlineColor (optional) The content box highlight outline color (default: transparent).
	ContentOutlineColor *DOMRGBA `json:"contentOutlineColor,omitempty"`
}

// ProtoReq name
func (m OverlayHighlightFrame) ProtoReq() string { return "Overlay.highlightFrame" }

// Call sends the request
func (m OverlayHighlightFrame) Call(c Client) error {
	return call(m.ProtoReq(), m, nil, c)
}

// OverlayHighlightNode Highlights DOM node with given id or with the given JavaScript object wrapper. Either nodeId or
// objectId must be specified.
type OverlayHighlightNode struct {

	// HighlightConfig A descriptor for the highlight appearance.
	HighlightConfig *OverlayHighlightConfig `json:"highlightConfig"`

	// NodeID (optional) Identifier of the node to highlight.
	NodeID DOMNodeID `json:"nodeId,omitempty"`

	// BackendNodeID (optional) Identifier of the backend node to highlight.
	BackendNodeID DOMBackendNodeID `json:"backendNodeId,omitempty"`

	// ObjectID (optional) JavaScript object id of the node to be highlighted.
	ObjectID RuntimeRemoteObjectID `json:"objectId,omitempty"`

	// Selector (optional) Selectors to highlight relevant nodes.
	Selector string `json:"selector,omitempty"`
}

// ProtoReq name
func (m OverlayHighlightNode) ProtoReq() string { return "Overlay.highlightNode" }

// Call sends the request
func (m OverlayHighlightNode) Call(c Client) error {
	return call(m.ProtoReq(), m, nil, c)
}

// OverlayHighlightQuad Highlights given quad. Coordinates are absolute with respect to the main frame viewport.
type OverlayHighlightQuad struct {

	// Quad Quad to highlight
	Quad DOMQuad `json:"quad"`

	// Color (optional) The highlight fill color (default: transparent).
	Color *DOMRGBA `json:"color,omitempty"`

	// OutlineColor (optional) The highlight outline color (default: transparent).
	OutlineColor *DOMRGBA `json:"outlineColor,omitempty"`
}

// ProtoReq name
func (m OverlayHighlightQuad) ProtoReq() string { return "Overlay.highlightQuad" }

// Call sends the request
func (m OverlayHighlightQuad) Call(c Client) error {
	return call(m.ProtoReq(), m, nil, c)
}

// OverlayHighlightRect Highlights given rectangle. Coordinates are absolute with respect to the main frame viewport.
type OverlayHighlightRect struct {

	// X X coordinate
	X int `json:"x"`

	// Y Y coordinate
	Y int `json:"y"`

	// Width Rectangle width
	Width int `json:"width"`

	// Height Rectangle height
	Height int `json:"height"`

	// Color (optional) The highlight fill color (default: transparent).
	Color *DOMRGBA `json:"color,omitempty"`

	// OutlineColor (optional) The highlight outline color (default: transparent).
	OutlineColor *DOMRGBA `json:"outlineColor,omitempty"`
}

// ProtoReq name
func (m OverlayHighlightRect) ProtoReq() string { return "Overlay.highlightRect" }

// Call sends the request
func (m OverlayHighlightRect) Call(c Client) error {
	return call(m.ProtoReq(), m, nil, c)
}

// OverlayHighlightSourceOrder Highlights the source order of the children of the DOM node with given id or with the given
// JavaScript object wrapper. Either nodeId or objectId must be specified.
type OverlayHighlightSourceOrder struct {

	// SourceOrderConfig A descriptor for the appearance of the overlay drawing.
	SourceOrderConfig *OverlaySourceOrderConfig `json:"sourceOrderConfig"`

	// NodeID (optional) Identifier of the node to highlight.
	NodeID DOMNodeID `json:"nodeId,omitempty"`

	// BackendNodeID (optional) Identifier of the backend node to highlight.
	BackendNodeID DOMBackendNodeID `json:"backendNodeId,omitempty"`

	// ObjectID (optional) JavaScript object id of the node to be highlighted.
	ObjectID RuntimeRemoteObjectID `json:"objectId,omitempty"`
}

// ProtoReq name
func (m OverlayHighlightSourceOrder) ProtoReq() string { return "Overlay.highlightSourceOrder" }

// Call sends the request
func (m OverlayHighlightSourceOrder) Call(c Client) error {
	return call(m.ProtoReq(), m, nil, c)
}

// OverlaySetInspectMode Enters the 'inspect' mode. In this mode, elements that user is hovering over are highlighted.
// Backend then generates 'inspectNodeRequested' event upon element selection.
type OverlaySetInspectMode struct {

	// Mode Set an inspection mode.
	Mode OverlayInspectMode `json:"mode"`

	// HighlightConfig (optional) A descriptor for the highlight appearance of hovered-over nodes. May be omitted if `enabled
	// == false`.
	HighlightConfig *OverlayHighlightConfig `json:"highlightConfig,omitempty"`
}

// ProtoReq name
func (m OverlaySetInspectMode) ProtoReq() string { return "Overlay.setInspectMode" }

// Call sends the request
func (m OverlaySetInspectMode) Call(c Client) error {
	return call(m.ProtoReq(), m, nil, c)
}

// OverlaySetShowAdHighlights Highlights owner element of all frames detected to be ads.
type OverlaySetShowAdHighlights struct {

	// Show True for showing ad highlights
	Show bool `json:"show"`
}

// ProtoReq name
func (m OverlaySetShowAdHighlights) ProtoReq() string { return "Overlay.setShowAdHighlights" }

// Call sends the request
func (m OverlaySetShowAdHighlights) Call(c Client) error {
	return call(m.ProtoReq(), m, nil, c)
}

// OverlaySetPausedInDebuggerMessage ...
type OverlaySetPausedInDebuggerMessage struct {

	// Message (optional) The message to display, also triggers resume and step over controls.
	Message string `json:"message,omitempty"`
}

// ProtoReq name
func (m OverlaySetPausedInDebuggerMessage) ProtoReq() string {
	return "Overlay.setPausedInDebuggerMessage"
}

// Call sends the request
func (m OverlaySetPausedInDebuggerMessage) Call(c Client) error {
	return call(m.ProtoReq(), m, nil, c)
}

// OverlaySetShowDebugBorders Requests that backend shows debug borders on layers
type OverlaySetShowDebugBorders struct {

	// Show True for showing debug borders
	Show bool `json:"show"`
}

// ProtoReq name
func (m OverlaySetShowDebugBorders) ProtoReq() string { return "Overlay.setShowDebugBorders" }

// Call sends the request
func (m OverlaySetShowDebugBorders) Call(c Client) error {
	return call(m.ProtoReq(), m, nil, c)
}

// OverlaySetShowFPSCounter Requests that backend shows the FPS counter
type OverlaySetShowFPSCounter struct {

	// Show True for showing the FPS counter
	Show bool `json:"show"`
}

// ProtoReq name
func (m OverlaySetShowFPSCounter) ProtoReq() string { return "Overlay.setShowFPSCounter" }

// Call sends the request
func (m OverlaySetShowFPSCounter) Call(c Client) error {
	return call(m.ProtoReq(), m, nil, c)
}

// OverlaySetShowGridOverlays Highlight multiple elements with the CSS Grid overlay.
type OverlaySetShowGridOverlays struct {

	// GridNodeHighlightConfigs An array of node identifiers and descriptors for the highlight appearance.
	GridNodeHighlightConfigs []*OverlayGridNodeHighlightConfig `json:"gridNodeHighlightConfigs"`
}

// ProtoReq name
func (m OverlaySetShowGridOverlays) ProtoReq() string { return "Overlay.setShowGridOverlays" }

// Call sends the request
func (m OverlaySetShowGridOverlays) Call(c Client) error {
	return call(m.ProtoReq(), m, nil, c)
}

// OverlaySetShowFlexOverlays ...
type OverlaySetShowFlexOverlays struct {

	// FlexNodeHighlightConfigs An array of node identifiers and descriptors for the highlight appearance.
	FlexNodeHighlightConfigs []*OverlayFlexNodeHighlightConfig `json:"flexNodeHighlightConfigs"`
}

// ProtoReq name
func (m OverlaySetShowFlexOverlays) ProtoReq() string { return "Overlay.setShowFlexOverlays" }

// Call sends the request
func (m OverlaySetShowFlexOverlays) Call(c Client) error {
	return call(m.ProtoReq(), m, nil, c)
}

// OverlaySetShowScrollSnapOverlays ...
type OverlaySetShowScrollSnapOverlays struct {

	// ScrollSnapHighlightConfigs An array of node identifiers and descriptors for the highlight appearance.
	ScrollSnapHighlightConfigs []*OverlayScrollSnapHighlightConfig `json:"scrollSnapHighlightConfigs"`
}

// ProtoReq name
func (m OverlaySetShowScrollSnapOverlays) ProtoReq() string {
	return "Overlay.setShowScrollSnapOverlays"
}

// Call sends the request
func (m OverlaySetShowScrollSnapOverlays) Call(c Client) error {
	return call(m.ProtoReq(), m, nil, c)
}

// OverlaySetShowContainerQueryOverlays ...
type OverlaySetShowContainerQueryOverlays struct {

	// ContainerQueryHighlightConfigs An array of node identifiers and descriptors for the highlight appearance.
	ContainerQueryHighlightConfigs []*OverlayContainerQueryHighlightConfig `json:"containerQueryHighlightConfigs"`
}

// ProtoReq name
func (m OverlaySetShowContainerQueryOverlays) ProtoReq() string {
	return "Overlay.setShowContainerQueryOverlays"
}

// Call sends the request
func (m OverlaySetShowContainerQueryOverlays) Call(c Client) error {
	return call(m.ProtoReq(), m, nil, c)
}

// OverlaySetShowPaintRects Requests that backend shows paint rectangles
type OverlaySetShowPaintRects struct {

	// Result True for showing paint rectangles
	Result bool `json:"result"`
}

// ProtoReq name
func (m OverlaySetShowPaintRects) ProtoReq() string { return "Overlay.setShowPaintRects" }

// Call sends the request
func (m OverlaySetShowPaintRects) Call(c Client) error {
	return call(m.ProtoReq(), m, nil, c)
}

// OverlaySetShowLayoutShiftRegions Requests that backend shows layout shift regions
type OverlaySetShowLayoutShiftRegions struct {

	// Result True for showing layout shift regions
	Result bool `json:"result"`
}

// ProtoReq name
func (m OverlaySetShowLayoutShiftRegions) ProtoReq() string {
	return "Overlay.setShowLayoutShiftRegions"
}

// Call sends the request
func (m OverlaySetShowLayoutShiftRegions) Call(c Client) error {
	return call(m.ProtoReq(), m, nil, c)
}

// OverlaySetShowScrollBottleneckRects Requests that backend shows scroll bottleneck rects
type OverlaySetShowScrollBottleneckRects struct {

	// Show True for showing scroll bottleneck rects
	Show bool `json:"show"`
}

// ProtoReq name
func (m OverlaySetShowScrollBottleneckRects) ProtoReq() string {
	return "Overlay.setShowScrollBottleneckRects"
}

// Call sends the request
func (m OverlaySetShowScrollBottleneckRects) Call(c Client) error {
	return call(m.ProtoReq(), m, nil, c)
}

// OverlaySetShowHitTestBorders (deprecated) Deprecated, no longer has any effect.
type OverlaySetShowHitTestBorders struct {

	// Show True for showing hit-test borders
	Show bool `json:"show"`
}

// ProtoReq name
func (m OverlaySetShowHitTestBorders) ProtoReq() string { return "Overlay.setShowHitTestBorders" }

// Call sends the request
func (m OverlaySetShowHitTestBorders) Call(c Client) error {
	return call(m.ProtoReq(), m, nil, c)
}

// OverlaySetShowWebVitals Request that backend shows an overlay with web vital metrics.
type OverlaySetShowWebVitals struct {

	// Show ...
	Show bool `json:"show"`
}

// ProtoReq name
func (m OverlaySetShowWebVitals) ProtoReq() string { return "Overlay.setShowWebVitals" }

// Call sends the request
func (m OverlaySetShowWebVitals) Call(c Client) error {
	return call(m.ProtoReq(), m, nil, c)
}

// OverlaySetShowViewportSizeOnResize Paints viewport size upon main frame resize.
type OverlaySetShowViewportSizeOnResize struct {

	// Show Whether to paint size or not.
	Show bool `json:"show"`
}

// ProtoReq name
func (m OverlaySetShowViewportSizeOnResize) ProtoReq() string {
	return "Overlay.setShowViewportSizeOnResize"
}

// Call sends the request
func (m OverlaySetShowViewportSizeOnResize) Call(c Client) error {
	return call(m.ProtoReq(), m, nil, c)
}

// OverlaySetShowHinge Add a dual screen device hinge
type OverlaySetShowHinge struct {

	// HingeConfig (optional) hinge data, null means hideHinge
	HingeConfig *OverlayHingeConfig `json:"hingeConfig,omitempty"`
}

// ProtoReq name
func (m OverlaySetShowHinge) ProtoReq() string { return "Overlay.setShowHinge" }

// Call sends the request
func (m OverlaySetShowHinge) Call(c Client) error {
	return call(m.ProtoReq(), m, nil, c)
}

// OverlaySetShowIsolatedElements Show elements in isolation mode with overlays.
type OverlaySetShowIsolatedElements struct {

	// IsolatedElementHighlightConfigs An array of node identifiers and descriptors for the highlight appearance.
	IsolatedElementHighlightConfigs []*OverlayIsolatedElementHighlightConfig `json:"isolatedElementHighlightConfigs"`
}

// ProtoReq name
func (m OverlaySetShowIsolatedElements) ProtoReq() string { return "Overlay.setShowIsolatedElements" }

// Call sends the request
func (m OverlaySetShowIsolatedElements) Call(c Client) error {
	return call(m.ProtoReq(), m, nil, c)
}

// OverlayInspectNodeRequested Fired when the node should be inspected. This happens after call to `setInspectMode` or when
// user manually inspects an element.
type OverlayInspectNodeRequested struct {

	// BackendNodeID Id of the node to inspect.
	BackendNodeID DOMBackendNodeID `json:"backendNodeId"`
}

// ProtoEvent name
func (evt OverlayInspectNodeRequested) ProtoEvent() string {
	return "Overlay.inspectNodeRequested"
}

// OverlayNodeHighlightRequested Fired when the node should be highlighted. This happens after call to `setInspectMode`.
type OverlayNodeHighlightRequested struct {

	// NodeID ...
	NodeID DOMNodeID `json:"nodeId"`
}

// ProtoEvent name
func (evt OverlayNodeHighlightRequested) ProtoEvent() string {
	return "Overlay.nodeHighlightRequested"
}

// OverlayScreenshotRequested Fired when user asks to capture screenshot of some area on the page.
type OverlayScreenshotRequested struct {

	// Viewport Viewport to capture, in device independent pixels (dip).
	Viewport *PageViewport `json:"viewport"`
}

// ProtoEvent name
func (evt OverlayScreenshotRequested) ProtoEvent() string {
	return "Overlay.screenshotRequested"
}

// OverlayInspectModeCanceled Fired when user cancels the inspect mode.
type OverlayInspectModeCanceled struct {
}

// ProtoEvent name
func (evt OverlayInspectModeCanceled) ProtoEvent() string {
	return "Overlay.inspectModeCanceled"
}
