package mturk_test

import (
	"github.com/mitchellh/goamz/aws"
	"github.com/calendreco/goamz/exp/mturk"
	"github.com/mitchellh/goamz/testutil"
	. "github.com/motain/gocheck"
	"net/url"
	"testing"
	// "log"
)

func Test(t *testing.T) {
	TestingT(t)
}

var _ = Suite(&S{})

type S struct {
	mturk *mturk.MTurk
}

var testServer = testutil.NewHTTPServer()

func (s *S) SetUpSuite(c *C) {
	testServer.Start()
	auth := aws.Auth{"AKIAJEPMI3TL2L3T253A", "aLIzrSLAsJQjBCYuEn5xMiUsxmN5urWo6VI/ONrG", ""}
	u, err := url.Parse(testServer.URL)
	if err != nil {
		panic(err.Error())
	}

	s.mturk = &mturk.MTurk{
		Auth: auth,
		URL:  u,
	}
}

func (s *S) TearDownTest(c *C) {
	testServer.Flush()
}

func (s *S) TestCreateHIT(c *C) {
	testServer.Response(200, nil, BasicHitResponse)

	question := mturk.ExternalQuestion{
		ExternalURL: "http://www.amazon.com",
		FrameHeight: 200,
	}
	reward := mturk.Price{
		Amount:       "0.01",
		CurrencyCode: "USD",
	}
	hit, err := s.mturk.CreateHIT("title", "description", question, reward, 1, 2, "key1,key2", 3, nil, "annotation")

	testServer.WaitRequest()

	c.Assert(err, IsNil)
	c.Assert(hit, NotNil)

	c.Assert(hit.HITId, Equals, "28J4IXKO2L927XKJTHO34OCDNASCDW")
	c.Assert(hit.HITTypeId, Equals, "2XZ7D1X3V0FKQVW7LU51S7PKKGFKDF")
}

func (s *S) TestSearchHITs(c *C) {
	testServer.Response(200, nil, SearchHITResponse)

	hitResult, err := s.mturk.SearchHITs()

	c.Assert(err, IsNil)
	c.Assert(hitResult, NotNil)

	c.Assert(hitResult.NumResults, Equals, uint(1))
	c.Assert(hitResult.PageNumber, Equals, uint(1))
	c.Assert(hitResult.TotalNumResults, Equals, uint(1))

	c.Assert(len(hitResult.HITs), Equals, 1)
	c.Assert(hitResult.HITs[0].HITId, Equals, "2BU26DG67D1XTE823B3OQ2JF2XWF83")
	c.Assert(hitResult.HITs[0].HITTypeId, Equals, "22OWJ5OPB0YV6IGL5727KP9U38P5XR")
	c.Assert(hitResult.HITs[0].CreationTime, Equals, "2011-12-28T19:56:20Z")
	c.Assert(hitResult.HITs[0].Title, Equals, "test hit")
	c.Assert(hitResult.HITs[0].Description, Equals, "please disregard, testing only")
	c.Assert(hitResult.HITs[0].HITStatus, Equals, "Reviewable")
	c.Assert(hitResult.HITs[0].MaxAssignments, Equals, uint(1))
	c.Assert(hitResult.HITs[0].Reward.Amount, Equals, "0.01")
	c.Assert(hitResult.HITs[0].Reward.CurrencyCode, Equals, "USD")
	c.Assert(hitResult.HITs[0].AutoApprovalDelayInSeconds, Equals, uint(2592000))
	c.Assert(hitResult.HITs[0].AssignmentDurationInSeconds, Equals, uint(30))
	c.Assert(hitResult.HITs[0].NumberOfAssignmentsPending, Equals, uint(0))
	c.Assert(hitResult.HITs[0].NumberOfAssignmentsAvailable, Equals, uint(1))
	c.Assert(hitResult.HITs[0].NumberOfAssignmentsCompleted, Equals, uint(0))
}

func (s *S) TestGetAssignment(c *C){
	testServer.Response(200, nil, BasicAssignmentResponse)
	assignmentResult, err := s.mturk.GetAssignment("3DR23U6WE6RZ27DEYI0P5KERFGUTEW")
	c.Assert(err, IsNil)
	c.Assert(assignmentResult, NotNil)
	c.Assert(assignmentResult.Assignment.AssignmentId, Equals, "3DR23U6WE6RZ27DEYI0P5KERFGUTEW")
	c.Assert(assignmentResult.Assignment.WorkerId, Equals, "A100VV1V2XDQI7")
	c.Assert(assignmentResult.Assignment.HITId, Equals, "307FVKVSYRSRPNLK67G92IFZ9ED74M")
	c.Assert(assignmentResult.Assignment.AssignmentStatus, Equals, "Approved")
	c.Assert(assignmentResult.Assignment.AutoApprovalTime, Equals, "2014-05-05T06:10:41Z")
	c.Assert(assignmentResult.Assignment.AcceptTime, Equals, "2014-05-04T22:07:54Z")
	c.Assert(assignmentResult.Assignment.SubmitTime, Equals, "2014-05-04T22:10:41Z")
	c.Assert(assignmentResult.Assignment.ApprovalTime, Equals, "2014-05-05T06:12:23Z")

	c.Assert(len(assignmentResult.Assignment.Answers.QuestionFormAnswers.Answers), Equals, 1)
	c.Assert(assignmentResult.Assignment.Answers.QuestionFormAnswers.Answers[0].QuestionIdentifier, Equals, "Q1HasEvents")
	c.Assert(assignmentResult.Assignment.Answers.QuestionFormAnswers.Answers[0].FreeText, Equals, "no")
}