package searchtest

import (
	"testing"

	"github.com/mattermost/mattermost-server/v5/store"
	"github.com/stretchr/testify/require"
)

var searchChannelStoreTests = []searchTest{
	{
		"Should be able to autocomplete a channel by name",
		testAutocompleteChannelByName,
		[]string{ENGINE_ALL},
	},
	{
		"Should be able to autocomplete a channel by display name",
		testAutocompleteChannelByDisplayName,
		[]string{ENGINE_ALL},
	},
	{
		"Should be able to autocomplete a channel by a part of its name when has parts splitted by - character",
		testAutocompleteChannelByNameSplittedWithDashChar,
		[]string{ENGINE_ALL},
	},
	{
		"Should be able to autocomplete a channel by a part of its name when has parts splitted by , character",
		testAutocompleteChannelByNameSplittedWithCommaChar,
		[]string{ENGINE_ALL},
	},
	{
		"Should be able to autocomplete a channel by a part of its name when has parts splitted by _ character",
		testAutocompleteChannelByNameSplittedWithUnderscoreChar,
		[]string{ENGINE_ALL},
	},
	{
		"Should be able to autocomplete a channel by a part of its display name when has parts splitted by whitespace character",
		testAutocompleteChannelByDisplayNameSplittedByWhitespaces,
		[]string{ENGINE_ALL},
	},
	{
		"Should be able to autocomplete retrieving all channels if the term is empty",
		testAutocompleteChannelByDisplayNameSplittedByWhitespaces,
		[]string{ENGINE_ALL},
	},
	{
		"Should be able to autocomplete channels in a case insensitive manner",
		testSearchChannelsInCaseInsensitiveManner,
		[]string{ENGINE_ALL},
	},
	{
		"Should autocomplete only returning public channels",
		testSearchOnlyPublicChannels,
		[]string{ENGINE_ALL},
	},
	{
		"Should support to autocomplete having a hyphen as the last character",
		testSearchShouldSupportHavingHyphenAsLastCharacter,
		[]string{ENGINE_ALL},
	},
	{
		"Should support to autocomplete with archived channels",
		testSearchShouldSupportAutocompleteWithArchivedChannels,
		[]string{ENGINE_ALL},
	},
}

func TestSearchChannelStore(t *testing.T, s store.Store, testEngine *SearchTestEngine) {
	th := &SearchTestHelper{
		Store: s,
	}
	err := th.InitFixtures()
	require.Nil(t, err)
	defer th.CleanFixtures()
	runTestSearch(t, testEngine, searchChannelStoreTests, th)
}

func testAutocompleteChannelByName(t *testing.T, th *SearchTestHelper) {
	res, err := th.Store.Channel().AutocompleteInTeam(th.Team.Id, "ChannelA", false)
	require.Nil(t, err)
	channelIds := make([]string, len(*res))
	for i, channel := range *res {
		channelIds[i] = channel.Id
	}
	require.ElementsMatch(t, []string{th.ChannelA.Id, th.ChannelAlternate.Id}, channelIds)
}

func testAutocompleteChannelByDisplayName(t *testing.T, th *SearchTestHelper) {
	t.Skip("Test not implemented yet")
}

func testAutocompleteChannelByNameSplittedWithDashChar(t *testing.T, th *SearchTestHelper) {
	t.Skip("Test not implemented yet")
}

func testAutocompleteChannelByNameSplittedWithCommaChar(t *testing.T, th *SearchTestHelper) {
	t.Skip("Test not implemented yet")
}

func testAutocompleteChannelByNameSplittedWithUnderscoreChar(t *testing.T, th *SearchTestHelper) {
	t.Skip("Test not implemented yet")
}

func testAutocompleteChannelByDisplayNameSplittedByWhitespaces(t *testing.T, th *SearchTestHelper) {
	t.Skip("Test not implemented yet")
}

func testSearchChannelsInCaseInsensitiveManner(t *testing.T, th *SearchTestHelper) {
	t.Skip("Test not implemented yet")
}

func testSearchOnlyPublicChannels(t *testing.T, th *SearchTestHelper) {
	t.Skip("Test not implemented yet")
}

func testSearchShouldSupportWildcardAfterHyphen(t *testing.T, th *SearchTestHelper) {
	t.Skip("Test not implemented yet")
}

func testSearchShouldSupportHavingHyphenAsLastCharacter(t *testing.T, th *SearchTestHelper) {
	t.Skip("Test not implemented yet")
}

func testSearchShouldSupportAutocompleteWithArchivedChannels(t *testing.T, th *SearchTestHelper) {
	t.Skip("Test not implemented yet")
}
