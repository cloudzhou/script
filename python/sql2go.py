
def to_goname(name):
    ss = name.split('_')
    goname = ''
    for s in ss:
        if s in ["id", "uid", "url"]:
            goname = goname + s.upper()
        else:
            goname = goname + s.capitalize()
    return goname

for x in open("/tmp/tt"):
    x = x.strip()
    xs = x.split(' ')
    column = xs[0]
    goname = to_goname(column)
    gotype = ''
    type = xs[1].lower()
    if 'bigint' in type:
        gotype = 'int64'
    elif 'int' in type:
        gotype = 'int'
    elif 'char' in type:
        gotype = 'string'
    elif 'float' in type or 'double' in type:
        gotype = 'float64'
    else:
        continue
    comment = x[x.index(' ')+1:]
    print '\t%s %s `ddb:"%s"` // %s' % (goname, gotype, column, comment)
