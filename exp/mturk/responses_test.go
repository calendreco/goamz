package mturk_test

var BasicHitResponse = `<?xml version="1.0"?>
<CreateHITResponse><OperationRequest><RequestId>643b794b-66b6-4427-bb8a-4d3df5c9a20e</RequestId></OperationRequest><HIT><Request><IsValid>True</IsValid></Request><HITId>28J4IXKO2L927XKJTHO34OCDNASCDW</HITId><HITTypeId>2XZ7D1X3V0FKQVW7LU51S7PKKGFKDF</HITTypeId></HIT></CreateHITResponse>
`

var SearchHITResponse = `<?xml version="1.0"?>
<SearchHITsResponse><OperationRequest><RequestId>38862d9c-f015-4177-a2d3-924110a9d6f2</RequestId></OperationRequest><SearchHITsResult><Request><IsValid>True</IsValid></Request><NumResults>1</NumResults><TotalNumResults>1</TotalNumResults><PageNumber>1</PageNumber><HIT><HITId>2BU26DG67D1XTE823B3OQ2JF2XWF83</HITId><HITTypeId>22OWJ5OPB0YV6IGL5727KP9U38P5XR</HITTypeId><CreationTime>2011-12-28T19:56:20Z</CreationTime><Title>test hit</Title><Description>please disregard, testing only</Description><HITStatus>Reviewable</HITStatus><MaxAssignments>1</MaxAssignments><Reward><Amount>0.01</Amount><CurrencyCode>USD</CurrencyCode><FormattedPrice>$0.01</FormattedPrice></Reward><AutoApprovalDelayInSeconds>2592000</AutoApprovalDelayInSeconds><Expiration>2011-12-28T19:56:50Z</Expiration><AssignmentDurationInSeconds>30</AssignmentDurationInSeconds><NumberOfAssignmentsPending>0</NumberOfAssignmentsPending><NumberOfAssignmentsAvailable>1</NumberOfAssignmentsAvailable><NumberOfAssignmentsCompleted>0</NumberOfAssignmentsCompleted></HIT></SearchHITsResult></SearchHITsResponse>
`

var BasicAssignmentResponse = `<?xml version="1.0"?>
<GetAssignmentResponse><OperationRequest><RequestId>c4954055-1a6c-4e5a-a33d-ff1fa7ac98b5</RequestId></OperationRequest><GetAssignmentResult><Request><IsValid>True</IsValid></Request><Assignment><AssignmentId>3DR23U6WE6RZ27DEYI0P5KERFGUTEW</AssignmentId><WorkerId>A100VV1V2XDQI7</WorkerId><HITId>307FVKVSYRSRPNLK67G92IFZ9ED74M</HITId><AssignmentStatus>Approved</AssignmentStatus><AutoApprovalTime>2014-05-05T06:10:41Z</AutoApprovalTime><AcceptTime>2014-05-04T22:07:54Z</AcceptTime><SubmitTime>2014-05-04T22:10:41Z</SubmitTime><ApprovalTime>2014-05-05T06:12:23Z</ApprovalTime><Answer>&lt;?xml version="1.0" encoding="UTF-8" standalone="no"?&gt;
&lt;QuestionFormAnswers xmlns="http://mechanicalturk.amazonaws.com/AWSMechanicalTurkDataSchemas/2005-10-01/QuestionFormAnswers.xsd"&gt;
&lt;Answer&gt;
&lt;QuestionIdentifier&gt;Q1HasEvents&lt;/QuestionIdentifier&gt;
&lt;FreeText&gt;no&lt;/FreeText&gt;
&lt;/Answer&gt;
&lt;/QuestionFormAnswers&gt;
</Answer></Assignment><HIT><HITId>307FVKVSYRSRPNLK67G92IFZ9ED74M</HITId><HITTypeId>3LJ9P5KQ15SEBJ4U29CY8J4HJJBTYQ</HITTypeId></HIT></GetAssignmentResult></GetAssignmentResponse>
`