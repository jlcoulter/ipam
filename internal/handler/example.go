package handler

import (
	//"encoding/json"
	"net/http"
	//"github.com/go-chi/chi/v5"
	"github.com/jlcoulter/ipam/internal/ipam"
)

type API struct {
	IPAM   *ipam.IPAM
	Router *http.ServeMux
}

func NewAPI(i *ipam.IPAM) *API {
	api := &API{IPAM: i, Router: http.NewServeMux()}
	// TODO: Register routes
	// POST /api/allocate
	// DELETE /api/hosts/{ip}
	// GET /api/hosts?subnet={cidr}
	// GET /api/hosts/{ip}
	// GET /api/subnets
	// POST /api/subnets
	// GET /api/subnets/{cidr}
	// POST /api/scan/{cidr}
	// POST /api/import
	return api
}

func (a *API) handleAllocate(w http.ResponseWriter, r *http.Request) {
	// TODO: Parse JSON body with subnet host and optional ip/mac/role/desc/tags
	// TODO: Call ipam.Allocate
	// TODO: Return JSON resp
}

func (a *API) handleRelease(w http.ResponseWriter, r *http.Request) {
	// TODO: Extract IP from url path
	// TODO: Call ipam.Release
	// TODO: Return JSON resp
}

func (a *API) handleList(w http.ResponseWriter, r *http.Request) {
	// TODO: Extract subnet from query param
	// TODO: Call ipam.List
	// TODO: Return JSON resp
}

func (a *API) handleShow(w http.ResponseWriter, r *http.Request) {
	// TODO: Extract IP from URL path
	// TODO: Call ipam.Show
	// TODO: Return JSON resp
}

func (a *API) handleSubnetList(w http.ResponseWriter, r *http.Request) {
	// TODO: Call ipam.Subnets
	// TODO: Return JSON resp
}

func (a *API) handleSubnetCreate(w http.ResponseWriter, r *http.Request) {
	// TODO: Parse JSON body with cidr, name, desc, vlan
	// TODO: Call ipam.AddSubnet
	// TODO: Return JSON resp
}

func (a *API) handleSubnet(w http.ResponseWriter, r *http.Request) {
	// TODO: Extract cidr from url path
	// TODO: calculate and return subnet summary
}

func (a *API) handleScan(w http.ResponseWriter, r *http.Request) {
	// TODO: Extract cidr from url path
	// TODO: Run ipam.ScanSubnet
	// TODO: Upsert discovered hosts
	// TODO: Return JSON resp with discovered hosts
}

func (a *API) handleImport(w http.ResponseWriter, r *http.Request) {
	// TODO: Parse uploaded nmap XML from multipart form
	// TODO: Run ipam.ImportNmap
	// TODO: Upsert discovered hosts
	// TODO: Return JSON resp with import summary
}

