package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name  string
	Age   int
	Email string
}

func main() {
	p := Person{Name: "John", Age: 30, Email: "john@example.com"}

	b, err := json.Marshal(p)

	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(string(b))
	fmt.Println(string(b))
}


// Reconcile is the main reconciliation loop for the controller
func (r *MyResourceReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
    log := r.Log.WithValues("myresource", req.NamespacedName)

    // Get the MyResource object
    myresource := &mygroupv1.MyResource{}
    if err := r.Get(ctx, req.NamespacedName, myresource); err != nil {
        if errors.IsNotFound(err) {
            // Object not found, possibly deleted
            return ctrl.Result{}, nil
        }
        log.Error(err, "Failed to get MyResource")
        return ctrl.Result{}, err
    }

    // Reconcile logic here

    return ctrl.Result{}, nil
}

package controllers

import (
	"context"

	mygroupv1 "github.com/example/myproject/api/v1"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

// MyResourceReconciler reconciles MyResource objects
type MyResourceReconciler struct {
	client.Client
	Log    *ctrl.Log
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=mygroup.example.com,resources=myresources,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=mygroup.example.com,resources=myresources/status,verbs=get;update;patch

// Reconcile handles the reconciliation logic for MyResource objects
func (r *MyResourceReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	log := r.Log.WithValues("myresource", req.NamespacedName)

	// Fetch the MyResource object
	myresource := &mygroupv1.MyResource{}
	if err := r.Get(ctx, req.NamespacedName, myresource); err != nil {
		return reconcile.Result{}, client.IgnoreNotFound(err)
	}

	// Add your reconciliation logic here

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager
func (r *MyResourceReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&mygroupv1.MyResource{}).
		Complete(r)
}

// +kubebuilder:webhook:path=/mutate-mygroup-example-com-v1-myresource,mutating=true,failurePolicy=fail,groups=mygroup.example.com,resources=myresources,verbs=create;update,versions=v1,name=mmyresource.kb.io

var _ admission.Handler = &MyResourceReconciler{}

// Handle handles admission requests
func (r *MyResourceReconciler) Handle(ctx context.Context, req admission.Request) admission.Response {
	// Add your admission webhook logic here
}

// InjectClient injects the client into the MyResourceReconciler
func (r *MyResourceReconciler) InjectClient(c client.Client) error {
	r.Client = c
	return nil
}

===

func (r *MyResourceReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
    // ...
    
    // 예시: 새로운 리소스 생성
    newResource := &mygroupv1.MyResource{
        ObjectMeta: metav1.ObjectMeta{
            Name:      "new-resource",
            Namespace: req.Namespace,
        },
        Spec: mygroupv1.MyResourceSpec{
            // 스펙 설정
        },
    }

    // 리소스 유효성 검사 및 변환
    if err := ctrl.SetControllerReference(owner, newResource, r.Scheme); err != nil {
        return ctrl.Result{}, err
    }

    // 리소스 생성
    if err := r.Create(ctx, newResource); err != nil {
        return ctrl.Result{}, err
    }

    // ...
    
    return ctrl.Result{}, nil
}