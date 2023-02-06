// Copyright 2020
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
// Code generated by gen-accessors; DO NOT EDIT.
package api

// GetDataType returns the DataType field if it's non-nil, zero value otherwise.
func (a *Attribute) GetDataType() string {
	if a == nil || a.DataType == nil {
		return ""
	}
	return *a.DataType
}

// GetDescription returns the Description field if it's non-nil, zero value otherwise.
func (a *Attribute) GetDescription() string {
	if a == nil || a.Description == nil {
		return ""
	}
	return *a.Description
}

// GetDisplayName returns the DisplayName field if it's non-nil, zero value otherwise.
func (a *Attribute) GetDisplayName() string {
	if a == nil || a.DisplayName == nil {
		return ""
	}
	return *a.DisplayName
}

// GetID returns the ID field if it's non-nil, zero value otherwise.
func (a *Attribute) GetID() string {
	if a == nil || a.ID == nil {
		return ""
	}
	return *a.ID
}

// GetIsSearchable returns the IsSearchable field if it's non-nil, zero value otherwise.
func (a *Attribute) GetIsSearchable() bool {
	if a == nil || a.IsSearchable == nil {
		return false
	}
	return *a.IsSearchable
}

// GetOrganizationId returns the OrganizationId field if it's non-nil, zero value otherwise.
func (a *Attribute) GetOrganizationId() string {
	if a == nil || a.OrganizationId == nil {
		return ""
	}
	return *a.OrganizationId
}

// HasSuggestedValues checks if Attribute has any SuggestedValues.
func (a *Attribute) HasSuggestedValues() bool {
	if a == nil || a.SuggestedValues == nil {
		return false
	}
	if len(a.SuggestedValues) == 0 {
		return false
	}
	return true
}

// GetTrafficTypeID returns the TrafficTypeID field if it's non-nil, zero value otherwise.
func (a *Attribute) GetTrafficTypeID() string {
	if a == nil || a.TrafficTypeID == nil {
		return ""
	}
	return *a.TrafficTypeID
}

// GetDataType returns the DataType field if it's non-nil, zero value otherwise.
func (a *AttributeRequest) GetDataType() string {
	if a == nil || a.DataType == nil {
		return ""
	}
	return *a.DataType
}

// GetDescription returns the Description field if it's non-nil, zero value otherwise.
func (a *AttributeRequest) GetDescription() string {
	if a == nil || a.Description == nil {
		return ""
	}
	return *a.Description
}

// GetDisplayName returns the DisplayName field if it's non-nil, zero value otherwise.
func (a *AttributeRequest) GetDisplayName() string {
	if a == nil || a.DisplayName == nil {
		return ""
	}
	return *a.DisplayName
}

// GetIdentifier returns the Identifier field if it's non-nil, zero value otherwise.
func (a *AttributeRequest) GetIdentifier() string {
	if a == nil || a.Identifier == nil {
		return ""
	}
	return *a.Identifier
}

// GetIsSearchable returns the IsSearchable field if it's non-nil, zero value otherwise.
func (a *AttributeRequest) GetIsSearchable() bool {
	if a == nil || a.IsSearchable == nil {
		return false
	}
	return *a.IsSearchable
}

// HasSuggestedValues checks if AttributeRequest has any SuggestedValues.
func (a *AttributeRequest) HasSuggestedValues() bool {
	if a == nil || a.SuggestedValues == nil {
		return false
	}
	if len(a.SuggestedValues) == 0 {
		return false
	}
	return true
}

// GetTrafficTypeID returns the TrafficTypeID field if it's non-nil, zero value otherwise.
func (a *AttributeRequest) GetTrafficTypeID() string {
	if a == nil || a.TrafficTypeID == nil {
		return ""
	}
	return *a.TrafficTypeID
}

// GetSize returns the Size field if it's non-nil, zero value otherwise.
func (b *Bucket) GetSize() int {
	if b == nil || b.Size == nil {
		return 0
	}
	return *b.Size
}

