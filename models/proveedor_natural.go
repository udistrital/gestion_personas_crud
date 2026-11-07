package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/beego/beego/v2/client/orm"
)

type ProveedorNatural struct {
	Id                                int       `orm:"column(id);pk;auto"`
	PaisNacimientoId                  int       `orm:"column(pais_nacimiento_id)"`
	Cargo                             string    `orm:"column(cargo);null"`
	MedicinaPrepagada                 bool      `orm:"column(medicina_prepagada);null"`
	ValorUvtPrepagada                 float64   `orm:"column(valor_uvt_prepagada);null"`
	InteresViviendaAfc                float64   `orm:"column(interes_vivienda_afc);null"`
	DependienteHijoMejorEdad          bool      `orm:"column(dependiente_hijo_mejor_edad);null"`
	DependienteHijoMenos23Estudiando  bool      `orm:"column(dependiente_hijo_menos23_estudiando);null"`
	DependienteHijoMas23Discapacitado bool      `orm:"column(dependiente_hijo_mas23_discapacitado);null"`
	DependienteConyuge                bool      `orm:"column(dependiente_conyuge);null"`
	DependientePadreOHermano          bool      `orm:"column(dependiente_padre_o_hermano);null"`
	Pensionado                        bool      `orm:"column(pensionado);null"`
	ValorAfc                          float64   `orm:"column(valor_afc);null"`
	ResponsableIva                    bool      `orm:"column(responsable_iva);null"`
	ResideColombia                    bool      `orm:"column(reside_colombia);null"`
	PensionVoluntaria                 bool      `orm:"column(pension_voluntaria);null"`
	ValorPensionVoluntaria            float64   `orm:"column(valor_pension_voluntaria);null"`
	ExperienciaLaboral                int       `orm:"column(experiencia_laboral);null"`
	ExperienciaProfesional            int       `orm:"column(experiencia_profesional);null"`
	Activo                            bool      `orm:"column(activo)"`
	FechaCreacion                     time.Time `orm:"column(fecha_creacion);type(timestamp without time zone)"`
	FechaModificacion                 time.Time `orm:"column(fecha_modificacion);type(timestamp without time zone);null"`
	UsuarioModificacionId             int       `orm:"column(usuario_modificacion_id);null"`
}

func (t *ProveedorNatural) TableName() string {
	return "proveedor_natural"
}

func init() {
	orm.RegisterModel(new(ProveedorNatural))
}

// AddProveedorNatural insert a new ProveedorNatural into database and returns
// last inserted Id on success.
func AddProveedorNatural(m *ProveedorNatural) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetProveedorNaturalById retrieves ProveedorNatural by Id. Returns error if
// Id doesn't exist
func GetProveedorNaturalById(id int) (v *ProveedorNatural, err error) {
	o := orm.NewOrm()
	v = &ProveedorNatural{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllProveedorNatural retrieves all ProveedorNatural matches certain condition. Returns empty list if
// no records exist
func GetAllProveedorNatural(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(ProveedorNatural)).RelatedSel()
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

	var l []ProveedorNatural
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

// UpdateProveedorNatural updates ProveedorNatural by Id and returns error if
// the record to be updated doesn't exist
func UpdateProveedorNaturalById(m *ProveedorNatural) (err error) {
	o := orm.NewOrm()
	v := ProveedorNatural{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteProveedorNatural deletes ProveedorNatural by Id and returns error if
// the record to be deleted doesn't exist
func DeleteProveedorNatural(id int) (err error) {
	o := orm.NewOrm()
	v := ProveedorNatural{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&ProveedorNatural{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
