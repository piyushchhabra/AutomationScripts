import re
import mysql.connector


def getfieldname(name):
    return re.sub(r'_([a-z])', lambda x: x.group(1).upper(), name)


def getclassname(name):
    return re.sub(r'_([a-z])', lambda x: x.group(1).upper(), name[0].upper()+name[1:])


def getFieldType(dbType):
    if 'int' in dbType:
        return 'Integer'
    elif 'varchar' in dbType:
        return 'String'
    elif 'longtext' in dbType:
        return 'String'
    elif 'date' in dbType:
        return 'Date'
    elif 'double' in dbType:
        return 'Double'
    else:
        return 'Object'


def printjavaclass(tableStructure, tablename):
    print('@Entity\n@Data\n@ToString\n@Table(name="'+tablename+'")')
    print('public class ' + getclassname(tablename) + ' {')
    print()
    for x in tableStructure:
        fieldName = x['Field']
        type = x['Type']
        nullAllow = x['Null']
        key = x['Key']
        deafult = x['Default']
        extra = x['Extra']
        if 'auto_increment' in extra:
            print('@Id\n@GeneratedValue(\nstrategy = GenerationType.IDENTITY\n)')
        else:
            print('@SerializedName("' + fieldName + '")')
            print('@Column("name=' + fieldName + '"', end='')
            if key == 'UNI':
                print(', unique=true', end='')
            if nullAllow == 'NO':
                print(', nullable = false)')
            else:
                print(')')

        print('private ' + getFieldType(type) + ' ' + getfieldname(fieldName) + ';')
        print('\n')
    print('}')


def main():
    db = mysql.connector.connect(
        host="localhost",
        user="root",
        passwd="",
        database="demo_test"
    )

    table_name = input("Enter table name:\n")
    cursor = db.cursor(dictionary=True)
    query = "desc "+table_name
    cursor.execute(query)

    result = cursor.fetchall()
    printjavaclass(result, table_name)


if __name__ == '__main__':
    main()