// GetTreatment returns the Treatment field if it's non-nil, zero value otherwise.
func (b *Bucket) GetTreatment() string {
	if b == nil || b.Treatment == nil {
		return ""
	}
	return *b.Treatment
}

// GetCombiner returns the Combiner field if it's non-nil, zero value otherwise.
func (c *Condition) GetCombiner() string {
	if c == nil || c.Combiner == nil {
		return ""
	}
	return *c.Combiner
}

// HasMatchers checks if Condition has any Matchers.
func (c *Condition) HasMatchers() bool {
	if c == nil || c.Matchers == nil {
		return false
	}
	if len(c.Matchers) == 0 {
		return false
	}
	return true
}

// GetID returns the ID field if it's non-nil, zero value otherwise.
func (e *Environment) GetID() string {
	if e == nil || e.ID == nil {
		return ""
	}
	return *e.ID
}

// GetName returns the Name field if it's non-nil, zero value otherwise.
func (e *Environment) GetName() string {
	if e == nil || e.Name == nil {
		return ""
	}
	return *e.Name
}

// GetProduction returns the Production field if it's non-nil, zero value otherwise.
func (e *Environment) GetProduction() bool {
	if e == nil || e.Production == nil {
		return false
	}
	return *e.Production
}

// GetName returns the Name field if it's non-nil, zero value otherwise.
func (e *EnvironmentRequest) GetName() string {
	if e == nil || e.Name == nil {
		return ""
	}
	return *e.Name
}

// GetProduction returns the Production field if it's non-nil, zero value otherwise.
func (e *EnvironmentRequest) GetProduction() bool {
	if e == nil || e.Production == nil {
		return false
	}
	return *e.Production
}

// GetCreationTime returns the CreationTime field if it's non-nil, zero value otherwise.
func (e *EnvironmentSegment) GetCreationTime() int64 {
	if e == nil || e.CreationTime == nil {
		return 0
	}
	return *e.CreationTime
}

// GetCreator returns the Creator field.
func (e *EnvironmentSegment) GetCreator() *User {
	if e == nil {
		return nil
	}
	return e.Creator
}

// GetDescription returns the Description field if it's non-nil, zero value otherwise.
func (e *EnvironmentSegment) GetDescription() string {
	if e == nil || e.Description == nil {
		return ""
	}
	return *e.Description
}

// GetEnvironment returns the Environment field if it's non-nil, zero value otherwise.
func (e *EnvironmentSegment) GetEnvironment() string {
	if e == nil || e.Environment == nil {
		return ""
	}
	return *e.Environment
}

// GetID returns the ID field if it's non-nil, zero value otherwise.
func (e *EnvironmentSegment) GetID() string {
	if e == nil || e.ID == nil {
		return ""
	}
	return *e.ID
}

// GetLastUpdateTime returns the LastUpdateTime field if it's non-nil, zero value otherwise.
func (e *EnvironmentSegment) GetLastUpdateTime() int64 {
	if e == nil || e.LastUpdateTime == nil {
		return 0
	}
	return *e.LastUpdateTime
}

// GetName returns the Name field if it's non-nil, zero value otherwise.
func (e *EnvironmentSegment) GetName() string {
	if e == nil || e.Name == nil {
		return ""
	}
	return *e.Name
}

// GetOrgID returns the OrgID field if it's non-nil, zero value otherwise.
func (e *EnvironmentSegment) GetOrgID() string {
	if e == nil || e.OrgID == nil {
		return ""
	}
	return *e.OrgID
}

// GetStatus returns the Status field if it's non-nil, zero value otherwise.
func (e *EnvironmentSegment) GetStatus() string {
	if e == nil || e.Status == nil {
		return ""
	}
	return *e.Status
}

// GetTrafficTypeID returns the TrafficTypeID field if it's non-nil, zero value otherwise.
func (e *EnvironmentSegment) GetTrafficTypeID() string {
	if e == nil || e.TrafficTypeID == nil {
		return ""
	}
	return *e.TrafficTypeID
}

// GetTrafficTypeURN returns the TrafficTypeURN field.
func (e *EnvironmentSegment) GetTrafficTypeURN() *TrafficType {
	if e == nil {
		return nil
	}
	return e.TrafficTypeURN
}

