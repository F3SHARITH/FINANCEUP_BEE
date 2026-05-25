import re
from pathlib import Path


def camel(s):
    return ''.join(part.capitalize() for part in s.split('_'))

root = Path(r"C:\Users\Alejandro Zorro\OneDrive\Escritorio\Repositorios_3237831\FINANCEUP_BEE")
model_dir = root / 'finanzas' / 'models'
wrapper_files = [
    'categoria.go', 'tipo_ingreso.go', 'nivel_riesgo.go', 'movimiento_inversion.go',
    'movimiento_meta.go', 'tipo_ingreso_inversion.go', 'finanzas.go',
    'tipo_ingreso_meta.go', 'editar_meta.go'
]

for name in wrapper_files:
    file = model_dir / name
    text = file.read_text(encoding='utf-8')
    model = camel(file.stem)
    new_import = 'import (\n\t"errors"\n\t"fmt"\n\t"reflect"\n\t"strings"\n\t"time"\n\n\t"github.com/beego/beego/v2/client/orm"\n)\n\n'
    text = re.sub(r'import \(.*?\n\)', new_import, text, flags=re.S)
    before = re.search(r'func Add' + re.escape(model) + r'\(', text)
    if not before:
        raise SystemExit(f"Could not find Add{model} in {name}")
    new_body_template = '''
func Add{model}(m *{model}) (id int64, err error) {{
\to := orm.NewOrm()
\tid, err = o.Insert(m)
\treturn
}}

func Get{model}ById(id int) (v *{model}, err error) {{
\to := orm.NewOrm()
\tv = &{model}{{Id: id}}
\tqs := o.QueryTable(new({model})).Filter("Id", id).RelatedSel()
\tif err = qs.One(v); err == nil {{
\t\treturn v, nil
\t}}
\treturn nil, err
}}

func GetAll{model}(query map[string]string, fields []string, sortby []string, order []string,
\toffset int64, limit int64) (ml []interface{}, err error) {{
\to := orm.NewOrm()
\tqs := o.QueryTable(new({model})).RelatedSel()
\tfor k, v := range query {{
\t\tk = strings.Replace(k, ".", "__", -1)
\t\tif strings.Contains(k, "isnull") {{
\t\t\tqs = qs.Filter(k, (v == "true" || v == "1"))
\t\t}} else {{
\t\t\tqs = qs.Filter(k, v)
\t\t}}
\t}}

\tvar sortFields []string
\tif len(sortby) != 0 {{
\t\tif len(sortby) == len(order) {{
\t\t\tfor i, v := range sortby {{
\t\t\t\torderby := ""
\t\t\t\tif order[i] == "desc" {{
\t\t\t\t\torderby = "-" + v
\t\t\t\t}} else if order[i] == "asc" {{
\t\t\t\t\torderby = v
\t\t\t\t}} else {{
\t\t\t\t\treturn nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
\t\t\t\t}}
\t\t\t\tsortFields = append(sortFields, orderby)
\t\t\t}}
\t\t}} else if len(order) == 1 {{
\t\t\tfor _, v := range sortby {{
\t\t\t\torderby := ""
\t\t\t\tif order[0] == "desc" {{
\t\t\t\t\torderby = "-" + v
\t\t\t\t}} else if order[0] == "asc" {{
\t\t\t\t\torderby = v
\t\t\t\t}} else {{
\t\t\t\t\treturn nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
\t\t\t\t}}
\t\t\t\tsortFields = append(sortFields, orderby)
\t\t\t}}
\t\t}} else {{
\t\t\treturn nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
\t\t}}
\t}} else if len(order) != 0 {{
\t\treturn nil, errors.New("Error: unused 'order' fields")
\t}}

\tvar l []{model}
\tqs = qs.OrderBy(sortFields...)
\tif _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {{
\t\tif len(fields) == 0 {{
\t\t\tfor _, v := range l {{
\t\t\t\tml = append(ml, v)
\t\t\t}}
\t\t}} else {{
\t\t\tfor _, v := range l {{
\t\t\t\tm := make(map[string]interface{{}})
\t\t\t\tval := reflect.ValueOf(v)
\t\t\t\tfor _, fname := range fields {{
\t\t\t\t\tm[fname] = val.FieldByName(fname).Interface()
\t\t\t\t}}
\t\t\t\tml = append(ml, m)
\t\t\t}}
\t\t}}
\t\treturn ml, nil
\t}}
\treturn nil, err
}}

func Update{model}ById(m *{model}) (err error) {{
\to := orm.NewOrm()
\tv := {model}{{Id: m.Id}}
\tif err = o.Read(&v); err == nil {{
\t\tvar num int64
\t\tif num, err = o.Update(m); err == nil {{
\t\t\tfmt.Println("Number of records updated in database:", num)
\t\t}}
\t}}
\treturn
}}

func Delete{model}(id int) (err error) {{
\to := orm.NewOrm()
\tv := {model}{{Id: id}}
\tif err = o.Read(&v); err == nil {{
\t\tvar num int64
\t\tif num, err = o.Delete(&{model}{{Id: id}}); err == nil {{
\t\t\tfmt.Println("Number of records deleted in database:", num)
\t\t}}
\t}}
\treturn
}}
'''
    new_body = new_body_template.replace('{model}', model)
    new_body = new_body.replace('{{{{', '{{').replace('}}}}', '}}').replace('{{', '{').replace('}}', '}')
    text = text[:before.start()] + new_body
    file.write_text(text, encoding='utf-8')

print('models done')
