package tests

import "github.com/stretchr/testify/require"

func (s *SuiteTest) TestGetFeedbacksByFeedbackType() {

	feedbacks, err := s.api.GetFeedbacksByFeedbackType("read", "{\"FeedbackType\":\"read\",\"UserId\":\"0-vortex\",\"ItemId\":\"nushell:nushell\"}", 10)

	s.T().Log(feedbacks)

	require.NoError(s.T(), err, err)
	require.NotNil(s.T(), feedbacks)
}