// HasKeys checks if EnvironmentSegmentKeysRequest has any Keys.
func (e *EnvironmentSegmentKeysRequest) HasKeys() bool {
	if e == nil || e.Keys == nil {
		return false
	}
	if len(e.Keys) == 0 {
		return false
	}
	return true
}

// GetLimit returns the Limit field if it's non-nil, zero value otherwise.
func (g *GenericListResult) GetLimit() int {
	if g == nil || g.Limit == nil {
		return 0
	}
	return *g.Limit
}

// GetOffset returns the Offset field if it's non-nil, zero value otherwise.
func (g *GenericListResult) GetOffset() int {
	if g == nil || g.Offset == nil {
		return 0
	}
	return *g.Offset
}

// GetTotalCount returns the TotalCount field if it's non-nil, zero value otherwise.
func (g *GenericListResult) GetTotalCount() int {
	if g == nil || g.TotalCount == nil {
		return 0
	}
	return *g.TotalCount
}

// GetDescription returns the Description field if it's non-nil, zero value otherwise.
func (g *Group) GetDescription() string {
	if g == nil || g.Description == nil {
		return ""
	}
	return *g.Description
}

// GetID returns the ID field if it's non-nil, zero value otherwise.
func (g *Group) GetID() string {
	if g == nil || g.ID == nil {
		return ""
	}
	return *g.ID
}

// GetName returns the Name field if it's non-nil, zero value otherwise.
func (g *Group) GetName() string {
	if g == nil || g.Name == nil {
		return ""
	}
	return *g.Name
}

// GetType returns the Type field if it's non-nil, zero value otherwise.
func (g *Group) GetType() string {
	if g == nil || g.Type == nil {
		return ""
	}
	return *g.Type
}

// GetCount returns the Count field if it's non-nil, zero value otherwise.
func (g *GroupListResult) GetCount() int {
	if g == nil || g.Count == nil {
		return 0
	}
	return *g.Count
}

// HasData checks if GroupListResult has any Data.
func (g *GroupListResult) HasData() bool {
	if g == nil || g.Data == nil {
		return false
	}
	if len(g.Data) == 0 {
		return false
	}
	return true
}

// GetLimit returns the Limit field if it's non-nil, zero value otherwise.
func (g *GroupListResult) GetLimit() int {
	if g == nil || g.Limit == nil {
		return 0
	}
	return *g.Limit
}

// GetNextMarker returns the NextMarker field if it's non-nil, zero value otherwise.
func (g *GroupListResult) GetNextMarker() string {
	if g == nil || g.NextMarker == nil {
		return ""
	}
	return *g.NextMarker
}

// GetPreviousMarker returns the PreviousMarker field if it's non-nil, zero value otherwise.
func (g *GroupListResult) GetPreviousMarker() string {
	if g == nil || g.PreviousMarker == nil {
		return ""
	}
	return *g.PreviousMarker
}

// GetAttribute returns the Attribute field if it's non-nil, zero value otherwise.
func (m *Matcher) GetAttribute() string {
	if m == nil || m.Attribute == nil {
		return ""
	}
	return *m.Attribute
}

// GetBool returns the Bool field if it's non-nil, zero value otherwise.
func (m *Matcher) GetBool() bool {
	if m == nil || m.Bool == nil {
		return false
	}
	return *m.Bool
}

// GetDate returns the Date field if it's non-nil, zero value otherwise.
func (m *Matcher) GetDate() int {
	if m == nil || m.Date == nil {
		return 0
	}
	return *m.Date
}

// GetNegate returns the Negate field if it's non-nil, zero value otherwise.
func (m *Matcher) GetNegate() bool {
	if m == nil || m.Negate == nil {
		return false
	}
	return *m.Negate
}

// GetNumber returns the Number field if it's non-nil, zero value otherwise.
func (m *Matcher) GetNumber() int {
	if m == nil || m.Number == nil {
		return 0
	}
	return *m.Number
}

