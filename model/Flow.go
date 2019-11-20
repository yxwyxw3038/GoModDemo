package model

type FlowBillModel struct {
	Flow         FlowView
	FlowStep     []FlowStep
	FlowStepUser []FlowStepUserView
}
