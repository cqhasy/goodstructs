package main

import (
	"fmt"
)

// 请求结构
type ReimbursementRequest struct {
	Amount      float64
	Description string
}

// Handler 接口
type Handler interface {
	SetNext(Handler)
	Handle(*ReimbursementRequest)
}

// 抽象基类：组合而非内嵌
type BaseHandlerImpl struct {
	next Handler
}

// 实现接口方法
func (h *BaseHandlerImpl) SetNext(next Handler) {
	h.next = next
}

func (h *BaseHandlerImpl) CallNext(request *ReimbursementRequest) {
	if h.next != nil {
		h.next.Handle(request)
	}
}

// ManagerHandler：经理处理
type ManagerHandler struct {
	base *BaseHandlerImpl
}

func NewManagerHandler() *ManagerHandler {
	return &ManagerHandler{base: &BaseHandlerImpl{}}
}

func (h *ManagerHandler) SetNext(next Handler) {
	h.base.SetNext(next)
}

func (h *ManagerHandler) Handle(request *ReimbursementRequest) {
	if request.Amount <= 1000 {
		fmt.Printf("经理审批通过：%s（金额：%.2f）\n", request.Description, request.Amount)
	} else {
		h.base.CallNext(request)
	}
}

// DepartmentHeadHandler：部门主管处理
type DepartmentHeadHandler struct {
	base *BaseHandlerImpl
}

func NewDepartmentHeadHandler() *DepartmentHeadHandler {
	return &DepartmentHeadHandler{base: &BaseHandlerImpl{}}
}

func (h *DepartmentHeadHandler) SetNext(next Handler) {
	h.base.SetNext(next)
}

func (h *DepartmentHeadHandler) Handle(request *ReimbursementRequest) {
	if request.Amount <= 5000 {
		fmt.Printf("部门主管审批通过：%s（金额：%.2f）\n", request.Description, request.Amount)
	} else {
		h.base.CallNext(request)
	}
}

// FinanceHandler：财务处理
type FinanceHandler struct {
	base *BaseHandlerImpl
}

func NewFinanceHandler() *FinanceHandler {
	return &FinanceHandler{base: &BaseHandlerImpl{}}
}

func (h *FinanceHandler) SetNext(next Handler) {
	h.base.SetNext(next)
}

func (h *FinanceHandler) Handle(request *ReimbursementRequest) {
	fmt.Printf("财务部门审批通过：%s（金额：%.2f）\n", request.Description, request.Amount)
}

// 入口
func main() {
	manager := NewManagerHandler()
	deptHead := NewDepartmentHeadHandler()
	finance := NewFinanceHandler()

	manager.SetNext(deptHead)
	deptHead.SetNext(finance)

	requests := []*ReimbursementRequest{
		{Amount: 800, Description: "购买办公用品"},
		{Amount: 3000, Description: "参加培训"},
		{Amount: 10000, Description: "举办团建活动"},
	}

	for _, req := range requests {
		manager.Handle(req)
	}
}