// GetString returns the String field if it's non-nil, zero value otherwise.
func (m *Matcher) GetString() string {
	if m == nil || m.String == nil {
		return ""
	}
	return *m.String
}

// HasStrings checks if Matcher has any Strings.
func (m *Matcher) HasStrings() bool {
	if m == nil || m.Strings == nil {
		return false
	}
	if len(m.Strings) == 0 {
		return false
	}
	return true
}

// GetType returns the Type field if it's non-nil, zero value otherwise.
func (m *Matcher) GetType() string {
	if m == nil || m.Type == nil {
		return ""
	}
	return *m.Type
}

// HasBuckets checks if Rule has any Buckets.
func (r *Rule) HasBuckets() bool {
	if r == nil || r.Buckets == nil {
		return false
	}
	if len(r.Buckets) == 0 {
		return false
	}
	return true
}

// GetCondition returns the Condition field.
func (r *Rule) GetCondition() *Condition {
	if r == nil {
		return nil
	}
	return r.Condition
}

// GetCreationTime returns the CreationTime field if it's non-nil, zero value otherwise.
func (s *Segment) GetCreationTime() int64 {
	if s == nil || s.CreationTime == nil {
		return 0
	}
	return *s.CreationTime
}

// GetDescription returns the Description field if it's non-nil, zero value otherwise.
func (s *Segment) GetDescription() string {
	if s == nil || s.Description == nil {
		return ""
	}
	return *s.Description
}

// GetEnvironment returns the Environment field.
func (s *Segment) GetEnvironment() *Environment {
	if s == nil {
		return nil
	}
	return s.Environment
}

// GetName returns the Name field if it's non-nil, zero value otherwise.
func (s *Segment) GetName() string {
	if s == nil || s.Name == nil {
		return ""
	}
	return *s.Name
}

// HasTags checks if Segment has any Tags.
func (s *Segment) HasTags() bool {
	if s == nil || s.Tags == nil {
		return false
	}
	if len(s.Tags) == 0 {
		return false
	}
	return true
}

// GetTrafficType returns the TrafficType field.
func (s *Segment) GetTrafficType() *TrafficType {
	if s == nil {
		return nil
	}
	return s.TrafficType
}

// GetKey returns the Key field if it's non-nil, zero value otherwise.
func (s *SegmentKey) GetKey() string {
	if s == nil || s.Key == nil {
		return ""
	}
	return *s.Key
}

// HasKeys checks if SegmentKeysList has any Keys.
func (s *SegmentKeysList) HasKeys() bool {
	if s == nil || s.Keys == nil {
		return false
	}
	if len(s.Keys) == 0 {
		return false
	}
	return true
}

// HasObjects checks if SegmentListResult has any Objects.
func (s *SegmentListResult) HasObjects() bool {
	if s == nil || s.Objects == nil {
		return false
	}
	if len(s.Objects) == 0 {
		return false
	}
	return true
}

// GetCreationTime returns the CreationTime field if it's non-nil, zero value otherwise.
func (s *Split) GetCreationTime() int64 {
	if s == nil || s.CreationTime == nil {
		return 0
	}
	return *s.CreationTime
}

// GetDescription returns the Description field if it's non-nil, zero value otherwise.
func (s *Split) GetDescription() string {
	if s == nil || s.Description == nil {
		return ""
	}
	return *s.Description
}

// GetID returns the ID field if it's non-nil, zero value otherwise.
func (s *Split) GetID() string {
	if s == nil || s.ID == nil {
		return ""
	}
	return *s.ID
}

// GetName returns the Name field if it's non-nil, zero value otherwise.
func (s *Split) GetName() string {
	if s == nil || s.Name == nil {
		return ""
	}
	return *s.Name
}

// GetRolloutStatus returns the RolloutStatus field.
func (s *Split) GetRolloutStatus() *SplitRolloutStatus {
	if s == nil {
		return nil
	}
	return s.RolloutStatus
}

