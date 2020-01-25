package mysql

import "testing"

import "github.com/stretchr/testify/assert"

var testRemoveTrailingCommaTables = []struct {
	in  string
	out string
}{
	{"asdfasdf,", "asdfasdf"},
	{"asdfasdf, ", "asdfasdf"},
	{"asdfasdf", "asdfasdf"},
	{"", ""},
	{",", ""},
}

func TestRemoveTrailingComma(t *testing.T) {
	for _, tt := range testRemoveTrailingCommaTables {
		assert.Equal(t, tt.out, removeTrailingComma(tt.in))
	}
}

var testRemoveKeyWordsTables = []struct {
	in  string
	out string
}{
	{
		"CREATE TABLE `auth_group_permissions` (\n" +
			"`id` int(11) NOT NULL AUTO_INCREMENT,\n" +
			"`group_id` int(11) NOT NULL,\n" +
			"`permission_id` int(11) NOT NULL,\n" +
			"PRIMARY KEY (`id`),\n" +
			"UNIQUE KEY `auth_group_permissions_group_id_permission_id_0cd325b0_uniq` (`group_id`,`permission_id`),\n" +
			"KEY `auth_group_permissio_permission_id_84c5c92e_fk_auth_perm` (`permission_id`),\n" +
			"CONSTRAINT `auth_group_permissio_permission_id_84c5c92e_fk_auth_perm` FOREIGN KEY (`permission_id`) REFERENCES `auth_permission` (`id`),\n" +
			"CONSTRAINT `auth_group_permissions_group_id_b120cbf9_fk_auth_group_id` FOREIGN KEY (`group_id`) REFERENCES `auth_group` (`id`)\n" +
			") ENGINE=InnoDB DEFAULT CHARSET=latin1",
		"CREATE TABLE `auth_group_permissions` (\n" +
			"`id` int(11) NOT NULL AUTO_INCREMENT,\n" +
			"`group_id` int(11) NOT NULL,\n" +
			"`permission_id` int(11) NOT NULL,\n" +
			"PRIMARY KEY (`id`),\n" +
			"UNIQUE KEY `auth_group_permissions_group_id_permission_id_0cd325b0_uniq` (`group_id`,`permission_id`),\n" +
			"KEY `auth_group_permissio_permission_id_84c5c92e_fk_auth_perm` (`permission_id`)\n" +
			") ENGINE=InnoDB DEFAULT CHARSET=latin1",
	},
}

func TestRemoveKeyWords(t *testing.T) {
	for _, tt := range testRemoveKeyWordsTables {
		assert.Equal(t, tt.out, removeKeywords(tt.in))
	}
}
