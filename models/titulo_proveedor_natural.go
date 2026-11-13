package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/beego/beego/v2/client/orm"
)

type TituloProveedorNatural struct {
	Id                              int        `orm:"column(id);pk;auto"`
	TituloObtenido                  string     `orm:"column(titulo_obtenido)"`
	NucleoBasicoConocimientoSniesId int        `orm:"column(nucleo_basico_conocimiento_snies_id)"`
	ProveedorId                     *Proveedor `orm:"column(proveedor_id);rel(fk)"`
	PerfilId                        int        `orm:"column(perfil_id)"`
	Activo                          bool       `orm:"column(activo)"`
	FechaCreacion                   time.Time  `orm:"column(fecha_creacion);type(timestamp without time zone)"`
	FechaModificacion               time.Time  `orm:"column(fecha_modificacion);type(timestamp without time zone);null"`
	UsuarioModificacionId           int        `orm:"column(usuario_modificacion_id);null"`
}

func (t *TituloProveedorNatural) TableName() string {
	return "titulo_proveedor_natural"
}

func init() {
	orm.RegisterModel(new(TituloProveedorNatural))
}

// AddTituloProveedorNatural insert a new TituloProveedorNatural into database and returns
// last inserted Id on success.
func AddTituloProveedorNatural(m *TituloProveedorNatural) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetTituloProveedorNaturalById retrieves TituloProveedorNatural by Id. Returns error if
// Id doesn't exist
func GetTituloProveedorNaturalById(id int) (v *TituloProveedorNatural, err error) {
	o := orm.NewOrm()
	v = &TituloProveedorNatural{}
	// RelatedSel expande las FK, igual que en GetAllTituloProveedorNatural
	if err = o.QueryTable(new(TituloProveedorNatural)).Filter("Id", id).RelatedSel().One(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllTituloProveedorNatural retrieves all TituloProveedorNatural matches certain condition. Returns empty list if
// no records exist
func GetAllTituloProveedorNatural(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(TituloProveedorNatural)).RelatedSel()
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		if strings.Contains(k, "isnull") {
			qs = qs.Filter(k, (v == "true" || v == "1"))
		} else {
			qs = qs.Filter(k, v)
		}
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []TituloProveedorNatural
	qs = qs.OrderBy(sortFields...)
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// UpdateTituloProveedorNatural updates TituloProveedorNatural by Id and returns error if
// the record to be updated doesn't exist
func UpdateTituloProveedorNaturalById(m *TituloProveedorNatural) (err error) {
	o := orm.NewOrm()
	v := TituloProveedorNatural{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m, "TituloObtenido", "NucleoBasicoConocimientoSniesId", "ProveedorId", "PerfilId",
			"Activo", "FechaModificacion", "UsuarioModificacionId"); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteTituloProveedorNatural deletes TituloProveedorNatural by Id and returns error if
// the record to be deleted doesn't exist
func DeleteTituloProveedorNatural(id int) (err error) {
	o := orm.NewOrm()
	v := TituloProveedorNatural{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&TituloProveedorNatural{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