// GetRolloutStatusTimestamp returns the RolloutStatusTimestamp field if it's non-nil, zero value otherwise.
func (s *Split) GetRolloutStatusTimestamp() int64 {
	if s == nil || s.RolloutStatusTimestamp == nil {
		return 0
	}
	return *s.RolloutStatusTimestamp
}

// HasTags checks if Split has any Tags.
func (s *Split) HasTags() bool {
	if s == nil || s.Tags == nil {
		return false
	}
	if len(s.Tags) == 0 {
		return false
	}
	return true
}

// GetTrafficType returns the TrafficType field.
func (s *Split) GetTrafficType() *TrafficType {
	if s == nil {
		return nil
	}
	return s.TrafficType
}

// GetCreationTIme returns the CreationTIme field if it's non-nil, zero value otherwise.
func (s *SplitDefinition) GetCreationTIme() int {
	if s == nil || s.CreationTIme == nil {
		return 0
	}
	return *s.CreationTIme
}

// HasDefaultRule checks if SplitDefinition has any DefaultRule.
func (s *SplitDefinition) HasDefaultRule() bool {
	if s == nil || s.DefaultRule == nil {
		return false
	}
	if len(s.DefaultRule) == 0 {
		return false
	}
	return true
}

// GetDefaultTreatment returns the DefaultTreatment field if it's non-nil, zero value otherwise.
func (s *SplitDefinition) GetDefaultTreatment() string {
	if s == nil || s.DefaultTreatment == nil {
		return ""
	}
	return *s.DefaultTreatment
}

// GetEnvironment returns the Environment field.
func (s *SplitDefinition) GetEnvironment() *Environment {
	if s == nil {
		return nil
	}
	return s.Environment
}

// GetID returns the ID field if it's non-nil, zero value otherwise.
func (s *SplitDefinition) GetID() string {
	if s == nil || s.ID == nil {
		return ""
	}
	return *s.ID
}

// GetKilled returns the Killed field if it's non-nil, zero value otherwise.
func (s *SplitDefinition) GetKilled() bool {
	if s == nil || s.Killed == nil {
		return false
	}
	return *s.Killed
}

// GetLastUpdateTime returns the LastUpdateTime field if it's non-nil, zero value otherwise.
func (s *SplitDefinition) GetLastUpdateTime() int {
	if s == nil || s.LastUpdateTime == nil {
		return 0
	}
	return *s.LastUpdateTime
}

// GetName returns the Name field if it's non-nil, zero value otherwise.
func (s *SplitDefinition) GetName() string {
	if s == nil || s.Name == nil {
		return ""
	}
	return *s.Name
}

// HasRules checks if SplitDefinition has any Rules.
func (s *SplitDefinition) HasRules() bool {
	if s == nil || s.Rules == nil {
		return false
	}
	if len(s.Rules) == 0 {
		return false
	}
	return true
}

// GetTrafficAllocation returns the TrafficAllocation field if it's non-nil, zero value otherwise.
func (s *SplitDefinition) GetTrafficAllocation() int {
	if s == nil || s.TrafficAllocation == nil {
		return 0
	}
	return *s.TrafficAllocation
}

// GetTrafficType returns the TrafficType field.
func (s *SplitDefinition) GetTrafficType() *TrafficType {
	if s == nil {
		return nil
	}
	return s.TrafficType
}

// HasTreatments checks if SplitDefinition has any Treatments.
func (s *SplitDefinition) HasTreatments() bool {
	if s == nil || s.Treatments == nil {
		return false
	}
	if len(s.Treatments) == 0 {
		return false
	}
	return true
}

// HasDefaultRule checks if SplitDefinitionRequest has any DefaultRule.
func (s *SplitDefinitionRequest) HasDefaultRule() bool {
	if s == nil || s.DefaultRule == nil {
		return false
	}
	if len(s.DefaultRule) == 0 {
		return false
	}
	return true
}

// HasRules checks if SplitDefinitionRequest has any Rules.
func (s *SplitDefinitionRequest) HasRules() bool {
	if s == nil || s.Rules == nil {
		return false
	}
	if len(s.Rules) == 0 {
		return false
	}
	return true
}

