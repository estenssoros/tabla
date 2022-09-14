package gopher_test

import (
	"fmt"
	"testing"

	"github.com/estenssoros/tabla/internal/gopher"
	"github.com/estenssoros/tabla/internal/mysql"
	"github.com/stretchr/testify/assert"
	"github.com/xwb1989/sqlparser"
)

var testDropCreateTables = []struct {
	in  string
	out string
	err bool
}{
	{
		"type Charter struct {" +
			"ID                       int       `db:\"id,int,11\"`\n" +
			"CreatedAt                time.Time `db:\"created_at,datetime\"`\n" +
			"UpdatedAt                time.Time `db:\"updated_at,datetime\"`\n" +
			"Status                   string    `db:\"status,varchar,15\"`\n" +
			"AddCom                   float64   `db:\"add_com,double\"`\n" +
			"Discount                 float64   `db:\"discount,double\"`\n" +
			"BodDate                  time.Time `db:\"bod_date,date\"`\n" +
			"BodLink                  string    `db:\"bod_link,varchar,100\"`\n" +
			"CbodDate                 time.Time `db:\"cbod_date,date\"`\n" +
			"CbodLink                 string    `db:\"cbod_link,varchar,100\"`\n" +
			"CpLink                   string    `db:\"cp_link,varchar,100\"`\n" +
			"CpStatus                 string    `db:\"cp_status,varchar,20\"`\n" +
			"Lashing                  int       `db:\"lashing,int,11\"`\n" +
			"LashingBonus             float64   `db:\"lashing_bonus,double\"`\n" +
			"Cev                      int       `db:\"cev,int,11\"`\n" +
			"Reefer                   float64   `db:\"reefer,double\"`\n" +
			"ReeferBonus              float64   `db:\"reefer_bonus,double\"`\n" +
			"ReeferRepairsOfficer     float64   `db:\"reefer_repairs_officer,double\"`\n" +
			"ReeferRepairsRating      float64   `db:\"reefer_repairs_rating,double\"`\n" +
			"ReeferRepairsCadet       float64   `db:\"reefer_repairs_cadet,double\"`\n" +
			"LeakageCleaning          float64   `db:\"leakage_cleaning,double\"`\n" +
			"FuelTesting              float64   `db:\"fuel_testing,double\"`\n" +
			"FuelClause               string    `db:\"fuel_clause,longtext\"`\n" +
			"FuelNotes                string    `db:\"fuel_notes,longtext\"`\n" +
			"AddPremiums              string    `db:\"add_premiums,longtext\"`\n" +
			"DgClause                 string    `db:\"dg_clause,longtext\"`\n" +
			"PaymentTerms             string    `db:\"payment_terms,longtext\"`\n" +
			"Surveyor                 string    `db:\"surveyor,longtext\"`\n" +
			"SurveyorCost             float64   `db:\"surveyor_cost,double\"`\n" +
			"CharterPi                string    `db:\"charter_pi,varchar,50\"`\n" +
			"SeaspanPi                string    `db:\"seaspan_pi,varchar,50\"`\n" +
			"LaycanFrom               time.Time `db:\"laycan_from,date\"`\n" +
			"LaycanTo                 time.Time `db:\"laycan_to,date\"`\n" +
			"RedeliveryDate           time.Time `db:\"redelivery_date,datetime,6\"`\n" +
			"RedeliveryLink           string    `db:\"redelivery_link,varchar,100\"`\n" +
			"RedeliveryPlace          string    `db:\"redelivery_place,varchar,50\"`\n" +
			"RedeliveryNoticeSchedule string    `db:\"redelivery_notice_schedule,varchar,50\"`\n" +
			"EarliestEnd              time.Time `db:\"earliest_end,date\"`\n" +
			"LatestEnd                time.Time `db:\"latest_end,date\"`\n" +
			"DeliveryDate             time.Time `db:\"delivery_date,date\"`\n" +
			"DeliveryLink             string    `db:\"delivery_link,varchar,100\"`\n" +
			"DeliveryPlace            string    `db:\"delivery_place,varchar,50\"`\n" +
			"DeliveryPeriod           string    `db:\"delivery_period,varchar,50\"`\n" +
			"DeliveryNoticeSchedule   string    `db:\"delivery_notice_schedule,varchar,50\"`\n" +
			"DeliveryNoticeTendered   time.Time `db:\"delivery_notice_tendered,datetime,6\"`\n" +
			"CustomerId               int       `db:\"customer_id,int,11\"`\n" +
			"VesselId                 int       `db:\"vessel_id,int,11\"`\n" +
			"}",
		"DROP TABLE IF EXISTS `charter`;\n" +
			"CREATE TABLE `charter` (\n" +
			"    `id` int(11)\n" +
			"    , `created_at` datetime\n" +
			"    , `updated_at` datetime\n" +
			"    , `status` varchar(15)\n" +
			"    , `add_com` double\n" +
			"    , `discount` double\n" +
			"    , `bod_date` date\n" +
			"    , `bod_link` varchar(100)\n" +
			"    , `cbod_date` date\n" +
			"    , `cbod_link` varchar(100)\n" +
			"    , `cp_link` varchar(100)\n" +
			"    , `cp_status` varchar(20)\n" +
			"    , `lashing` int(11)\n" +
			"    , `lashing_bonus` double\n" +
			"    , `cev` int(11)\n" +
			"    , `reefer` double\n" +
			"    , `reefer_bonus` double\n" +
			"    , `reefer_repairs_officer` double\n" +
			"    , `reefer_repairs_rating` double\n" +
			"    , `reefer_repairs_cadet` double\n" +
			"    , `leakage_cleaning` double\n" +
			"    , `fuel_testing` double\n" +
			"    , `fuel_clause` longtext\n" +
			"    , `fuel_notes` longtext\n" +
			"    , `add_premiums` longtext\n" +
			"    , `dg_clause` longtext\n" +
			"    , `payment_terms` longtext\n" +
			"    , `surveyor` longtext\n" +
			"    , `surveyor_cost` double\n" +
			"    , `charter_pi` varchar(50)\n" +
			"    , `seaspan_pi` varchar(50)\n" +
			"    , `laycan_from` date\n" +
			"    , `laycan_to` date\n" +
			"    , `redelivery_date` datetime(6)\n" +
			"    , `redelivery_link` varchar(100)\n" +
			"    , `redelivery_place` varchar(50)\n" +
			"    , `redelivery_notice_schedule` varchar(50)\n" +
			"    , `earliest_end` date\n" +
			"    , `latest_end` date\n" +
			"    , `delivery_date` date\n" +
			"    , `delivery_link` varchar(100)\n" +
			"    , `delivery_place` varchar(50)\n" +
			"    , `delivery_period` varchar(50)\n" +
			"    , `delivery_notice_schedule` varchar(50)\n" +
			"    , `delivery_notice_tendered` datetime(6)\n" +
			"    , `customer_id` int(11)\n" +
			"    , `vessel_id` int(11)\n" +
			"    , PRIMARY KEY(`id`)\n" +
			");",
		false,
	},
	{"asdf", "", true},
}

func TestDropCreate(t *testing.T) {
	for i, tt := range testDropCreateTables {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			dropCreate, err := gopher.DropCreate(tt.in, mysql.Dialect{})
			if tt.err {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
				assert.Equal(t, tt.out, dropCreate)
			}
		})
	}
}

type testConverter struct{}

func (c testConverter) ColDefToGoField(colDef *sqlparser.ColumnDefinition, nulls bool) (*gopher.GoField, error) {
	return &gopher.GoField{}, nil
}

func (c testConverter) PrepareStatment(sql string) string {
	return sql
}

var testParseSrcTables = []struct {
	in  string
	err bool
}{
	{"CREATE TABLE `asdf` (id int);", false},
	{"DROP TABLE `asdf`;", true},
	{"INSERT INTO `asdf` (id) VALUES ('asdf');", true},
	{"ASDF `asdf` (id) VALUES ('asdf');", true},
}

func TestParseSQLToGoStruct(t *testing.T) {
	for i, tt := range testParseSrcTables {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			_, err := gopher.ParseSQLToGoStruct(tt.in, testConverter{}, false)
			if tt.err {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}