// HasTreatments checks if SplitDefinitionRequest has any Treatments.
func (s *SplitDefinitionRequest) HasTreatments() bool {
	if s == nil || s.Treatments == nil {
		return false
	}
	if len(s.Treatments) == 0 {
		return false
	}
	return true
}

// HasObjects checks if SplitDefinitions has any Objects.
func (s *SplitDefinitions) HasObjects() bool {
	if s == nil || s.Objects == nil {
		return false
	}
	if len(s.Objects) == 0 {
		return false
	}
	return true
}

// GetID returns the ID field if it's non-nil, zero value otherwise.
func (s *SplitRolloutStatus) GetID() string {
	if s == nil || s.ID == nil {
		return ""
	}
	return *s.ID
}

// GetName returns the Name field if it's non-nil, zero value otherwise.
func (s *SplitRolloutStatus) GetName() string {
	if s == nil || s.Name == nil {
		return ""
	}
	return *s.Name
}

// HasObjects checks if Splits has any Objects.
func (s *Splits) HasObjects() bool {
	if s == nil || s.Objects == nil {
		return false
	}
	if len(s.Objects) == 0 {
		return false
	}
	return true
}

// GetName returns the Name field if it's non-nil, zero value otherwise.
func (t *Tag) GetName() string {
	if t == nil || t.Name == nil {
		return ""
	}
	return *t.Name
}

// GetDisplayAttributeID returns the DisplayAttributeID field if it's non-nil, zero value otherwise.
func (t *TrafficType) GetDisplayAttributeID() string {
	if t == nil || t.DisplayAttributeID == nil {
		return ""
	}
	return *t.DisplayAttributeID
}

// GetID returns the ID field if it's non-nil, zero value otherwise.
func (t *TrafficType) GetID() string {
	if t == nil || t.ID == nil {
		return ""
	}
	return *t.ID
}

// GetName returns the Name field if it's non-nil, zero value otherwise.
func (t *TrafficType) GetName() string {
	if t == nil || t.Name == nil {
		return ""
	}
	return *t.Name
}

// GetType returns the Type field if it's non-nil, zero value otherwise.
func (t *TrafficType) GetType() string {
	if t == nil || t.Type == nil {
		return ""
	}
	return *t.Type
}

// GetWorkspace returns the Workspace field.
func (t *TrafficType) GetWorkspace() *Workspace {
	if t == nil {
		return nil
	}
	return t.Workspace
}

// GetConfigurations returns the Configurations field if it's non-nil, zero value otherwise.
func (t *Treatment) GetConfigurations() string {
	if t == nil || t.Configurations == nil {
		return ""
	}
	return *t.Configurations
}

// GetDescription returns the Description field if it's non-nil, zero value otherwise.
func (t *Treatment) GetDescription() string {
	if t == nil || t.Description == nil {
		return ""
	}
	return *t.Description
}

// HasKeys checks if Treatment has any Keys.
func (t *Treatment) HasKeys() bool {
	if t == nil || t.Keys == nil {
		return false
	}
	if len(t.Keys) == 0 {
		return false
	}
	return true
}

// GetName returns the Name field if it's non-nil, zero value otherwise.
func (t *Treatment) GetName() string {
	if t == nil || t.Name == nil {
		return ""
	}
	return *t.Name
}

// HasSegments checks if Treatment has any Segments.
func (t *Treatment) HasSegments() bool {
	if t == nil || t.Segments == nil {
		return false
	}
	if len(t.Segments) == 0 {
		return false
	}
	return true
}

// GetEmail returns the Email field if it's non-nil, zero value otherwise.
func (u *User) GetEmail() string {
	if u == nil || u.Email == nil {
		return ""
	}
	return *u.Email
}

// HasGroups checks if User has any Groups.
func (u *User) HasGroups() bool {
	if u == nil || u.Groups == nil {
		return false
	}
	if len(u.Groups) == 0 {
		return false
	}
	return true
}

// GetID returns the ID field if it's non-nil, zero value otherwise.
func (u *User) GetID() string {
	if u == nil || u.ID == nil {
		return ""
	}
	return *u.ID
}

// GetName returns the Name field if it's non-nil, zero value otherwise.
func (u *User) GetName() string {
	if u == nil || u.Name == nil {
		return ""
	}
	return *u.Name
}

// GetStatus returns the Status field if it's non-nil, zero value otherwise.
func (u *User) GetStatus() string {
	if u == nil || u.Status == nil {
		return ""
	}
	return *u.Status
}

// GetTFA returns the TFA field if it's non-nil, zero value otherwise.
func (u *User) GetTFA() bool {
	if u == nil || u.TFA == nil {
		return false
	}
	return *u.TFA
}

// GetType returns the Type field if it's non-nil, zero value otherwise.
func (u *User) GetType() string {
	if u == nil || u.Type == nil {
		return ""
	}
	return *u.Type
}

// HasGroups checks if UserCreateRequest has any Groups.
func (u *UserCreateRequest) HasGroups() bool {
	if u == nil || u.Groups == nil {
		return false
	}
	if len(u.Groups) == 0 {
		return false
	}
	return true
}

// GetCount returns the Count field if it's non-nil, zero value otherwise.
func (u *UserListResult) GetCount() int {
	if u == nil || u.Count == nil {
		return 0
	}
	return *u.Count
}

// HasData checks if UserListResult has any Data.
func (u *UserListResult) HasData() bool {
	if u == nil || u.Data == nil {
		return false
	}
	if len(u.Data) == 0 {
		return false
	}
	return true
}

// GetLimit returns the Limit field if it's non-nil, zero value otherwise.
func (u *UserListResult) GetLimit() int {
	if u == nil || u.Limit == nil {
		return 0
	}
	return *u.Limit
}

// GetNextMarker returns the NextMarker field if it's non-nil, zero value otherwise.
func (u *UserListResult) GetNextMarker() string {
	if u == nil || u.NextMarker == nil {
		return ""
	}
	return *u.NextMarker
}

// GetPreviousMarker returns the PreviousMarker field if it's non-nil, zero value otherwise.
func (u *UserListResult) GetPreviousMarker() string {
	if u == nil || u.PreviousMarker == nil {
		return ""
	}
	return *u.PreviousMarker
}

// GetTFA returns the TFA field if it's non-nil, zero value otherwise.
func (u *UserUpdateRequest) GetTFA() bool {
	if u == nil || u.TFA == nil {
		return false
	}
	return *u.TFA
}

// GetID returns the ID field if it's non-nil, zero value otherwise.
func (w *Workspace) GetID() string {
	if w == nil || w.ID == nil {
		return ""
	}
	return *w.ID
}

// GetName returns the Name field if it's non-nil, zero value otherwise.
func (w *Workspace) GetName() string {
	if w == nil || w.Name == nil {
		return ""
	}
	return *w.Name
}

// GetRequiresTitleAndComments returns the RequiresTitleAndComments field if it's non-nil, zero value otherwise.
func (w *Workspace) GetRequiresTitleAndComments() bool {
	if w == nil || w.RequiresTitleAndComments == nil {
		return false
	}
	return *w.RequiresTitleAndComments
}

// GetType returns the Type field if it's non-nil, zero value otherwise.
func (w *Workspace) GetType() string {
	if w == nil || w.Type == nil {
		return ""
	}
	return *w.Type
}

// GetName returns the Name field if it's non-nil, zero value otherwise.
func (w *WorkspaceRequest) GetName() string {
	if w == nil || w.Name == nil {
		return ""
	}
	return *w.Name
}

// GetRequiresTitleAndComments returns the RequiresTitleAndComments field if it's non-nil, zero value otherwise.
func (w *WorkspaceRequest) GetRequiresTitleAndComments() bool {
	if w == nil || w.RequiresTitleAndComments == nil {
		return false
	}
	return *w.RequiresTitleAndComments
}

// HasObjects checks if Workspaces has any Objects.
func (w *Workspaces) HasObjects() bool {
	if w == nil || w.Objects == nil {
		return false
	}
	if len(w.Objects) == 0 {
		return false
	}
	return true
}
