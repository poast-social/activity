package propertyattachment

import (
	"fmt"
	vocab "github.com/go-fed/activity/streams/vocab"
	"net/url"
)

// AttachmentPropertyIterator is an iterator for a property. It is permitted to be
// one of multiple value types. At most, one type of value can be present, or
// none at all. Setting a value will clear the other types of values so that
// only one of the 'Is' methods will return true. It is possible to clear all
// values, so that this property is empty.
type AttachmentPropertyIterator struct {
	ObjectMember                vocab.ObjectInterface
	LinkMember                  vocab.LinkInterface
	AcceptMember                vocab.AcceptInterface
	ActivityMember              vocab.ActivityInterface
	AddMember                   vocab.AddInterface
	AnnounceMember              vocab.AnnounceInterface
	ApplicationMember           vocab.ApplicationInterface
	ArriveMember                vocab.ArriveInterface
	ArticleMember               vocab.ArticleInterface
	AudioMember                 vocab.AudioInterface
	BlockMember                 vocab.BlockInterface
	CollectionMember            vocab.CollectionInterface
	CollectionPageMember        vocab.CollectionPageInterface
	CreateMember                vocab.CreateInterface
	DeleteMember                vocab.DeleteInterface
	DislikeMember               vocab.DislikeInterface
	DocumentMember              vocab.DocumentInterface
	EventMember                 vocab.EventInterface
	FlagMember                  vocab.FlagInterface
	FollowMember                vocab.FollowInterface
	GroupMember                 vocab.GroupInterface
	IgnoreMember                vocab.IgnoreInterface
	ImageMember                 vocab.ImageInterface
	IntransitiveActivityMember  vocab.IntransitiveActivityInterface
	InviteMember                vocab.InviteInterface
	JoinMember                  vocab.JoinInterface
	LeaveMember                 vocab.LeaveInterface
	LikeMember                  vocab.LikeInterface
	ListenMember                vocab.ListenInterface
	MentionMember               vocab.MentionInterface
	MoveMember                  vocab.MoveInterface
	NoteMember                  vocab.NoteInterface
	OfferMember                 vocab.OfferInterface
	OrderedCollectionMember     vocab.OrderedCollectionInterface
	OrderedCollectionPageMember vocab.OrderedCollectionPageInterface
	OrganizationMember          vocab.OrganizationInterface
	PageMember                  vocab.PageInterface
	PersonMember                vocab.PersonInterface
	PlaceMember                 vocab.PlaceInterface
	ProfileMember               vocab.ProfileInterface
	QuestionMember              vocab.QuestionInterface
	ReadMember                  vocab.ReadInterface
	RejectMember                vocab.RejectInterface
	RelationshipMember          vocab.RelationshipInterface
	RemoveMember                vocab.RemoveInterface
	ServiceMember               vocab.ServiceInterface
	TentativeAcceptMember       vocab.TentativeAcceptInterface
	TentativeRejectMember       vocab.TentativeRejectInterface
	TombstoneMember             vocab.TombstoneInterface
	TravelMember                vocab.TravelInterface
	UndoMember                  vocab.UndoInterface
	UpdateMember                vocab.UpdateInterface
	VideoMember                 vocab.VideoInterface
	ViewMember                  vocab.ViewInterface
	unknown                     []byte
	iri                         *url.URL
	alias                       string
	myIdx                       int
	parent                      vocab.AttachmentPropertyInterface
}

// NewAttachmentPropertyIterator creates a new attachment property.
func NewAttachmentPropertyIterator() *AttachmentPropertyIterator {
	return &AttachmentPropertyIterator{alias: ""}
}

// deserializeAttachmentPropertyIterator creates an iterator from an element that
// has been unmarshalled from a text or binary format.
func deserializeAttachmentPropertyIterator(i interface{}, aliasMap map[string]string) (*AttachmentPropertyIterator, error) {
	alias := ""
	if a, ok := aliasMap["https://www.w3.org/TR/activitystreams-vocabulary"]; ok {
		alias = a
	}
	if s, ok := i.(string); ok {
		u, err := url.Parse(s)
		// If error exists, don't error out -- skip this and treat as unknown string ([]byte) at worst
		// Also, if no scheme exists, don't treat it as a URL -- net/url is greedy
		if err == nil && len(u.Scheme) > 0 {
			this := &AttachmentPropertyIterator{
				alias: alias,
				iri:   u,
			}
			return this, nil
		}
	}
	if m, ok := i.(map[string]interface{}); ok {
		if v, err := mgr.DeserializeObjectActivityStreams()(m, aliasMap); err == nil {
			this := &AttachmentPropertyIterator{
				ObjectMember: v,
				alias:        alias,
			}
			return this, nil
		} else if v, err := mgr.DeserializeLinkActivityStreams()(m, aliasMap); err == nil {
			this := &AttachmentPropertyIterator{
				LinkMember: v,
				alias:      alias,
			}
			return this, nil
		} else if v, err := mgr.DeserializeAcceptActivityStreams()(m, aliasMap); err == nil {
			this := &AttachmentPropertyIterator{
				AcceptMember: v,
				alias:        alias,
			}
			return this, nil
		} else if v, err := mgr.DeserializeActivityActivityStreams()(m, aliasMap); err == nil {
			this := &AttachmentPropertyIterator{
				ActivityMember: v,
				alias:          alias,
			}
			return this, nil
		} else if v, err := mgr.DeserializeAddActivityStreams()(m, aliasMap); err == nil {
			this := &AttachmentPropertyIterator{
				AddMember: v,
				alias:     alias,
			}
			return this, nil
		} else if v, err := mgr.DeserializeAnnounceActivityStreams()(m, aliasMap); err == nil {
			this := &AttachmentPropertyIterator{
				AnnounceMember: v,
				alias:          alias,
			}
			return this, nil
		} else if v, err := mgr.DeserializeApplicationActivityStreams()(m, aliasMap); err == nil {
			this := &AttachmentPropertyIterator{
				ApplicationMember: v,
				alias:             alias,
			}
			return this, nil
		} else if v, err := mgr.DeserializeArriveActivityStreams()(m, aliasMap); err == nil {
			this := &AttachmentPropertyIterator{
				ArriveMember: v,
				alias:        alias,
			}
			return this, nil
		} else if v, err := mgr.DeserializeArticleActivityStreams()(m, aliasMap); err == nil {
			this := &AttachmentPropertyIterator{
				ArticleMember: v,
				alias:         alias,
			}
			return this, nil
		} else if v, err := mgr.DeserializeAudioActivityStreams()(m, aliasMap); err == nil {
			this := &AttachmentPropertyIterator{
				AudioMember: v,
				alias:       alias,
			}
			return this, nil
		} else if v, err := mgr.DeserializeBlockActivityStreams()(m, aliasMap); err == nil {
			this := &AttachmentPropertyIterator{
				BlockMember: v,
				alias:       alias,
			}
			return this, nil
		} else if v, err := mgr.DeserializeCollectionActivityStreams()(m, aliasMap); err == nil {
			this := &AttachmentPropertyIterator{
				CollectionMember: v,
				alias:            alias,
			}
			return this, nil
		} else if v, err := mgr.DeserializeCollectionPageActivityStreams()(m, aliasMap); err == nil {
			this := &AttachmentPropertyIterator{
				CollectionPageMember: v,
				alias:                alias,
			}
			return this, nil
		} else if v, err := mgr.DeserializeCreateActivityStreams()(m, aliasMap); err == nil {
			this := &AttachmentPropertyIterator{
				CreateMember: v,
				alias:        alias,
			}
			return this, nil
		} else if v, err := mgr.DeserializeDeleteActivityStreams()(m, aliasMap); err == nil {
			this := &AttachmentPropertyIterator{
				DeleteMember: v,
				alias:        alias,
			}
			return this, nil
		} else if v, err := mgr.DeserializeDislikeActivityStreams()(m, aliasMap); err == nil {
			this := &AttachmentPropertyIterator{
				DislikeMember: v,
				alias:         alias,
			}
			return this, nil
		} else if v, err := mgr.DeserializeDocumentActivityStreams()(m, aliasMap); err == nil {
			this := &AttachmentPropertyIterator{
				DocumentMember: v,
				alias:          alias,
			}
			return this, nil
		} else if v, err := mgr.DeserializeEventActivityStreams()(m, aliasMap); err == nil {
			this := &AttachmentPropertyIterator{
				EventMember: v,
				alias:       alias,
			}
			return this, nil
		} else if v, err := mgr.DeserializeFlagActivityStreams()(m, aliasMap); err == nil {
			this := &AttachmentPropertyIterator{
				FlagMember: v,
				alias:      alias,
			}
			return this, nil
		} else if v, err := mgr.DeserializeFollowActivityStreams()(m, aliasMap); err == nil {
			this := &AttachmentPropertyIterator{
				FollowMember: v,
				alias:        alias,
			}
			return this, nil
		} else if v, err := mgr.DeserializeGroupActivityStreams()(m, aliasMap); err == nil {
			this := &AttachmentPropertyIterator{
				GroupMember: v,
				alias:       alias,
			}
			return this, nil
		} else if v, err := mgr.DeserializeIgnoreActivityStreams()(m, aliasMap); err == nil {
			this := &AttachmentPropertyIterator{
				IgnoreMember: v,
				alias:        alias,
			}
			return this, nil
		} else if v, err := mgr.DeserializeImageActivityStreams()(m, aliasMap); err == nil {
			this := &AttachmentPropertyIterator{
				ImageMember: v,
				alias:       alias,
			}
			return this, nil
		} else if v, err := mgr.DeserializeIntransitiveActivityActivityStreams()(m, aliasMap); err == nil {
			this := &AttachmentPropertyIterator{
				IntransitiveActivityMember: v,
				alias:                      alias,
			}
			return this, nil
		} else if v, err := mgr.DeserializeInviteActivityStreams()(m, aliasMap); err == nil {
			this := &AttachmentPropertyIterator{
				InviteMember: v,
				alias:        alias,
			}
			return this, nil
		} else if v, err := mgr.DeserializeJoinActivityStreams()(m, aliasMap); err == nil {
			this := &AttachmentPropertyIterator{
				JoinMember: v,
				alias:      alias,
			}
			return this, nil
		} else if v, err := mgr.DeserializeLeaveActivityStreams()(m, aliasMap); err == nil {
			this := &AttachmentPropertyIterator{
				LeaveMember: v,
				alias:       alias,
			}
			return this, nil
		} else if v, err := mgr.DeserializeLikeActivityStreams()(m, aliasMap); err == nil {
			this := &AttachmentPropertyIterator{
				LikeMember: v,
				alias:      alias,
			}
			return this, nil
		} else if v, err := mgr.DeserializeListenActivityStreams()(m, aliasMap); err == nil {
			this := &AttachmentPropertyIterator{
				ListenMember: v,
				alias:        alias,
			}
			return this, nil
		} else if v, err := mgr.DeserializeMentionActivityStreams()(m, aliasMap); err == nil {
			this := &AttachmentPropertyIterator{
				MentionMember: v,
				alias:         alias,
			}
			return this, nil
		} else if v, err := mgr.DeserializeMoveActivityStreams()(m, aliasMap); err == nil {
			this := &AttachmentPropertyIterator{
				MoveMember: v,
				alias:      alias,
			}
			return this, nil
		} else if v, err := mgr.DeserializeNoteActivityStreams()(m, aliasMap); err == nil {
			this := &AttachmentPropertyIterator{
				NoteMember: v,
				alias:      alias,
			}
			return this, nil
		} else if v, err := mgr.DeserializeOfferActivityStreams()(m, aliasMap); err == nil {
			this := &AttachmentPropertyIterator{
				OfferMember: v,
				alias:       alias,
			}
			return this, nil
		} else if v, err := mgr.DeserializeOrderedCollectionActivityStreams()(m, aliasMap); err == nil {
			this := &AttachmentPropertyIterator{
				OrderedCollectionMember: v,
				alias:                   alias,
			}
			return this, nil
		} else if v, err := mgr.DeserializeOrderedCollectionPageActivityStreams()(m, aliasMap); err == nil {
			this := &AttachmentPropertyIterator{
				OrderedCollectionPageMember: v,
				alias:                       alias,
			}
			return this, nil
		} else if v, err := mgr.DeserializeOrganizationActivityStreams()(m, aliasMap); err == nil {
			this := &AttachmentPropertyIterator{
				OrganizationMember: v,
				alias:              alias,
			}
			return this, nil
		} else if v, err := mgr.DeserializePageActivityStreams()(m, aliasMap); err == nil {
			this := &AttachmentPropertyIterator{
				PageMember: v,
				alias:      alias,
			}
			return this, nil
		} else if v, err := mgr.DeserializePersonActivityStreams()(m, aliasMap); err == nil {
			this := &AttachmentPropertyIterator{
				PersonMember: v,
				alias:        alias,
			}
			return this, nil
		} else if v, err := mgr.DeserializePlaceActivityStreams()(m, aliasMap); err == nil {
			this := &AttachmentPropertyIterator{
				PlaceMember: v,
				alias:       alias,
			}
			return this, nil
		} else if v, err := mgr.DeserializeProfileActivityStreams()(m, aliasMap); err == nil {
			this := &AttachmentPropertyIterator{
				ProfileMember: v,
				alias:         alias,
			}
			return this, nil
		} else if v, err := mgr.DeserializeQuestionActivityStreams()(m, aliasMap); err == nil {
			this := &AttachmentPropertyIterator{
				QuestionMember: v,
				alias:          alias,
			}
			return this, nil
		} else if v, err := mgr.DeserializeReadActivityStreams()(m, aliasMap); err == nil {
			this := &AttachmentPropertyIterator{
				ReadMember: v,
				alias:      alias,
			}
			return this, nil
		} else if v, err := mgr.DeserializeRejectActivityStreams()(m, aliasMap); err == nil {
			this := &AttachmentPropertyIterator{
				RejectMember: v,
				alias:        alias,
			}
			return this, nil
		} else if v, err := mgr.DeserializeRelationshipActivityStreams()(m, aliasMap); err == nil {
			this := &AttachmentPropertyIterator{
				RelationshipMember: v,
				alias:              alias,
			}
			return this, nil
		} else if v, err := mgr.DeserializeRemoveActivityStreams()(m, aliasMap); err == nil {
			this := &AttachmentPropertyIterator{
				RemoveMember: v,
				alias:        alias,
			}
			return this, nil
		} else if v, err := mgr.DeserializeServiceActivityStreams()(m, aliasMap); err == nil {
			this := &AttachmentPropertyIterator{
				ServiceMember: v,
				alias:         alias,
			}
			return this, nil
		} else if v, err := mgr.DeserializeTentativeAcceptActivityStreams()(m, aliasMap); err == nil {
			this := &AttachmentPropertyIterator{
				TentativeAcceptMember: v,
				alias:                 alias,
			}
			return this, nil
		} else if v, err := mgr.DeserializeTentativeRejectActivityStreams()(m, aliasMap); err == nil {
			this := &AttachmentPropertyIterator{
				TentativeRejectMember: v,
				alias:                 alias,
			}
			return this, nil
		} else if v, err := mgr.DeserializeTombstoneActivityStreams()(m, aliasMap); err == nil {
			this := &AttachmentPropertyIterator{
				TombstoneMember: v,
				alias:           alias,
			}
			return this, nil
		} else if v, err := mgr.DeserializeTravelActivityStreams()(m, aliasMap); err == nil {
			this := &AttachmentPropertyIterator{
				TravelMember: v,
				alias:        alias,
			}
			return this, nil
		} else if v, err := mgr.DeserializeUndoActivityStreams()(m, aliasMap); err == nil {
			this := &AttachmentPropertyIterator{
				UndoMember: v,
				alias:      alias,
			}
			return this, nil
		} else if v, err := mgr.DeserializeUpdateActivityStreams()(m, aliasMap); err == nil {
			this := &AttachmentPropertyIterator{
				UpdateMember: v,
				alias:        alias,
			}
			return this, nil
		} else if v, err := mgr.DeserializeVideoActivityStreams()(m, aliasMap); err == nil {
			this := &AttachmentPropertyIterator{
				VideoMember: v,
				alias:       alias,
			}
			return this, nil
		} else if v, err := mgr.DeserializeViewActivityStreams()(m, aliasMap); err == nil {
			this := &AttachmentPropertyIterator{
				ViewMember: v,
				alias:      alias,
			}
			return this, nil
		}
	} else if str, ok := i.(string); ok {
		this := &AttachmentPropertyIterator{
			alias:   alias,
			unknown: []byte(str),
		}
		return this, nil
	}
	return nil, fmt.Errorf("could not deserialize %q property", "attachment")
}

// GetAccept returns the value of this property. When IsAccept returns false,
// GetAccept will return an arbitrary value.
func (this AttachmentPropertyIterator) GetAccept() vocab.AcceptInterface {
	return this.AcceptMember
}

// GetActivity returns the value of this property. When IsActivity returns false,
// GetActivity will return an arbitrary value.
func (this AttachmentPropertyIterator) GetActivity() vocab.ActivityInterface {
	return this.ActivityMember
}

// GetAdd returns the value of this property. When IsAdd returns false, GetAdd
// will return an arbitrary value.
func (this AttachmentPropertyIterator) GetAdd() vocab.AddInterface {
	return this.AddMember
}

// GetAnnounce returns the value of this property. When IsAnnounce returns false,
// GetAnnounce will return an arbitrary value.
func (this AttachmentPropertyIterator) GetAnnounce() vocab.AnnounceInterface {
	return this.AnnounceMember
}

// GetApplication returns the value of this property. When IsApplication returns
// false, GetApplication will return an arbitrary value.
func (this AttachmentPropertyIterator) GetApplication() vocab.ApplicationInterface {
	return this.ApplicationMember
}

// GetArrive returns the value of this property. When IsArrive returns false,
// GetArrive will return an arbitrary value.
func (this AttachmentPropertyIterator) GetArrive() vocab.ArriveInterface {
	return this.ArriveMember
}

// GetArticle returns the value of this property. When IsArticle returns false,
// GetArticle will return an arbitrary value.
func (this AttachmentPropertyIterator) GetArticle() vocab.ArticleInterface {
	return this.ArticleMember
}

// GetAudio returns the value of this property. When IsAudio returns false,
// GetAudio will return an arbitrary value.
func (this AttachmentPropertyIterator) GetAudio() vocab.AudioInterface {
	return this.AudioMember
}

// GetBlock returns the value of this property. When IsBlock returns false,
// GetBlock will return an arbitrary value.
func (this AttachmentPropertyIterator) GetBlock() vocab.BlockInterface {
	return this.BlockMember
}

// GetCollection returns the value of this property. When IsCollection returns
// false, GetCollection will return an arbitrary value.
func (this AttachmentPropertyIterator) GetCollection() vocab.CollectionInterface {
	return this.CollectionMember
}

// GetCollectionPage returns the value of this property. When IsCollectionPage
// returns false, GetCollectionPage will return an arbitrary value.
func (this AttachmentPropertyIterator) GetCollectionPage() vocab.CollectionPageInterface {
	return this.CollectionPageMember
}

// GetCreate returns the value of this property. When IsCreate returns false,
// GetCreate will return an arbitrary value.
func (this AttachmentPropertyIterator) GetCreate() vocab.CreateInterface {
	return this.CreateMember
}

// GetDelete returns the value of this property. When IsDelete returns false,
// GetDelete will return an arbitrary value.
func (this AttachmentPropertyIterator) GetDelete() vocab.DeleteInterface {
	return this.DeleteMember
}

// GetDislike returns the value of this property. When IsDislike returns false,
// GetDislike will return an arbitrary value.
func (this AttachmentPropertyIterator) GetDislike() vocab.DislikeInterface {
	return this.DislikeMember
}

// GetDocument returns the value of this property. When IsDocument returns false,
// GetDocument will return an arbitrary value.
func (this AttachmentPropertyIterator) GetDocument() vocab.DocumentInterface {
	return this.DocumentMember
}

// GetEvent returns the value of this property. When IsEvent returns false,
// GetEvent will return an arbitrary value.
func (this AttachmentPropertyIterator) GetEvent() vocab.EventInterface {
	return this.EventMember
}

// GetFlag returns the value of this property. When IsFlag returns false, GetFlag
// will return an arbitrary value.
func (this AttachmentPropertyIterator) GetFlag() vocab.FlagInterface {
	return this.FlagMember
}

// GetFollow returns the value of this property. When IsFollow returns false,
// GetFollow will return an arbitrary value.
func (this AttachmentPropertyIterator) GetFollow() vocab.FollowInterface {
	return this.FollowMember
}

// GetGroup returns the value of this property. When IsGroup returns false,
// GetGroup will return an arbitrary value.
func (this AttachmentPropertyIterator) GetGroup() vocab.GroupInterface {
	return this.GroupMember
}

// GetIRI returns the IRI of this property. When IsIRI returns false, GetIRI will
// return an arbitrary value.
func (this AttachmentPropertyIterator) GetIRI() *url.URL {
	return this.iri
}

// GetIgnore returns the value of this property. When IsIgnore returns false,
// GetIgnore will return an arbitrary value.
func (this AttachmentPropertyIterator) GetIgnore() vocab.IgnoreInterface {
	return this.IgnoreMember
}

// GetImage returns the value of this property. When IsImage returns false,
// GetImage will return an arbitrary value.
func (this AttachmentPropertyIterator) GetImage() vocab.ImageInterface {
	return this.ImageMember
}

// GetIntransitiveActivity returns the value of this property. When
// IsIntransitiveActivity returns false, GetIntransitiveActivity will return
// an arbitrary value.
func (this AttachmentPropertyIterator) GetIntransitiveActivity() vocab.IntransitiveActivityInterface {
	return this.IntransitiveActivityMember
}

// GetInvite returns the value of this property. When IsInvite returns false,
// GetInvite will return an arbitrary value.
func (this AttachmentPropertyIterator) GetInvite() vocab.InviteInterface {
	return this.InviteMember
}

// GetJoin returns the value of this property. When IsJoin returns false, GetJoin
// will return an arbitrary value.
func (this AttachmentPropertyIterator) GetJoin() vocab.JoinInterface {
	return this.JoinMember
}

// GetLeave returns the value of this property. When IsLeave returns false,
// GetLeave will return an arbitrary value.
func (this AttachmentPropertyIterator) GetLeave() vocab.LeaveInterface {
	return this.LeaveMember
}

// GetLike returns the value of this property. When IsLike returns false, GetLike
// will return an arbitrary value.
func (this AttachmentPropertyIterator) GetLike() vocab.LikeInterface {
	return this.LikeMember
}

// GetLink returns the value of this property. When IsLink returns false, GetLink
// will return an arbitrary value.
func (this AttachmentPropertyIterator) GetLink() vocab.LinkInterface {
	return this.LinkMember
}

// GetListen returns the value of this property. When IsListen returns false,
// GetListen will return an arbitrary value.
func (this AttachmentPropertyIterator) GetListen() vocab.ListenInterface {
	return this.ListenMember
}

// GetMention returns the value of this property. When IsMention returns false,
// GetMention will return an arbitrary value.
func (this AttachmentPropertyIterator) GetMention() vocab.MentionInterface {
	return this.MentionMember
}

// GetMove returns the value of this property. When IsMove returns false, GetMove
// will return an arbitrary value.
func (this AttachmentPropertyIterator) GetMove() vocab.MoveInterface {
	return this.MoveMember
}

// GetNote returns the value of this property. When IsNote returns false, GetNote
// will return an arbitrary value.
func (this AttachmentPropertyIterator) GetNote() vocab.NoteInterface {
	return this.NoteMember
}

// GetObject returns the value of this property. When IsObject returns false,
// GetObject will return an arbitrary value.
func (this AttachmentPropertyIterator) GetObject() vocab.ObjectInterface {
	return this.ObjectMember
}

// GetOffer returns the value of this property. When IsOffer returns false,
// GetOffer will return an arbitrary value.
func (this AttachmentPropertyIterator) GetOffer() vocab.OfferInterface {
	return this.OfferMember
}

// GetOrderedCollection returns the value of this property. When
// IsOrderedCollection returns false, GetOrderedCollection will return an
// arbitrary value.
func (this AttachmentPropertyIterator) GetOrderedCollection() vocab.OrderedCollectionInterface {
	return this.OrderedCollectionMember
}

// GetOrderedCollectionPage returns the value of this property. When
// IsOrderedCollectionPage returns false, GetOrderedCollectionPage will return
// an arbitrary value.
func (this AttachmentPropertyIterator) GetOrderedCollectionPage() vocab.OrderedCollectionPageInterface {
	return this.OrderedCollectionPageMember
}

// GetOrganization returns the value of this property. When IsOrganization returns
// false, GetOrganization will return an arbitrary value.
func (this AttachmentPropertyIterator) GetOrganization() vocab.OrganizationInterface {
	return this.OrganizationMember
}

// GetPage returns the value of this property. When IsPage returns false, GetPage
// will return an arbitrary value.
func (this AttachmentPropertyIterator) GetPage() vocab.PageInterface {
	return this.PageMember
}

// GetPerson returns the value of this property. When IsPerson returns false,
// GetPerson will return an arbitrary value.
func (this AttachmentPropertyIterator) GetPerson() vocab.PersonInterface {
	return this.PersonMember
}

// GetPlace returns the value of this property. When IsPlace returns false,
// GetPlace will return an arbitrary value.
func (this AttachmentPropertyIterator) GetPlace() vocab.PlaceInterface {
	return this.PlaceMember
}

// GetProfile returns the value of this property. When IsProfile returns false,
// GetProfile will return an arbitrary value.
func (this AttachmentPropertyIterator) GetProfile() vocab.ProfileInterface {
	return this.ProfileMember
}

// GetQuestion returns the value of this property. When IsQuestion returns false,
// GetQuestion will return an arbitrary value.
func (this AttachmentPropertyIterator) GetQuestion() vocab.QuestionInterface {
	return this.QuestionMember
}

// GetRead returns the value of this property. When IsRead returns false, GetRead
// will return an arbitrary value.
func (this AttachmentPropertyIterator) GetRead() vocab.ReadInterface {
	return this.ReadMember
}

// GetReject returns the value of this property. When IsReject returns false,
// GetReject will return an arbitrary value.
func (this AttachmentPropertyIterator) GetReject() vocab.RejectInterface {
	return this.RejectMember
}

// GetRelationship returns the value of this property. When IsRelationship returns
// false, GetRelationship will return an arbitrary value.
func (this AttachmentPropertyIterator) GetRelationship() vocab.RelationshipInterface {
	return this.RelationshipMember
}

// GetRemove returns the value of this property. When IsRemove returns false,
// GetRemove will return an arbitrary value.
func (this AttachmentPropertyIterator) GetRemove() vocab.RemoveInterface {
	return this.RemoveMember
}

// GetService returns the value of this property. When IsService returns false,
// GetService will return an arbitrary value.
func (this AttachmentPropertyIterator) GetService() vocab.ServiceInterface {
	return this.ServiceMember
}

// GetTentativeAccept returns the value of this property. When IsTentativeAccept
// returns false, GetTentativeAccept will return an arbitrary value.
func (this AttachmentPropertyIterator) GetTentativeAccept() vocab.TentativeAcceptInterface {
	return this.TentativeAcceptMember
}

// GetTentativeReject returns the value of this property. When IsTentativeReject
// returns false, GetTentativeReject will return an arbitrary value.
func (this AttachmentPropertyIterator) GetTentativeReject() vocab.TentativeRejectInterface {
	return this.TentativeRejectMember
}

// GetTombstone returns the value of this property. When IsTombstone returns
// false, GetTombstone will return an arbitrary value.
func (this AttachmentPropertyIterator) GetTombstone() vocab.TombstoneInterface {
	return this.TombstoneMember
}

// GetTravel returns the value of this property. When IsTravel returns false,
// GetTravel will return an arbitrary value.
func (this AttachmentPropertyIterator) GetTravel() vocab.TravelInterface {
	return this.TravelMember
}

// GetUndo returns the value of this property. When IsUndo returns false, GetUndo
// will return an arbitrary value.
func (this AttachmentPropertyIterator) GetUndo() vocab.UndoInterface {
	return this.UndoMember
}

// GetUpdate returns the value of this property. When IsUpdate returns false,
// GetUpdate will return an arbitrary value.
func (this AttachmentPropertyIterator) GetUpdate() vocab.UpdateInterface {
	return this.UpdateMember
}

// GetVideo returns the value of this property. When IsVideo returns false,
// GetVideo will return an arbitrary value.
func (this AttachmentPropertyIterator) GetVideo() vocab.VideoInterface {
	return this.VideoMember
}

// GetView returns the value of this property. When IsView returns false, GetView
// will return an arbitrary value.
func (this AttachmentPropertyIterator) GetView() vocab.ViewInterface {
	return this.ViewMember
}

// HasAny returns true if any of the different values is set.
func (this AttachmentPropertyIterator) HasAny() bool {
	return this.IsObject() ||
		this.IsLink() ||
		this.IsAccept() ||
		this.IsActivity() ||
		this.IsAdd() ||
		this.IsAnnounce() ||
		this.IsApplication() ||
		this.IsArrive() ||
		this.IsArticle() ||
		this.IsAudio() ||
		this.IsBlock() ||
		this.IsCollection() ||
		this.IsCollectionPage() ||
		this.IsCreate() ||
		this.IsDelete() ||
		this.IsDislike() ||
		this.IsDocument() ||
		this.IsEvent() ||
		this.IsFlag() ||
		this.IsFollow() ||
		this.IsGroup() ||
		this.IsIgnore() ||
		this.IsImage() ||
		this.IsIntransitiveActivity() ||
		this.IsInvite() ||
		this.IsJoin() ||
		this.IsLeave() ||
		this.IsLike() ||
		this.IsListen() ||
		this.IsMention() ||
		this.IsMove() ||
		this.IsNote() ||
		this.IsOffer() ||
		this.IsOrderedCollection() ||
		this.IsOrderedCollectionPage() ||
		this.IsOrganization() ||
		this.IsPage() ||
		this.IsPerson() ||
		this.IsPlace() ||
		this.IsProfile() ||
		this.IsQuestion() ||
		this.IsRead() ||
		this.IsReject() ||
		this.IsRelationship() ||
		this.IsRemove() ||
		this.IsService() ||
		this.IsTentativeAccept() ||
		this.IsTentativeReject() ||
		this.IsTombstone() ||
		this.IsTravel() ||
		this.IsUndo() ||
		this.IsUpdate() ||
		this.IsVideo() ||
		this.IsView() ||
		this.iri != nil
}

// IsAccept returns true if this property has a type of "Accept". When true, use
// the GetAccept and SetAccept methods to access and set this property.
func (this AttachmentPropertyIterator) IsAccept() bool {
	return this.AcceptMember != nil
}

// IsActivity returns true if this property has a type of "Activity". When true,
// use the GetActivity and SetActivity methods to access and set this property.
func (this AttachmentPropertyIterator) IsActivity() bool {
	return this.ActivityMember != nil
}

// IsAdd returns true if this property has a type of "Add". When true, use the
// GetAdd and SetAdd methods to access and set this property.
func (this AttachmentPropertyIterator) IsAdd() bool {
	return this.AddMember != nil
}

// IsAnnounce returns true if this property has a type of "Announce". When true,
// use the GetAnnounce and SetAnnounce methods to access and set this property.
func (this AttachmentPropertyIterator) IsAnnounce() bool {
	return this.AnnounceMember != nil
}

// IsApplication returns true if this property has a type of "Application". When
// true, use the GetApplication and SetApplication methods to access and set
// this property.
func (this AttachmentPropertyIterator) IsApplication() bool {
	return this.ApplicationMember != nil
}

// IsArrive returns true if this property has a type of "Arrive". When true, use
// the GetArrive and SetArrive methods to access and set this property.
func (this AttachmentPropertyIterator) IsArrive() bool {
	return this.ArriveMember != nil
}

// IsArticle returns true if this property has a type of "Article". When true, use
// the GetArticle and SetArticle methods to access and set this property.
func (this AttachmentPropertyIterator) IsArticle() bool {
	return this.ArticleMember != nil
}

// IsAudio returns true if this property has a type of "Audio". When true, use the
// GetAudio and SetAudio methods to access and set this property.
func (this AttachmentPropertyIterator) IsAudio() bool {
	return this.AudioMember != nil
}

// IsBlock returns true if this property has a type of "Block". When true, use the
// GetBlock and SetBlock methods to access and set this property.
func (this AttachmentPropertyIterator) IsBlock() bool {
	return this.BlockMember != nil
}

// IsCollection returns true if this property has a type of "Collection". When
// true, use the GetCollection and SetCollection methods to access and set
// this property.
func (this AttachmentPropertyIterator) IsCollection() bool {
	return this.CollectionMember != nil
}

// IsCollectionPage returns true if this property has a type of "CollectionPage".
// When true, use the GetCollectionPage and SetCollectionPage methods to
// access and set this property.
func (this AttachmentPropertyIterator) IsCollectionPage() bool {
	return this.CollectionPageMember != nil
}

// IsCreate returns true if this property has a type of "Create". When true, use
// the GetCreate and SetCreate methods to access and set this property.
func (this AttachmentPropertyIterator) IsCreate() bool {
	return this.CreateMember != nil
}

// IsDelete returns true if this property has a type of "Delete". When true, use
// the GetDelete and SetDelete methods to access and set this property.
func (this AttachmentPropertyIterator) IsDelete() bool {
	return this.DeleteMember != nil
}

// IsDislike returns true if this property has a type of "Dislike". When true, use
// the GetDislike and SetDislike methods to access and set this property.
func (this AttachmentPropertyIterator) IsDislike() bool {
	return this.DislikeMember != nil
}

// IsDocument returns true if this property has a type of "Document". When true,
// use the GetDocument and SetDocument methods to access and set this property.
func (this AttachmentPropertyIterator) IsDocument() bool {
	return this.DocumentMember != nil
}

// IsEvent returns true if this property has a type of "Event". When true, use the
// GetEvent and SetEvent methods to access and set this property.
func (this AttachmentPropertyIterator) IsEvent() bool {
	return this.EventMember != nil
}

// IsFlag returns true if this property has a type of "Flag". When true, use the
// GetFlag and SetFlag methods to access and set this property.
func (this AttachmentPropertyIterator) IsFlag() bool {
	return this.FlagMember != nil
}

// IsFollow returns true if this property has a type of "Follow". When true, use
// the GetFollow and SetFollow methods to access and set this property.
func (this AttachmentPropertyIterator) IsFollow() bool {
	return this.FollowMember != nil
}

// IsGroup returns true if this property has a type of "Group". When true, use the
// GetGroup and SetGroup methods to access and set this property.
func (this AttachmentPropertyIterator) IsGroup() bool {
	return this.GroupMember != nil
}

// IsIRI returns true if this property is an IRI. When true, use GetIRI and SetIRI
// to access and set this property
func (this AttachmentPropertyIterator) IsIRI() bool {
	return this.iri != nil
}

// IsIgnore returns true if this property has a type of "Ignore". When true, use
// the GetIgnore and SetIgnore methods to access and set this property.
func (this AttachmentPropertyIterator) IsIgnore() bool {
	return this.IgnoreMember != nil
}

// IsImage returns true if this property has a type of "Image". When true, use the
// GetImage and SetImage methods to access and set this property.
func (this AttachmentPropertyIterator) IsImage() bool {
	return this.ImageMember != nil
}

// IsIntransitiveActivity returns true if this property has a type of
// "IntransitiveActivity". When true, use the GetIntransitiveActivity and
// SetIntransitiveActivity methods to access and set this property.
func (this AttachmentPropertyIterator) IsIntransitiveActivity() bool {
	return this.IntransitiveActivityMember != nil
}

// IsInvite returns true if this property has a type of "Invite". When true, use
// the GetInvite and SetInvite methods to access and set this property.
func (this AttachmentPropertyIterator) IsInvite() bool {
	return this.InviteMember != nil
}

// IsJoin returns true if this property has a type of "Join". When true, use the
// GetJoin and SetJoin methods to access and set this property.
func (this AttachmentPropertyIterator) IsJoin() bool {
	return this.JoinMember != nil
}

// IsLeave returns true if this property has a type of "Leave". When true, use the
// GetLeave and SetLeave methods to access and set this property.
func (this AttachmentPropertyIterator) IsLeave() bool {
	return this.LeaveMember != nil
}

// IsLike returns true if this property has a type of "Like". When true, use the
// GetLike and SetLike methods to access and set this property.
func (this AttachmentPropertyIterator) IsLike() bool {
	return this.LikeMember != nil
}

// IsLink returns true if this property has a type of "Link". When true, use the
// GetLink and SetLink methods to access and set this property.
func (this AttachmentPropertyIterator) IsLink() bool {
	return this.LinkMember != nil
}

// IsListen returns true if this property has a type of "Listen". When true, use
// the GetListen and SetListen methods to access and set this property.
func (this AttachmentPropertyIterator) IsListen() bool {
	return this.ListenMember != nil
}

// IsMention returns true if this property has a type of "Mention". When true, use
// the GetMention and SetMention methods to access and set this property.
func (this AttachmentPropertyIterator) IsMention() bool {
	return this.MentionMember != nil
}

// IsMove returns true if this property has a type of "Move". When true, use the
// GetMove and SetMove methods to access and set this property.
func (this AttachmentPropertyIterator) IsMove() bool {
	return this.MoveMember != nil
}

// IsNote returns true if this property has a type of "Note". When true, use the
// GetNote and SetNote methods to access and set this property.
func (this AttachmentPropertyIterator) IsNote() bool {
	return this.NoteMember != nil
}

// IsObject returns true if this property has a type of "Object". When true, use
// the GetObject and SetObject methods to access and set this property.
func (this AttachmentPropertyIterator) IsObject() bool {
	return this.ObjectMember != nil
}

// IsOffer returns true if this property has a type of "Offer". When true, use the
// GetOffer and SetOffer methods to access and set this property.
func (this AttachmentPropertyIterator) IsOffer() bool {
	return this.OfferMember != nil
}

// IsOrderedCollection returns true if this property has a type of
// "OrderedCollection". When true, use the GetOrderedCollection and
// SetOrderedCollection methods to access and set this property.
func (this AttachmentPropertyIterator) IsOrderedCollection() bool {
	return this.OrderedCollectionMember != nil
}

// IsOrderedCollectionPage returns true if this property has a type of
// "OrderedCollectionPage". When true, use the GetOrderedCollectionPage and
// SetOrderedCollectionPage methods to access and set this property.
func (this AttachmentPropertyIterator) IsOrderedCollectionPage() bool {
	return this.OrderedCollectionPageMember != nil
}

// IsOrganization returns true if this property has a type of "Organization". When
// true, use the GetOrganization and SetOrganization methods to access and set
// this property.
func (this AttachmentPropertyIterator) IsOrganization() bool {
	return this.OrganizationMember != nil
}

// IsPage returns true if this property has a type of "Page". When true, use the
// GetPage and SetPage methods to access and set this property.
func (this AttachmentPropertyIterator) IsPage() bool {
	return this.PageMember != nil
}

// IsPerson returns true if this property has a type of "Person". When true, use
// the GetPerson and SetPerson methods to access and set this property.
func (this AttachmentPropertyIterator) IsPerson() bool {
	return this.PersonMember != nil
}

// IsPlace returns true if this property has a type of "Place". When true, use the
// GetPlace and SetPlace methods to access and set this property.
func (this AttachmentPropertyIterator) IsPlace() bool {
	return this.PlaceMember != nil
}

// IsProfile returns true if this property has a type of "Profile". When true, use
// the GetProfile and SetProfile methods to access and set this property.
func (this AttachmentPropertyIterator) IsProfile() bool {
	return this.ProfileMember != nil
}

// IsQuestion returns true if this property has a type of "Question". When true,
// use the GetQuestion and SetQuestion methods to access and set this property.
func (this AttachmentPropertyIterator) IsQuestion() bool {
	return this.QuestionMember != nil
}

// IsRead returns true if this property has a type of "Read". When true, use the
// GetRead and SetRead methods to access and set this property.
func (this AttachmentPropertyIterator) IsRead() bool {
	return this.ReadMember != nil
}

// IsReject returns true if this property has a type of "Reject". When true, use
// the GetReject and SetReject methods to access and set this property.
func (this AttachmentPropertyIterator) IsReject() bool {
	return this.RejectMember != nil
}

// IsRelationship returns true if this property has a type of "Relationship". When
// true, use the GetRelationship and SetRelationship methods to access and set
// this property.
func (this AttachmentPropertyIterator) IsRelationship() bool {
	return this.RelationshipMember != nil
}

// IsRemove returns true if this property has a type of "Remove". When true, use
// the GetRemove and SetRemove methods to access and set this property.
func (this AttachmentPropertyIterator) IsRemove() bool {
	return this.RemoveMember != nil
}

// IsService returns true if this property has a type of "Service". When true, use
// the GetService and SetService methods to access and set this property.
func (this AttachmentPropertyIterator) IsService() bool {
	return this.ServiceMember != nil
}

// IsTentativeAccept returns true if this property has a type of
// "TentativeAccept". When true, use the GetTentativeAccept and
// SetTentativeAccept methods to access and set this property.
func (this AttachmentPropertyIterator) IsTentativeAccept() bool {
	return this.TentativeAcceptMember != nil
}

// IsTentativeReject returns true if this property has a type of
// "TentativeReject". When true, use the GetTentativeReject and
// SetTentativeReject methods to access and set this property.
func (this AttachmentPropertyIterator) IsTentativeReject() bool {
	return this.TentativeRejectMember != nil
}

// IsTombstone returns true if this property has a type of "Tombstone". When true,
// use the GetTombstone and SetTombstone methods to access and set this
// property.
func (this AttachmentPropertyIterator) IsTombstone() bool {
	return this.TombstoneMember != nil
}

// IsTravel returns true if this property has a type of "Travel". When true, use
// the GetTravel and SetTravel methods to access and set this property.
func (this AttachmentPropertyIterator) IsTravel() bool {
	return this.TravelMember != nil
}

// IsUndo returns true if this property has a type of "Undo". When true, use the
// GetUndo and SetUndo methods to access and set this property.
func (this AttachmentPropertyIterator) IsUndo() bool {
	return this.UndoMember != nil
}

// IsUpdate returns true if this property has a type of "Update". When true, use
// the GetUpdate and SetUpdate methods to access and set this property.
func (this AttachmentPropertyIterator) IsUpdate() bool {
	return this.UpdateMember != nil
}

// IsVideo returns true if this property has a type of "Video". When true, use the
// GetVideo and SetVideo methods to access and set this property.
func (this AttachmentPropertyIterator) IsVideo() bool {
	return this.VideoMember != nil
}

// IsView returns true if this property has a type of "View". When true, use the
// GetView and SetView methods to access and set this property.
func (this AttachmentPropertyIterator) IsView() bool {
	return this.ViewMember != nil
}

// JSONLDContext returns the JSONLD URIs required in the context string for this
// property and the specific values that are set. The value in the map is the
// alias used to import the property's value or values.
func (this AttachmentPropertyIterator) JSONLDContext() map[string]string {
	m := map[string]string{"https://www.w3.org/TR/activitystreams-vocabulary": this.alias}
	var child map[string]string
	if this.IsObject() {
		child = this.GetObject().JSONLDContext()
	} else if this.IsLink() {
		child = this.GetLink().JSONLDContext()
	} else if this.IsAccept() {
		child = this.GetAccept().JSONLDContext()
	} else if this.IsActivity() {
		child = this.GetActivity().JSONLDContext()
	} else if this.IsAdd() {
		child = this.GetAdd().JSONLDContext()
	} else if this.IsAnnounce() {
		child = this.GetAnnounce().JSONLDContext()
	} else if this.IsApplication() {
		child = this.GetApplication().JSONLDContext()
	} else if this.IsArrive() {
		child = this.GetArrive().JSONLDContext()
	} else if this.IsArticle() {
		child = this.GetArticle().JSONLDContext()
	} else if this.IsAudio() {
		child = this.GetAudio().JSONLDContext()
	} else if this.IsBlock() {
		child = this.GetBlock().JSONLDContext()
	} else if this.IsCollection() {
		child = this.GetCollection().JSONLDContext()
	} else if this.IsCollectionPage() {
		child = this.GetCollectionPage().JSONLDContext()
	} else if this.IsCreate() {
		child = this.GetCreate().JSONLDContext()
	} else if this.IsDelete() {
		child = this.GetDelete().JSONLDContext()
	} else if this.IsDislike() {
		child = this.GetDislike().JSONLDContext()
	} else if this.IsDocument() {
		child = this.GetDocument().JSONLDContext()
	} else if this.IsEvent() {
		child = this.GetEvent().JSONLDContext()
	} else if this.IsFlag() {
		child = this.GetFlag().JSONLDContext()
	} else if this.IsFollow() {
		child = this.GetFollow().JSONLDContext()
	} else if this.IsGroup() {
		child = this.GetGroup().JSONLDContext()
	} else if this.IsIgnore() {
		child = this.GetIgnore().JSONLDContext()
	} else if this.IsImage() {
		child = this.GetImage().JSONLDContext()
	} else if this.IsIntransitiveActivity() {
		child = this.GetIntransitiveActivity().JSONLDContext()
	} else if this.IsInvite() {
		child = this.GetInvite().JSONLDContext()
	} else if this.IsJoin() {
		child = this.GetJoin().JSONLDContext()
	} else if this.IsLeave() {
		child = this.GetLeave().JSONLDContext()
	} else if this.IsLike() {
		child = this.GetLike().JSONLDContext()
	} else if this.IsListen() {
		child = this.GetListen().JSONLDContext()
	} else if this.IsMention() {
		child = this.GetMention().JSONLDContext()
	} else if this.IsMove() {
		child = this.GetMove().JSONLDContext()
	} else if this.IsNote() {
		child = this.GetNote().JSONLDContext()
	} else if this.IsOffer() {
		child = this.GetOffer().JSONLDContext()
	} else if this.IsOrderedCollection() {
		child = this.GetOrderedCollection().JSONLDContext()
	} else if this.IsOrderedCollectionPage() {
		child = this.GetOrderedCollectionPage().JSONLDContext()
	} else if this.IsOrganization() {
		child = this.GetOrganization().JSONLDContext()
	} else if this.IsPage() {
		child = this.GetPage().JSONLDContext()
	} else if this.IsPerson() {
		child = this.GetPerson().JSONLDContext()
	} else if this.IsPlace() {
		child = this.GetPlace().JSONLDContext()
	} else if this.IsProfile() {
		child = this.GetProfile().JSONLDContext()
	} else if this.IsQuestion() {
		child = this.GetQuestion().JSONLDContext()
	} else if this.IsRead() {
		child = this.GetRead().JSONLDContext()
	} else if this.IsReject() {
		child = this.GetReject().JSONLDContext()
	} else if this.IsRelationship() {
		child = this.GetRelationship().JSONLDContext()
	} else if this.IsRemove() {
		child = this.GetRemove().JSONLDContext()
	} else if this.IsService() {
		child = this.GetService().JSONLDContext()
	} else if this.IsTentativeAccept() {
		child = this.GetTentativeAccept().JSONLDContext()
	} else if this.IsTentativeReject() {
		child = this.GetTentativeReject().JSONLDContext()
	} else if this.IsTombstone() {
		child = this.GetTombstone().JSONLDContext()
	} else if this.IsTravel() {
		child = this.GetTravel().JSONLDContext()
	} else if this.IsUndo() {
		child = this.GetUndo().JSONLDContext()
	} else if this.IsUpdate() {
		child = this.GetUpdate().JSONLDContext()
	} else if this.IsVideo() {
		child = this.GetVideo().JSONLDContext()
	} else if this.IsView() {
		child = this.GetView().JSONLDContext()
	}
	/*
	   Since the literal maps in this function are determined at
	   code-generation time, this loop should not overwrite an existing key with a
	   new value.
	*/
	for k, v := range child {
		m[k] = v
	}
	return m
}

// KindIndex computes an arbitrary value for indexing this kind of value. This is
// a leaky API detail only for folks looking to replace the go-fed
// implementation. Applications should not use this method.
func (this AttachmentPropertyIterator) KindIndex() int {
	if this.IsObject() {
		return 0
	}
	if this.IsLink() {
		return 1
	}
	if this.IsAccept() {
		return 2
	}
	if this.IsActivity() {
		return 3
	}
	if this.IsAdd() {
		return 4
	}
	if this.IsAnnounce() {
		return 5
	}
	if this.IsApplication() {
		return 6
	}
	if this.IsArrive() {
		return 7
	}
	if this.IsArticle() {
		return 8
	}
	if this.IsAudio() {
		return 9
	}
	if this.IsBlock() {
		return 10
	}
	if this.IsCollection() {
		return 11
	}
	if this.IsCollectionPage() {
		return 12
	}
	if this.IsCreate() {
		return 13
	}
	if this.IsDelete() {
		return 14
	}
	if this.IsDislike() {
		return 15
	}
	if this.IsDocument() {
		return 16
	}
	if this.IsEvent() {
		return 17
	}
	if this.IsFlag() {
		return 18
	}
	if this.IsFollow() {
		return 19
	}
	if this.IsGroup() {
		return 20
	}
	if this.IsIgnore() {
		return 21
	}
	if this.IsImage() {
		return 22
	}
	if this.IsIntransitiveActivity() {
		return 23
	}
	if this.IsInvite() {
		return 24
	}
	if this.IsJoin() {
		return 25
	}
	if this.IsLeave() {
		return 26
	}
	if this.IsLike() {
		return 27
	}
	if this.IsListen() {
		return 28
	}
	if this.IsMention() {
		return 29
	}
	if this.IsMove() {
		return 30
	}
	if this.IsNote() {
		return 31
	}
	if this.IsOffer() {
		return 32
	}
	if this.IsOrderedCollection() {
		return 33
	}
	if this.IsOrderedCollectionPage() {
		return 34
	}
	if this.IsOrganization() {
		return 35
	}
	if this.IsPage() {
		return 36
	}
	if this.IsPerson() {
		return 37
	}
	if this.IsPlace() {
		return 38
	}
	if this.IsProfile() {
		return 39
	}
	if this.IsQuestion() {
		return 40
	}
	if this.IsRead() {
		return 41
	}
	if this.IsReject() {
		return 42
	}
	if this.IsRelationship() {
		return 43
	}
	if this.IsRemove() {
		return 44
	}
	if this.IsService() {
		return 45
	}
	if this.IsTentativeAccept() {
		return 46
	}
	if this.IsTentativeReject() {
		return 47
	}
	if this.IsTombstone() {
		return 48
	}
	if this.IsTravel() {
		return 49
	}
	if this.IsUndo() {
		return 50
	}
	if this.IsUpdate() {
		return 51
	}
	if this.IsVideo() {
		return 52
	}
	if this.IsView() {
		return 53
	}
	if this.IsIRI() {
		return -2
	}
	return -1
}

// LessThan compares two instances of this property with an arbitrary but stable
// comparison. Applications should not use this because it is only meant to
// help alternative implementations to go-fed to be able to normalize
// nonfunctional properties.
func (this AttachmentPropertyIterator) LessThan(o vocab.AttachmentPropertyIteratorInterface) bool {
	idx1 := this.KindIndex()
	idx2 := o.KindIndex()
	if idx1 < idx2 {
		return true
	} else if idx1 > idx2 {
		return false
	} else if this.IsObject() {
		return this.GetObject().LessThan(o.GetObject())
	} else if this.IsLink() {
		return this.GetLink().LessThan(o.GetLink())
	} else if this.IsAccept() {
		return this.GetAccept().LessThan(o.GetAccept())
	} else if this.IsActivity() {
		return this.GetActivity().LessThan(o.GetActivity())
	} else if this.IsAdd() {
		return this.GetAdd().LessThan(o.GetAdd())
	} else if this.IsAnnounce() {
		return this.GetAnnounce().LessThan(o.GetAnnounce())
	} else if this.IsApplication() {
		return this.GetApplication().LessThan(o.GetApplication())
	} else if this.IsArrive() {
		return this.GetArrive().LessThan(o.GetArrive())
	} else if this.IsArticle() {
		return this.GetArticle().LessThan(o.GetArticle())
	} else if this.IsAudio() {
		return this.GetAudio().LessThan(o.GetAudio())
	} else if this.IsBlock() {
		return this.GetBlock().LessThan(o.GetBlock())
	} else if this.IsCollection() {
		return this.GetCollection().LessThan(o.GetCollection())
	} else if this.IsCollectionPage() {
		return this.GetCollectionPage().LessThan(o.GetCollectionPage())
	} else if this.IsCreate() {
		return this.GetCreate().LessThan(o.GetCreate())
	} else if this.IsDelete() {
		return this.GetDelete().LessThan(o.GetDelete())
	} else if this.IsDislike() {
		return this.GetDislike().LessThan(o.GetDislike())
	} else if this.IsDocument() {
		return this.GetDocument().LessThan(o.GetDocument())
	} else if this.IsEvent() {
		return this.GetEvent().LessThan(o.GetEvent())
	} else if this.IsFlag() {
		return this.GetFlag().LessThan(o.GetFlag())
	} else if this.IsFollow() {
		return this.GetFollow().LessThan(o.GetFollow())
	} else if this.IsGroup() {
		return this.GetGroup().LessThan(o.GetGroup())
	} else if this.IsIgnore() {
		return this.GetIgnore().LessThan(o.GetIgnore())
	} else if this.IsImage() {
		return this.GetImage().LessThan(o.GetImage())
	} else if this.IsIntransitiveActivity() {
		return this.GetIntransitiveActivity().LessThan(o.GetIntransitiveActivity())
	} else if this.IsInvite() {
		return this.GetInvite().LessThan(o.GetInvite())
	} else if this.IsJoin() {
		return this.GetJoin().LessThan(o.GetJoin())
	} else if this.IsLeave() {
		return this.GetLeave().LessThan(o.GetLeave())
	} else if this.IsLike() {
		return this.GetLike().LessThan(o.GetLike())
	} else if this.IsListen() {
		return this.GetListen().LessThan(o.GetListen())
	} else if this.IsMention() {
		return this.GetMention().LessThan(o.GetMention())
	} else if this.IsMove() {
		return this.GetMove().LessThan(o.GetMove())
	} else if this.IsNote() {
		return this.GetNote().LessThan(o.GetNote())
	} else if this.IsOffer() {
		return this.GetOffer().LessThan(o.GetOffer())
	} else if this.IsOrderedCollection() {
		return this.GetOrderedCollection().LessThan(o.GetOrderedCollection())
	} else if this.IsOrderedCollectionPage() {
		return this.GetOrderedCollectionPage().LessThan(o.GetOrderedCollectionPage())
	} else if this.IsOrganization() {
		return this.GetOrganization().LessThan(o.GetOrganization())
	} else if this.IsPage() {
		return this.GetPage().LessThan(o.GetPage())
	} else if this.IsPerson() {
		return this.GetPerson().LessThan(o.GetPerson())
	} else if this.IsPlace() {
		return this.GetPlace().LessThan(o.GetPlace())
	} else if this.IsProfile() {
		return this.GetProfile().LessThan(o.GetProfile())
	} else if this.IsQuestion() {
		return this.GetQuestion().LessThan(o.GetQuestion())
	} else if this.IsRead() {
		return this.GetRead().LessThan(o.GetRead())
	} else if this.IsReject() {
		return this.GetReject().LessThan(o.GetReject())
	} else if this.IsRelationship() {
		return this.GetRelationship().LessThan(o.GetRelationship())
	} else if this.IsRemove() {
		return this.GetRemove().LessThan(o.GetRemove())
	} else if this.IsService() {
		return this.GetService().LessThan(o.GetService())
	} else if this.IsTentativeAccept() {
		return this.GetTentativeAccept().LessThan(o.GetTentativeAccept())
	} else if this.IsTentativeReject() {
		return this.GetTentativeReject().LessThan(o.GetTentativeReject())
	} else if this.IsTombstone() {
		return this.GetTombstone().LessThan(o.GetTombstone())
	} else if this.IsTravel() {
		return this.GetTravel().LessThan(o.GetTravel())
	} else if this.IsUndo() {
		return this.GetUndo().LessThan(o.GetUndo())
	} else if this.IsUpdate() {
		return this.GetUpdate().LessThan(o.GetUpdate())
	} else if this.IsVideo() {
		return this.GetVideo().LessThan(o.GetVideo())
	} else if this.IsView() {
		return this.GetView().LessThan(o.GetView())
	} else if this.IsIRI() {
		return this.iri.String() < o.GetIRI().String()
	}
	return false
}

// Name returns the name of this property: "attachment".
func (this AttachmentPropertyIterator) Name() string {
	return "attachment"
}

// Next returns the next iterator, or nil if there is no next iterator.
func (this AttachmentPropertyIterator) Next() vocab.AttachmentPropertyIteratorInterface {
	if this.myIdx+1 >= this.parent.Len() {
		return nil
	} else {
		return this.parent.At(this.myIdx + 1)
	}
}

// Prev returns the previous iterator, or nil if there is no previous iterator.
func (this AttachmentPropertyIterator) Prev() vocab.AttachmentPropertyIteratorInterface {
	if this.myIdx-1 < 0 {
		return nil
	} else {
		return this.parent.At(this.myIdx - 1)
	}
}

// SetAccept sets the value of this property. Calling IsAccept afterwards returns
// true.
func (this *AttachmentPropertyIterator) SetAccept(v vocab.AcceptInterface) {
	this.clear()
	this.AcceptMember = v
}

// SetActivity sets the value of this property. Calling IsActivity afterwards
// returns true.
func (this *AttachmentPropertyIterator) SetActivity(v vocab.ActivityInterface) {
	this.clear()
	this.ActivityMember = v
}

// SetAdd sets the value of this property. Calling IsAdd afterwards returns true.
func (this *AttachmentPropertyIterator) SetAdd(v vocab.AddInterface) {
	this.clear()
	this.AddMember = v
}

// SetAnnounce sets the value of this property. Calling IsAnnounce afterwards
// returns true.
func (this *AttachmentPropertyIterator) SetAnnounce(v vocab.AnnounceInterface) {
	this.clear()
	this.AnnounceMember = v
}

// SetApplication sets the value of this property. Calling IsApplication
// afterwards returns true.
func (this *AttachmentPropertyIterator) SetApplication(v vocab.ApplicationInterface) {
	this.clear()
	this.ApplicationMember = v
}

// SetArrive sets the value of this property. Calling IsArrive afterwards returns
// true.
func (this *AttachmentPropertyIterator) SetArrive(v vocab.ArriveInterface) {
	this.clear()
	this.ArriveMember = v
}

// SetArticle sets the value of this property. Calling IsArticle afterwards
// returns true.
func (this *AttachmentPropertyIterator) SetArticle(v vocab.ArticleInterface) {
	this.clear()
	this.ArticleMember = v
}

// SetAudio sets the value of this property. Calling IsAudio afterwards returns
// true.
func (this *AttachmentPropertyIterator) SetAudio(v vocab.AudioInterface) {
	this.clear()
	this.AudioMember = v
}

// SetBlock sets the value of this property. Calling IsBlock afterwards returns
// true.
func (this *AttachmentPropertyIterator) SetBlock(v vocab.BlockInterface) {
	this.clear()
	this.BlockMember = v
}

// SetCollection sets the value of this property. Calling IsCollection afterwards
// returns true.
func (this *AttachmentPropertyIterator) SetCollection(v vocab.CollectionInterface) {
	this.clear()
	this.CollectionMember = v
}

// SetCollectionPage sets the value of this property. Calling IsCollectionPage
// afterwards returns true.
func (this *AttachmentPropertyIterator) SetCollectionPage(v vocab.CollectionPageInterface) {
	this.clear()
	this.CollectionPageMember = v
}

// SetCreate sets the value of this property. Calling IsCreate afterwards returns
// true.
func (this *AttachmentPropertyIterator) SetCreate(v vocab.CreateInterface) {
	this.clear()
	this.CreateMember = v
}

// SetDelete sets the value of this property. Calling IsDelete afterwards returns
// true.
func (this *AttachmentPropertyIterator) SetDelete(v vocab.DeleteInterface) {
	this.clear()
	this.DeleteMember = v
}

// SetDislike sets the value of this property. Calling IsDislike afterwards
// returns true.
func (this *AttachmentPropertyIterator) SetDislike(v vocab.DislikeInterface) {
	this.clear()
	this.DislikeMember = v
}

// SetDocument sets the value of this property. Calling IsDocument afterwards
// returns true.
func (this *AttachmentPropertyIterator) SetDocument(v vocab.DocumentInterface) {
	this.clear()
	this.DocumentMember = v
}

// SetEvent sets the value of this property. Calling IsEvent afterwards returns
// true.
func (this *AttachmentPropertyIterator) SetEvent(v vocab.EventInterface) {
	this.clear()
	this.EventMember = v
}

// SetFlag sets the value of this property. Calling IsFlag afterwards returns true.
func (this *AttachmentPropertyIterator) SetFlag(v vocab.FlagInterface) {
	this.clear()
	this.FlagMember = v
}

// SetFollow sets the value of this property. Calling IsFollow afterwards returns
// true.
func (this *AttachmentPropertyIterator) SetFollow(v vocab.FollowInterface) {
	this.clear()
	this.FollowMember = v
}

// SetGroup sets the value of this property. Calling IsGroup afterwards returns
// true.
func (this *AttachmentPropertyIterator) SetGroup(v vocab.GroupInterface) {
	this.clear()
	this.GroupMember = v
}

// SetIRI sets the value of this property. Calling IsIRI afterwards returns true.
func (this *AttachmentPropertyIterator) SetIRI(v *url.URL) {
	this.clear()
	this.iri = v
}

// SetIgnore sets the value of this property. Calling IsIgnore afterwards returns
// true.
func (this *AttachmentPropertyIterator) SetIgnore(v vocab.IgnoreInterface) {
	this.clear()
	this.IgnoreMember = v
}

// SetImage sets the value of this property. Calling IsImage afterwards returns
// true.
func (this *AttachmentPropertyIterator) SetImage(v vocab.ImageInterface) {
	this.clear()
	this.ImageMember = v
}

// SetIntransitiveActivity sets the value of this property. Calling
// IsIntransitiveActivity afterwards returns true.
func (this *AttachmentPropertyIterator) SetIntransitiveActivity(v vocab.IntransitiveActivityInterface) {
	this.clear()
	this.IntransitiveActivityMember = v
}

// SetInvite sets the value of this property. Calling IsInvite afterwards returns
// true.
func (this *AttachmentPropertyIterator) SetInvite(v vocab.InviteInterface) {
	this.clear()
	this.InviteMember = v
}

// SetJoin sets the value of this property. Calling IsJoin afterwards returns true.
func (this *AttachmentPropertyIterator) SetJoin(v vocab.JoinInterface) {
	this.clear()
	this.JoinMember = v
}

// SetLeave sets the value of this property. Calling IsLeave afterwards returns
// true.
func (this *AttachmentPropertyIterator) SetLeave(v vocab.LeaveInterface) {
	this.clear()
	this.LeaveMember = v
}

// SetLike sets the value of this property. Calling IsLike afterwards returns true.
func (this *AttachmentPropertyIterator) SetLike(v vocab.LikeInterface) {
	this.clear()
	this.LikeMember = v
}

// SetLink sets the value of this property. Calling IsLink afterwards returns true.
func (this *AttachmentPropertyIterator) SetLink(v vocab.LinkInterface) {
	this.clear()
	this.LinkMember = v
}

// SetListen sets the value of this property. Calling IsListen afterwards returns
// true.
func (this *AttachmentPropertyIterator) SetListen(v vocab.ListenInterface) {
	this.clear()
	this.ListenMember = v
}

// SetMention sets the value of this property. Calling IsMention afterwards
// returns true.
func (this *AttachmentPropertyIterator) SetMention(v vocab.MentionInterface) {
	this.clear()
	this.MentionMember = v
}

// SetMove sets the value of this property. Calling IsMove afterwards returns true.
func (this *AttachmentPropertyIterator) SetMove(v vocab.MoveInterface) {
	this.clear()
	this.MoveMember = v
}

// SetNote sets the value of this property. Calling IsNote afterwards returns true.
func (this *AttachmentPropertyIterator) SetNote(v vocab.NoteInterface) {
	this.clear()
	this.NoteMember = v
}

// SetObject sets the value of this property. Calling IsObject afterwards returns
// true.
func (this *AttachmentPropertyIterator) SetObject(v vocab.ObjectInterface) {
	this.clear()
	this.ObjectMember = v
}

// SetOffer sets the value of this property. Calling IsOffer afterwards returns
// true.
func (this *AttachmentPropertyIterator) SetOffer(v vocab.OfferInterface) {
	this.clear()
	this.OfferMember = v
}

// SetOrderedCollection sets the value of this property. Calling
// IsOrderedCollection afterwards returns true.
func (this *AttachmentPropertyIterator) SetOrderedCollection(v vocab.OrderedCollectionInterface) {
	this.clear()
	this.OrderedCollectionMember = v
}

// SetOrderedCollectionPage sets the value of this property. Calling
// IsOrderedCollectionPage afterwards returns true.
func (this *AttachmentPropertyIterator) SetOrderedCollectionPage(v vocab.OrderedCollectionPageInterface) {
	this.clear()
	this.OrderedCollectionPageMember = v
}

// SetOrganization sets the value of this property. Calling IsOrganization
// afterwards returns true.
func (this *AttachmentPropertyIterator) SetOrganization(v vocab.OrganizationInterface) {
	this.clear()
	this.OrganizationMember = v
}

// SetPage sets the value of this property. Calling IsPage afterwards returns true.
func (this *AttachmentPropertyIterator) SetPage(v vocab.PageInterface) {
	this.clear()
	this.PageMember = v
}

// SetPerson sets the value of this property. Calling IsPerson afterwards returns
// true.
func (this *AttachmentPropertyIterator) SetPerson(v vocab.PersonInterface) {
	this.clear()
	this.PersonMember = v
}

// SetPlace sets the value of this property. Calling IsPlace afterwards returns
// true.
func (this *AttachmentPropertyIterator) SetPlace(v vocab.PlaceInterface) {
	this.clear()
	this.PlaceMember = v
}

// SetProfile sets the value of this property. Calling IsProfile afterwards
// returns true.
func (this *AttachmentPropertyIterator) SetProfile(v vocab.ProfileInterface) {
	this.clear()
	this.ProfileMember = v
}

// SetQuestion sets the value of this property. Calling IsQuestion afterwards
// returns true.
func (this *AttachmentPropertyIterator) SetQuestion(v vocab.QuestionInterface) {
	this.clear()
	this.QuestionMember = v
}

// SetRead sets the value of this property. Calling IsRead afterwards returns true.
func (this *AttachmentPropertyIterator) SetRead(v vocab.ReadInterface) {
	this.clear()
	this.ReadMember = v
}

// SetReject sets the value of this property. Calling IsReject afterwards returns
// true.
func (this *AttachmentPropertyIterator) SetReject(v vocab.RejectInterface) {
	this.clear()
	this.RejectMember = v
}

// SetRelationship sets the value of this property. Calling IsRelationship
// afterwards returns true.
func (this *AttachmentPropertyIterator) SetRelationship(v vocab.RelationshipInterface) {
	this.clear()
	this.RelationshipMember = v
}

// SetRemove sets the value of this property. Calling IsRemove afterwards returns
// true.
func (this *AttachmentPropertyIterator) SetRemove(v vocab.RemoveInterface) {
	this.clear()
	this.RemoveMember = v
}

// SetService sets the value of this property. Calling IsService afterwards
// returns true.
func (this *AttachmentPropertyIterator) SetService(v vocab.ServiceInterface) {
	this.clear()
	this.ServiceMember = v
}

// SetTentativeAccept sets the value of this property. Calling IsTentativeAccept
// afterwards returns true.
func (this *AttachmentPropertyIterator) SetTentativeAccept(v vocab.TentativeAcceptInterface) {
	this.clear()
	this.TentativeAcceptMember = v
}

// SetTentativeReject sets the value of this property. Calling IsTentativeReject
// afterwards returns true.
func (this *AttachmentPropertyIterator) SetTentativeReject(v vocab.TentativeRejectInterface) {
	this.clear()
	this.TentativeRejectMember = v
}

// SetTombstone sets the value of this property. Calling IsTombstone afterwards
// returns true.
func (this *AttachmentPropertyIterator) SetTombstone(v vocab.TombstoneInterface) {
	this.clear()
	this.TombstoneMember = v
}

// SetTravel sets the value of this property. Calling IsTravel afterwards returns
// true.
func (this *AttachmentPropertyIterator) SetTravel(v vocab.TravelInterface) {
	this.clear()
	this.TravelMember = v
}

// SetUndo sets the value of this property. Calling IsUndo afterwards returns true.
func (this *AttachmentPropertyIterator) SetUndo(v vocab.UndoInterface) {
	this.clear()
	this.UndoMember = v
}

// SetUpdate sets the value of this property. Calling IsUpdate afterwards returns
// true.
func (this *AttachmentPropertyIterator) SetUpdate(v vocab.UpdateInterface) {
	this.clear()
	this.UpdateMember = v
}

// SetVideo sets the value of this property. Calling IsVideo afterwards returns
// true.
func (this *AttachmentPropertyIterator) SetVideo(v vocab.VideoInterface) {
	this.clear()
	this.VideoMember = v
}

// SetView sets the value of this property. Calling IsView afterwards returns true.
func (this *AttachmentPropertyIterator) SetView(v vocab.ViewInterface) {
	this.clear()
	this.ViewMember = v
}

// clear ensures no value of this property is set. Calling HasAny or any of the
// 'Is' methods afterwards will return false.
func (this *AttachmentPropertyIterator) clear() {
	this.ObjectMember = nil
	this.LinkMember = nil
	this.AcceptMember = nil
	this.ActivityMember = nil
	this.AddMember = nil
	this.AnnounceMember = nil
	this.ApplicationMember = nil
	this.ArriveMember = nil
	this.ArticleMember = nil
	this.AudioMember = nil
	this.BlockMember = nil
	this.CollectionMember = nil
	this.CollectionPageMember = nil
	this.CreateMember = nil
	this.DeleteMember = nil
	this.DislikeMember = nil
	this.DocumentMember = nil
	this.EventMember = nil
	this.FlagMember = nil
	this.FollowMember = nil
	this.GroupMember = nil
	this.IgnoreMember = nil
	this.ImageMember = nil
	this.IntransitiveActivityMember = nil
	this.InviteMember = nil
	this.JoinMember = nil
	this.LeaveMember = nil
	this.LikeMember = nil
	this.ListenMember = nil
	this.MentionMember = nil
	this.MoveMember = nil
	this.NoteMember = nil
	this.OfferMember = nil
	this.OrderedCollectionMember = nil
	this.OrderedCollectionPageMember = nil
	this.OrganizationMember = nil
	this.PageMember = nil
	this.PersonMember = nil
	this.PlaceMember = nil
	this.ProfileMember = nil
	this.QuestionMember = nil
	this.ReadMember = nil
	this.RejectMember = nil
	this.RelationshipMember = nil
	this.RemoveMember = nil
	this.ServiceMember = nil
	this.TentativeAcceptMember = nil
	this.TentativeRejectMember = nil
	this.TombstoneMember = nil
	this.TravelMember = nil
	this.UndoMember = nil
	this.UpdateMember = nil
	this.VideoMember = nil
	this.ViewMember = nil
	this.unknown = nil
	this.iri = nil
}

// serialize converts this into an interface representation suitable for
// marshalling into a text or binary format. Applications should not need this
// function as most typical use cases serialize types instead of individual
// properties. It is exposed for alternatives to go-fed implementations to use.
func (this AttachmentPropertyIterator) serialize() (interface{}, error) {
	if this.IsObject() {
		return this.GetObject().Serialize()
	} else if this.IsLink() {
		return this.GetLink().Serialize()
	} else if this.IsAccept() {
		return this.GetAccept().Serialize()
	} else if this.IsActivity() {
		return this.GetActivity().Serialize()
	} else if this.IsAdd() {
		return this.GetAdd().Serialize()
	} else if this.IsAnnounce() {
		return this.GetAnnounce().Serialize()
	} else if this.IsApplication() {
		return this.GetApplication().Serialize()
	} else if this.IsArrive() {
		return this.GetArrive().Serialize()
	} else if this.IsArticle() {
		return this.GetArticle().Serialize()
	} else if this.IsAudio() {
		return this.GetAudio().Serialize()
	} else if this.IsBlock() {
		return this.GetBlock().Serialize()
	} else if this.IsCollection() {
		return this.GetCollection().Serialize()
	} else if this.IsCollectionPage() {
		return this.GetCollectionPage().Serialize()
	} else if this.IsCreate() {
		return this.GetCreate().Serialize()
	} else if this.IsDelete() {
		return this.GetDelete().Serialize()
	} else if this.IsDislike() {
		return this.GetDislike().Serialize()
	} else if this.IsDocument() {
		return this.GetDocument().Serialize()
	} else if this.IsEvent() {
		return this.GetEvent().Serialize()
	} else if this.IsFlag() {
		return this.GetFlag().Serialize()
	} else if this.IsFollow() {
		return this.GetFollow().Serialize()
	} else if this.IsGroup() {
		return this.GetGroup().Serialize()
	} else if this.IsIgnore() {
		return this.GetIgnore().Serialize()
	} else if this.IsImage() {
		return this.GetImage().Serialize()
	} else if this.IsIntransitiveActivity() {
		return this.GetIntransitiveActivity().Serialize()
	} else if this.IsInvite() {
		return this.GetInvite().Serialize()
	} else if this.IsJoin() {
		return this.GetJoin().Serialize()
	} else if this.IsLeave() {
		return this.GetLeave().Serialize()
	} else if this.IsLike() {
		return this.GetLike().Serialize()
	} else if this.IsListen() {
		return this.GetListen().Serialize()
	} else if this.IsMention() {
		return this.GetMention().Serialize()
	} else if this.IsMove() {
		return this.GetMove().Serialize()
	} else if this.IsNote() {
		return this.GetNote().Serialize()
	} else if this.IsOffer() {
		return this.GetOffer().Serialize()
	} else if this.IsOrderedCollection() {
		return this.GetOrderedCollection().Serialize()
	} else if this.IsOrderedCollectionPage() {
		return this.GetOrderedCollectionPage().Serialize()
	} else if this.IsOrganization() {
		return this.GetOrganization().Serialize()
	} else if this.IsPage() {
		return this.GetPage().Serialize()
	} else if this.IsPerson() {
		return this.GetPerson().Serialize()
	} else if this.IsPlace() {
		return this.GetPlace().Serialize()
	} else if this.IsProfile() {
		return this.GetProfile().Serialize()
	} else if this.IsQuestion() {
		return this.GetQuestion().Serialize()
	} else if this.IsRead() {
		return this.GetRead().Serialize()
	} else if this.IsReject() {
		return this.GetReject().Serialize()
	} else if this.IsRelationship() {
		return this.GetRelationship().Serialize()
	} else if this.IsRemove() {
		return this.GetRemove().Serialize()
	} else if this.IsService() {
		return this.GetService().Serialize()
	} else if this.IsTentativeAccept() {
		return this.GetTentativeAccept().Serialize()
	} else if this.IsTentativeReject() {
		return this.GetTentativeReject().Serialize()
	} else if this.IsTombstone() {
		return this.GetTombstone().Serialize()
	} else if this.IsTravel() {
		return this.GetTravel().Serialize()
	} else if this.IsUndo() {
		return this.GetUndo().Serialize()
	} else if this.IsUpdate() {
		return this.GetUpdate().Serialize()
	} else if this.IsVideo() {
		return this.GetVideo().Serialize()
	} else if this.IsView() {
		return this.GetView().Serialize()
	} else if this.IsIRI() {
		return this.iri.String(), nil
	}
	return string(this.unknown), nil
}

// AttachmentProperty is the non-functional property "attachment". It is permitted
// to have one or more values, and of different value types.
type AttachmentProperty struct {
	properties []*AttachmentPropertyIterator
	alias      string
}

// DeserializeAttachmentProperty creates a "attachment" property from an interface
// representation that has been unmarshalled from a text or binary format.
func DeserializeAttachmentProperty(m map[string]interface{}, aliasMap map[string]string) (vocab.AttachmentPropertyInterface, error) {
	alias := ""
	if a, ok := aliasMap["https://www.w3.org/TR/activitystreams-vocabulary"]; ok {
		alias = a
	}
	propName := "attachment"
	if len(alias) > 0 {
		propName = fmt.Sprintf("%s:%s", alias, "attachment")
	}
	if i, ok := m[propName]; ok {
		this := &AttachmentProperty{
			alias:      alias,
			properties: []*AttachmentPropertyIterator{},
		}
		if list, ok := i.([]interface{}); ok {
			for _, iterator := range list {
				if p, err := deserializeAttachmentPropertyIterator(iterator, aliasMap); err != nil {
					return this, err
				} else if p != nil {
					this.properties = append(this.properties, p)
				}
			}
		} else {
			if p, err := deserializeAttachmentPropertyIterator(i, aliasMap); err != nil {
				return this, err
			} else if p != nil {
				this.properties = append(this.properties, p)
			}
		}
		// Set up the properties for iteration.
		for idx, ele := range this.properties {
			ele.parent = this
			ele.myIdx = idx
		}
		return this, nil
	}
	return nil, nil
}

// NewAttachmentProperty creates a new attachment property.
func NewAttachmentProperty() *AttachmentProperty {
	return &AttachmentProperty{alias: ""}
}

// AppendAccept appends a Accept value to the back of a list of the property
// "attachment". Invalidates iterators that are traversing using Prev.
func (this *AttachmentProperty) AppendAccept(v vocab.AcceptInterface) {
	this.properties = append(this.properties, &AttachmentPropertyIterator{
		AcceptMember: v,
		alias:        this.alias,
		myIdx:        this.Len(),
		parent:       this,
	})
}

// AppendActivity appends a Activity value to the back of a list of the property
// "attachment". Invalidates iterators that are traversing using Prev.
func (this *AttachmentProperty) AppendActivity(v vocab.ActivityInterface) {
	this.properties = append(this.properties, &AttachmentPropertyIterator{
		ActivityMember: v,
		alias:          this.alias,
		myIdx:          this.Len(),
		parent:         this,
	})
}

// AppendAdd appends a Add value to the back of a list of the property
// "attachment". Invalidates iterators that are traversing using Prev.
func (this *AttachmentProperty) AppendAdd(v vocab.AddInterface) {
	this.properties = append(this.properties, &AttachmentPropertyIterator{
		AddMember: v,
		alias:     this.alias,
		myIdx:     this.Len(),
		parent:    this,
	})
}

// AppendAnnounce appends a Announce value to the back of a list of the property
// "attachment". Invalidates iterators that are traversing using Prev.
func (this *AttachmentProperty) AppendAnnounce(v vocab.AnnounceInterface) {
	this.properties = append(this.properties, &AttachmentPropertyIterator{
		AnnounceMember: v,
		alias:          this.alias,
		myIdx:          this.Len(),
		parent:         this,
	})
}

// AppendApplication appends a Application value to the back of a list of the
// property "attachment". Invalidates iterators that are traversing using Prev.
func (this *AttachmentProperty) AppendApplication(v vocab.ApplicationInterface) {
	this.properties = append(this.properties, &AttachmentPropertyIterator{
		ApplicationMember: v,
		alias:             this.alias,
		myIdx:             this.Len(),
		parent:            this,
	})
}

// AppendArrive appends a Arrive value to the back of a list of the property
// "attachment". Invalidates iterators that are traversing using Prev.
func (this *AttachmentProperty) AppendArrive(v vocab.ArriveInterface) {
	this.properties = append(this.properties, &AttachmentPropertyIterator{
		ArriveMember: v,
		alias:        this.alias,
		myIdx:        this.Len(),
		parent:       this,
	})
}

// AppendArticle appends a Article value to the back of a list of the property
// "attachment". Invalidates iterators that are traversing using Prev.
func (this *AttachmentProperty) AppendArticle(v vocab.ArticleInterface) {
	this.properties = append(this.properties, &AttachmentPropertyIterator{
		ArticleMember: v,
		alias:         this.alias,
		myIdx:         this.Len(),
		parent:        this,
	})
}

// AppendAudio appends a Audio value to the back of a list of the property
// "attachment". Invalidates iterators that are traversing using Prev.
func (this *AttachmentProperty) AppendAudio(v vocab.AudioInterface) {
	this.properties = append(this.properties, &AttachmentPropertyIterator{
		AudioMember: v,
		alias:       this.alias,
		myIdx:       this.Len(),
		parent:      this,
	})
}

// AppendBlock appends a Block value to the back of a list of the property
// "attachment". Invalidates iterators that are traversing using Prev.
func (this *AttachmentProperty) AppendBlock(v vocab.BlockInterface) {
	this.properties = append(this.properties, &AttachmentPropertyIterator{
		BlockMember: v,
		alias:       this.alias,
		myIdx:       this.Len(),
		parent:      this,
	})
}

// AppendCollection appends a Collection value to the back of a list of the
// property "attachment". Invalidates iterators that are traversing using Prev.
func (this *AttachmentProperty) AppendCollection(v vocab.CollectionInterface) {
	this.properties = append(this.properties, &AttachmentPropertyIterator{
		CollectionMember: v,
		alias:            this.alias,
		myIdx:            this.Len(),
		parent:           this,
	})
}

// AppendCollectionPage appends a CollectionPage value to the back of a list of
// the property "attachment". Invalidates iterators that are traversing using
// Prev.
func (this *AttachmentProperty) AppendCollectionPage(v vocab.CollectionPageInterface) {
	this.properties = append(this.properties, &AttachmentPropertyIterator{
		CollectionPageMember: v,
		alias:                this.alias,
		myIdx:                this.Len(),
		parent:               this,
	})
}

// AppendCreate appends a Create value to the back of a list of the property
// "attachment". Invalidates iterators that are traversing using Prev.
func (this *AttachmentProperty) AppendCreate(v vocab.CreateInterface) {
	this.properties = append(this.properties, &AttachmentPropertyIterator{
		CreateMember: v,
		alias:        this.alias,
		myIdx:        this.Len(),
		parent:       this,
	})
}

// AppendDelete appends a Delete value to the back of a list of the property
// "attachment". Invalidates iterators that are traversing using Prev.
func (this *AttachmentProperty) AppendDelete(v vocab.DeleteInterface) {
	this.properties = append(this.properties, &AttachmentPropertyIterator{
		DeleteMember: v,
		alias:        this.alias,
		myIdx:        this.Len(),
		parent:       this,
	})
}

// AppendDislike appends a Dislike value to the back of a list of the property
// "attachment". Invalidates iterators that are traversing using Prev.
func (this *AttachmentProperty) AppendDislike(v vocab.DislikeInterface) {
	this.properties = append(this.properties, &AttachmentPropertyIterator{
		DislikeMember: v,
		alias:         this.alias,
		myIdx:         this.Len(),
		parent:        this,
	})
}

// AppendDocument appends a Document value to the back of a list of the property
// "attachment". Invalidates iterators that are traversing using Prev.
func (this *AttachmentProperty) AppendDocument(v vocab.DocumentInterface) {
	this.properties = append(this.properties, &AttachmentPropertyIterator{
		DocumentMember: v,
		alias:          this.alias,
		myIdx:          this.Len(),
		parent:         this,
	})
}

// AppendEvent appends a Event value to the back of a list of the property
// "attachment". Invalidates iterators that are traversing using Prev.
func (this *AttachmentProperty) AppendEvent(v vocab.EventInterface) {
	this.properties = append(this.properties, &AttachmentPropertyIterator{
		EventMember: v,
		alias:       this.alias,
		myIdx:       this.Len(),
		parent:      this,
	})
}

// AppendFlag appends a Flag value to the back of a list of the property
// "attachment". Invalidates iterators that are traversing using Prev.
func (this *AttachmentProperty) AppendFlag(v vocab.FlagInterface) {
	this.properties = append(this.properties, &AttachmentPropertyIterator{
		FlagMember: v,
		alias:      this.alias,
		myIdx:      this.Len(),
		parent:     this,
	})
}

// AppendFollow appends a Follow value to the back of a list of the property
// "attachment". Invalidates iterators that are traversing using Prev.
func (this *AttachmentProperty) AppendFollow(v vocab.FollowInterface) {
	this.properties = append(this.properties, &AttachmentPropertyIterator{
		FollowMember: v,
		alias:        this.alias,
		myIdx:        this.Len(),
		parent:       this,
	})
}

// AppendGroup appends a Group value to the back of a list of the property
// "attachment". Invalidates iterators that are traversing using Prev.
func (this *AttachmentProperty) AppendGroup(v vocab.GroupInterface) {
	this.properties = append(this.properties, &AttachmentPropertyIterator{
		GroupMember: v,
		alias:       this.alias,
		myIdx:       this.Len(),
		parent:      this,
	})
}

// AppendIRI appends an IRI value to the back of a list of the property
// "attachment"
func (this *AttachmentProperty) AppendIRI(v *url.URL) {
	this.properties = append(this.properties, &AttachmentPropertyIterator{
		alias:  this.alias,
		iri:    v,
		myIdx:  this.Len(),
		parent: this,
	})
}

// AppendIgnore appends a Ignore value to the back of a list of the property
// "attachment". Invalidates iterators that are traversing using Prev.
func (this *AttachmentProperty) AppendIgnore(v vocab.IgnoreInterface) {
	this.properties = append(this.properties, &AttachmentPropertyIterator{
		IgnoreMember: v,
		alias:        this.alias,
		myIdx:        this.Len(),
		parent:       this,
	})
}

// AppendImage appends a Image value to the back of a list of the property
// "attachment". Invalidates iterators that are traversing using Prev.
func (this *AttachmentProperty) AppendImage(v vocab.ImageInterface) {
	this.properties = append(this.properties, &AttachmentPropertyIterator{
		ImageMember: v,
		alias:       this.alias,
		myIdx:       this.Len(),
		parent:      this,
	})
}

// AppendIntransitiveActivity appends a IntransitiveActivity value to the back of
// a list of the property "attachment". Invalidates iterators that are
// traversing using Prev.
func (this *AttachmentProperty) AppendIntransitiveActivity(v vocab.IntransitiveActivityInterface) {
	this.properties = append(this.properties, &AttachmentPropertyIterator{
		IntransitiveActivityMember: v,
		alias:                      this.alias,
		myIdx:                      this.Len(),
		parent:                     this,
	})
}

// AppendInvite appends a Invite value to the back of a list of the property
// "attachment". Invalidates iterators that are traversing using Prev.
func (this *AttachmentProperty) AppendInvite(v vocab.InviteInterface) {
	this.properties = append(this.properties, &AttachmentPropertyIterator{
		InviteMember: v,
		alias:        this.alias,
		myIdx:        this.Len(),
		parent:       this,
	})
}

// AppendJoin appends a Join value to the back of a list of the property
// "attachment". Invalidates iterators that are traversing using Prev.
func (this *AttachmentProperty) AppendJoin(v vocab.JoinInterface) {
	this.properties = append(this.properties, &AttachmentPropertyIterator{
		JoinMember: v,
		alias:      this.alias,
		myIdx:      this.Len(),
		parent:     this,
	})
}

// AppendLeave appends a Leave value to the back of a list of the property
// "attachment". Invalidates iterators that are traversing using Prev.
func (this *AttachmentProperty) AppendLeave(v vocab.LeaveInterface) {
	this.properties = append(this.properties, &AttachmentPropertyIterator{
		LeaveMember: v,
		alias:       this.alias,
		myIdx:       this.Len(),
		parent:      this,
	})
}

// AppendLike appends a Like value to the back of a list of the property
// "attachment". Invalidates iterators that are traversing using Prev.
func (this *AttachmentProperty) AppendLike(v vocab.LikeInterface) {
	this.properties = append(this.properties, &AttachmentPropertyIterator{
		LikeMember: v,
		alias:      this.alias,
		myIdx:      this.Len(),
		parent:     this,
	})
}

// AppendLink appends a Link value to the back of a list of the property
// "attachment". Invalidates iterators that are traversing using Prev.
func (this *AttachmentProperty) AppendLink(v vocab.LinkInterface) {
	this.properties = append(this.properties, &AttachmentPropertyIterator{
		LinkMember: v,
		alias:      this.alias,
		myIdx:      this.Len(),
		parent:     this,
	})
}

// AppendListen appends a Listen value to the back of a list of the property
// "attachment". Invalidates iterators that are traversing using Prev.
func (this *AttachmentProperty) AppendListen(v vocab.ListenInterface) {
	this.properties = append(this.properties, &AttachmentPropertyIterator{
		ListenMember: v,
		alias:        this.alias,
		myIdx:        this.Len(),
		parent:       this,
	})
}

// AppendMention appends a Mention value to the back of a list of the property
// "attachment". Invalidates iterators that are traversing using Prev.
func (this *AttachmentProperty) AppendMention(v vocab.MentionInterface) {
	this.properties = append(this.properties, &AttachmentPropertyIterator{
		MentionMember: v,
		alias:         this.alias,
		myIdx:         this.Len(),
		parent:        this,
	})
}

// AppendMove appends a Move value to the back of a list of the property
// "attachment". Invalidates iterators that are traversing using Prev.
func (this *AttachmentProperty) AppendMove(v vocab.MoveInterface) {
	this.properties = append(this.properties, &AttachmentPropertyIterator{
		MoveMember: v,
		alias:      this.alias,
		myIdx:      this.Len(),
		parent:     this,
	})
}

// AppendNote appends a Note value to the back of a list of the property
// "attachment". Invalidates iterators that are traversing using Prev.
func (this *AttachmentProperty) AppendNote(v vocab.NoteInterface) {
	this.properties = append(this.properties, &AttachmentPropertyIterator{
		NoteMember: v,
		alias:      this.alias,
		myIdx:      this.Len(),
		parent:     this,
	})
}

// AppendObject appends a Object value to the back of a list of the property
// "attachment". Invalidates iterators that are traversing using Prev.
func (this *AttachmentProperty) AppendObject(v vocab.ObjectInterface) {
	this.properties = append(this.properties, &AttachmentPropertyIterator{
		ObjectMember: v,
		alias:        this.alias,
		myIdx:        this.Len(),
		parent:       this,
	})
}

// AppendOffer appends a Offer value to the back of a list of the property
// "attachment". Invalidates iterators that are traversing using Prev.
func (this *AttachmentProperty) AppendOffer(v vocab.OfferInterface) {
	this.properties = append(this.properties, &AttachmentPropertyIterator{
		OfferMember: v,
		alias:       this.alias,
		myIdx:       this.Len(),
		parent:      this,
	})
}

// AppendOrderedCollection appends a OrderedCollection value to the back of a list
// of the property "attachment". Invalidates iterators that are traversing
// using Prev.
func (this *AttachmentProperty) AppendOrderedCollection(v vocab.OrderedCollectionInterface) {
	this.properties = append(this.properties, &AttachmentPropertyIterator{
		OrderedCollectionMember: v,
		alias:                   this.alias,
		myIdx:                   this.Len(),
		parent:                  this,
	})
}

// AppendOrderedCollectionPage appends a OrderedCollectionPage value to the back
// of a list of the property "attachment". Invalidates iterators that are
// traversing using Prev.
func (this *AttachmentProperty) AppendOrderedCollectionPage(v vocab.OrderedCollectionPageInterface) {
	this.properties = append(this.properties, &AttachmentPropertyIterator{
		OrderedCollectionPageMember: v,
		alias:                       this.alias,
		myIdx:                       this.Len(),
		parent:                      this,
	})
}

// AppendOrganization appends a Organization value to the back of a list of the
// property "attachment". Invalidates iterators that are traversing using Prev.
func (this *AttachmentProperty) AppendOrganization(v vocab.OrganizationInterface) {
	this.properties = append(this.properties, &AttachmentPropertyIterator{
		OrganizationMember: v,
		alias:              this.alias,
		myIdx:              this.Len(),
		parent:             this,
	})
}

// AppendPage appends a Page value to the back of a list of the property
// "attachment". Invalidates iterators that are traversing using Prev.
func (this *AttachmentProperty) AppendPage(v vocab.PageInterface) {
	this.properties = append(this.properties, &AttachmentPropertyIterator{
		PageMember: v,
		alias:      this.alias,
		myIdx:      this.Len(),
		parent:     this,
	})
}

// AppendPerson appends a Person value to the back of a list of the property
// "attachment". Invalidates iterators that are traversing using Prev.
func (this *AttachmentProperty) AppendPerson(v vocab.PersonInterface) {
	this.properties = append(this.properties, &AttachmentPropertyIterator{
		PersonMember: v,
		alias:        this.alias,
		myIdx:        this.Len(),
		parent:       this,
	})
}

// AppendPlace appends a Place value to the back of a list of the property
// "attachment". Invalidates iterators that are traversing using Prev.
func (this *AttachmentProperty) AppendPlace(v vocab.PlaceInterface) {
	this.properties = append(this.properties, &AttachmentPropertyIterator{
		PlaceMember: v,
		alias:       this.alias,
		myIdx:       this.Len(),
		parent:      this,
	})
}

// AppendProfile appends a Profile value to the back of a list of the property
// "attachment". Invalidates iterators that are traversing using Prev.
func (this *AttachmentProperty) AppendProfile(v vocab.ProfileInterface) {
	this.properties = append(this.properties, &AttachmentPropertyIterator{
		ProfileMember: v,
		alias:         this.alias,
		myIdx:         this.Len(),
		parent:        this,
	})
}

// AppendQuestion appends a Question value to the back of a list of the property
// "attachment". Invalidates iterators that are traversing using Prev.
func (this *AttachmentProperty) AppendQuestion(v vocab.QuestionInterface) {
	this.properties = append(this.properties, &AttachmentPropertyIterator{
		QuestionMember: v,
		alias:          this.alias,
		myIdx:          this.Len(),
		parent:         this,
	})
}

// AppendRead appends a Read value to the back of a list of the property
// "attachment". Invalidates iterators that are traversing using Prev.
func (this *AttachmentProperty) AppendRead(v vocab.ReadInterface) {
	this.properties = append(this.properties, &AttachmentPropertyIterator{
		ReadMember: v,
		alias:      this.alias,
		myIdx:      this.Len(),
		parent:     this,
	})
}

// AppendReject appends a Reject value to the back of a list of the property
// "attachment". Invalidates iterators that are traversing using Prev.
func (this *AttachmentProperty) AppendReject(v vocab.RejectInterface) {
	this.properties = append(this.properties, &AttachmentPropertyIterator{
		RejectMember: v,
		alias:        this.alias,
		myIdx:        this.Len(),
		parent:       this,
	})
}

// AppendRelationship appends a Relationship value to the back of a list of the
// property "attachment". Invalidates iterators that are traversing using Prev.
func (this *AttachmentProperty) AppendRelationship(v vocab.RelationshipInterface) {
	this.properties = append(this.properties, &AttachmentPropertyIterator{
		RelationshipMember: v,
		alias:              this.alias,
		myIdx:              this.Len(),
		parent:             this,
	})
}

// AppendRemove appends a Remove value to the back of a list of the property
// "attachment". Invalidates iterators that are traversing using Prev.
func (this *AttachmentProperty) AppendRemove(v vocab.RemoveInterface) {
	this.properties = append(this.properties, &AttachmentPropertyIterator{
		RemoveMember: v,
		alias:        this.alias,
		myIdx:        this.Len(),
		parent:       this,
	})
}

// AppendService appends a Service value to the back of a list of the property
// "attachment". Invalidates iterators that are traversing using Prev.
func (this *AttachmentProperty) AppendService(v vocab.ServiceInterface) {
	this.properties = append(this.properties, &AttachmentPropertyIterator{
		ServiceMember: v,
		alias:         this.alias,
		myIdx:         this.Len(),
		parent:        this,
	})
}

// AppendTentativeAccept appends a TentativeAccept value to the back of a list of
// the property "attachment". Invalidates iterators that are traversing using
// Prev.
func (this *AttachmentProperty) AppendTentativeAccept(v vocab.TentativeAcceptInterface) {
	this.properties = append(this.properties, &AttachmentPropertyIterator{
		TentativeAcceptMember: v,
		alias:                 this.alias,
		myIdx:                 this.Len(),
		parent:                this,
	})
}

// AppendTentativeReject appends a TentativeReject value to the back of a list of
// the property "attachment". Invalidates iterators that are traversing using
// Prev.
func (this *AttachmentProperty) AppendTentativeReject(v vocab.TentativeRejectInterface) {
	this.properties = append(this.properties, &AttachmentPropertyIterator{
		TentativeRejectMember: v,
		alias:                 this.alias,
		myIdx:                 this.Len(),
		parent:                this,
	})
}

// AppendTombstone appends a Tombstone value to the back of a list of the property
// "attachment". Invalidates iterators that are traversing using Prev.
func (this *AttachmentProperty) AppendTombstone(v vocab.TombstoneInterface) {
	this.properties = append(this.properties, &AttachmentPropertyIterator{
		TombstoneMember: v,
		alias:           this.alias,
		myIdx:           this.Len(),
		parent:          this,
	})
}

// AppendTravel appends a Travel value to the back of a list of the property
// "attachment". Invalidates iterators that are traversing using Prev.
func (this *AttachmentProperty) AppendTravel(v vocab.TravelInterface) {
	this.properties = append(this.properties, &AttachmentPropertyIterator{
		TravelMember: v,
		alias:        this.alias,
		myIdx:        this.Len(),
		parent:       this,
	})
}

// AppendUndo appends a Undo value to the back of a list of the property
// "attachment". Invalidates iterators that are traversing using Prev.
func (this *AttachmentProperty) AppendUndo(v vocab.UndoInterface) {
	this.properties = append(this.properties, &AttachmentPropertyIterator{
		UndoMember: v,
		alias:      this.alias,
		myIdx:      this.Len(),
		parent:     this,
	})
}

// AppendUpdate appends a Update value to the back of a list of the property
// "attachment". Invalidates iterators that are traversing using Prev.
func (this *AttachmentProperty) AppendUpdate(v vocab.UpdateInterface) {
	this.properties = append(this.properties, &AttachmentPropertyIterator{
		UpdateMember: v,
		alias:        this.alias,
		myIdx:        this.Len(),
		parent:       this,
	})
}

// AppendVideo appends a Video value to the back of a list of the property
// "attachment". Invalidates iterators that are traversing using Prev.
func (this *AttachmentProperty) AppendVideo(v vocab.VideoInterface) {
	this.properties = append(this.properties, &AttachmentPropertyIterator{
		VideoMember: v,
		alias:       this.alias,
		myIdx:       this.Len(),
		parent:      this,
	})
}

// AppendView appends a View value to the back of a list of the property
// "attachment". Invalidates iterators that are traversing using Prev.
func (this *AttachmentProperty) AppendView(v vocab.ViewInterface) {
	this.properties = append(this.properties, &AttachmentPropertyIterator{
		ViewMember: v,
		alias:      this.alias,
		myIdx:      this.Len(),
		parent:     this,
	})
}

// At returns the property value for the specified index. Panics if the index is
// out of bounds.
func (this AttachmentProperty) At(index int) vocab.AttachmentPropertyIteratorInterface {
	return this.properties[index]
}

// Begin returns the first iterator, or nil if empty. Can be used with the
// iterator's Next method and this property's End method to iterate from front
// to back through all values.
func (this AttachmentProperty) Begin() vocab.AttachmentPropertyIteratorInterface {
	if this.Empty() {
		return nil
	} else {
		return this.properties[0]
	}
}

// Empty returns returns true if there are no elements.
func (this AttachmentProperty) Empty() bool {
	return this.Len() == 0
}

// End returns beyond-the-last iterator, which is nil. Can be used with the
// iterator's Next method and this property's Begin method to iterate from
// front to back through all values.
func (this AttachmentProperty) End() vocab.AttachmentPropertyIteratorInterface {
	return nil
}

// JSONLDContext returns the JSONLD URIs required in the context string for this
// property and the specific values that are set. The value in the map is the
// alias used to import the property's value or values.
func (this AttachmentProperty) JSONLDContext() map[string]string {
	m := map[string]string{"https://www.w3.org/TR/activitystreams-vocabulary": this.alias}
	for _, elem := range this.properties {
		child := elem.JSONLDContext()
		/*
		   Since the literal maps in this function are determined at
		   code-generation time, this loop should not overwrite an existing key with a
		   new value.
		*/
		for k, v := range child {
			m[k] = v
		}
	}
	return m
}

// KindIndex computes an arbitrary value for indexing this kind of value. This is
// a leaky API method specifically needed only for alternate implementations
// for go-fed. Applications should not use this method. Panics if the index is
// out of bounds.
func (this AttachmentProperty) KindIndex(idx int) int {
	return this.properties[idx].KindIndex()
}

// Len returns the number of values that exist for the "attachment" property.
func (this AttachmentProperty) Len() (length int) {
	return len(this.properties)
}

// Less computes whether another property is less than this one. Mixing types
// results in a consistent but arbitrary ordering
func (this AttachmentProperty) Less(i, j int) bool {
	idx1 := this.KindIndex(i)
	idx2 := this.KindIndex(j)
	if idx1 < idx2 {
		return true
	} else if idx1 == idx2 {
		if idx1 == 0 {
			lhs := this.properties[i].GetObject()
			rhs := this.properties[j].GetObject()
			return lhs.LessThan(rhs)
		} else if idx1 == 1 {
			lhs := this.properties[i].GetLink()
			rhs := this.properties[j].GetLink()
			return lhs.LessThan(rhs)
		} else if idx1 == 2 {
			lhs := this.properties[i].GetAccept()
			rhs := this.properties[j].GetAccept()
			return lhs.LessThan(rhs)
		} else if idx1 == 3 {
			lhs := this.properties[i].GetActivity()
			rhs := this.properties[j].GetActivity()
			return lhs.LessThan(rhs)
		} else if idx1 == 4 {
			lhs := this.properties[i].GetAdd()
			rhs := this.properties[j].GetAdd()
			return lhs.LessThan(rhs)
		} else if idx1 == 5 {
			lhs := this.properties[i].GetAnnounce()
			rhs := this.properties[j].GetAnnounce()
			return lhs.LessThan(rhs)
		} else if idx1 == 6 {
			lhs := this.properties[i].GetApplication()
			rhs := this.properties[j].GetApplication()
			return lhs.LessThan(rhs)
		} else if idx1 == 7 {
			lhs := this.properties[i].GetArrive()
			rhs := this.properties[j].GetArrive()
			return lhs.LessThan(rhs)
		} else if idx1 == 8 {
			lhs := this.properties[i].GetArticle()
			rhs := this.properties[j].GetArticle()
			return lhs.LessThan(rhs)
		} else if idx1 == 9 {
			lhs := this.properties[i].GetAudio()
			rhs := this.properties[j].GetAudio()
			return lhs.LessThan(rhs)
		} else if idx1 == 10 {
			lhs := this.properties[i].GetBlock()
			rhs := this.properties[j].GetBlock()
			return lhs.LessThan(rhs)
		} else if idx1 == 11 {
			lhs := this.properties[i].GetCollection()
			rhs := this.properties[j].GetCollection()
			return lhs.LessThan(rhs)
		} else if idx1 == 12 {
			lhs := this.properties[i].GetCollectionPage()
			rhs := this.properties[j].GetCollectionPage()
			return lhs.LessThan(rhs)
		} else if idx1 == 13 {
			lhs := this.properties[i].GetCreate()
			rhs := this.properties[j].GetCreate()
			return lhs.LessThan(rhs)
		} else if idx1 == 14 {
			lhs := this.properties[i].GetDelete()
			rhs := this.properties[j].GetDelete()
			return lhs.LessThan(rhs)
		} else if idx1 == 15 {
			lhs := this.properties[i].GetDislike()
			rhs := this.properties[j].GetDislike()
			return lhs.LessThan(rhs)
		} else if idx1 == 16 {
			lhs := this.properties[i].GetDocument()
			rhs := this.properties[j].GetDocument()
			return lhs.LessThan(rhs)
		} else if idx1 == 17 {
			lhs := this.properties[i].GetEvent()
			rhs := this.properties[j].GetEvent()
			return lhs.LessThan(rhs)
		} else if idx1 == 18 {
			lhs := this.properties[i].GetFlag()
			rhs := this.properties[j].GetFlag()
			return lhs.LessThan(rhs)
		} else if idx1 == 19 {
			lhs := this.properties[i].GetFollow()
			rhs := this.properties[j].GetFollow()
			return lhs.LessThan(rhs)
		} else if idx1 == 20 {
			lhs := this.properties[i].GetGroup()
			rhs := this.properties[j].GetGroup()
			return lhs.LessThan(rhs)
		} else if idx1 == 21 {
			lhs := this.properties[i].GetIgnore()
			rhs := this.properties[j].GetIgnore()
			return lhs.LessThan(rhs)
		} else if idx1 == 22 {
			lhs := this.properties[i].GetImage()
			rhs := this.properties[j].GetImage()
			return lhs.LessThan(rhs)
		} else if idx1 == 23 {
			lhs := this.properties[i].GetIntransitiveActivity()
			rhs := this.properties[j].GetIntransitiveActivity()
			return lhs.LessThan(rhs)
		} else if idx1 == 24 {
			lhs := this.properties[i].GetInvite()
			rhs := this.properties[j].GetInvite()
			return lhs.LessThan(rhs)
		} else if idx1 == 25 {
			lhs := this.properties[i].GetJoin()
			rhs := this.properties[j].GetJoin()
			return lhs.LessThan(rhs)
		} else if idx1 == 26 {
			lhs := this.properties[i].GetLeave()
			rhs := this.properties[j].GetLeave()
			return lhs.LessThan(rhs)
		} else if idx1 == 27 {
			lhs := this.properties[i].GetLike()
			rhs := this.properties[j].GetLike()
			return lhs.LessThan(rhs)
		} else if idx1 == 28 {
			lhs := this.properties[i].GetListen()
			rhs := this.properties[j].GetListen()
			return lhs.LessThan(rhs)
		} else if idx1 == 29 {
			lhs := this.properties[i].GetMention()
			rhs := this.properties[j].GetMention()
			return lhs.LessThan(rhs)
		} else if idx1 == 30 {
			lhs := this.properties[i].GetMove()
			rhs := this.properties[j].GetMove()
			return lhs.LessThan(rhs)
		} else if idx1 == 31 {
			lhs := this.properties[i].GetNote()
			rhs := this.properties[j].GetNote()
			return lhs.LessThan(rhs)
		} else if idx1 == 32 {
			lhs := this.properties[i].GetOffer()
			rhs := this.properties[j].GetOffer()
			return lhs.LessThan(rhs)
		} else if idx1 == 33 {
			lhs := this.properties[i].GetOrderedCollection()
			rhs := this.properties[j].GetOrderedCollection()
			return lhs.LessThan(rhs)
		} else if idx1 == 34 {
			lhs := this.properties[i].GetOrderedCollectionPage()
			rhs := this.properties[j].GetOrderedCollectionPage()
			return lhs.LessThan(rhs)
		} else if idx1 == 35 {
			lhs := this.properties[i].GetOrganization()
			rhs := this.properties[j].GetOrganization()
			return lhs.LessThan(rhs)
		} else if idx1 == 36 {
			lhs := this.properties[i].GetPage()
			rhs := this.properties[j].GetPage()
			return lhs.LessThan(rhs)
		} else if idx1 == 37 {
			lhs := this.properties[i].GetPerson()
			rhs := this.properties[j].GetPerson()
			return lhs.LessThan(rhs)
		} else if idx1 == 38 {
			lhs := this.properties[i].GetPlace()
			rhs := this.properties[j].GetPlace()
			return lhs.LessThan(rhs)
		} else if idx1 == 39 {
			lhs := this.properties[i].GetProfile()
			rhs := this.properties[j].GetProfile()
			return lhs.LessThan(rhs)
		} else if idx1 == 40 {
			lhs := this.properties[i].GetQuestion()
			rhs := this.properties[j].GetQuestion()
			return lhs.LessThan(rhs)
		} else if idx1 == 41 {
			lhs := this.properties[i].GetRead()
			rhs := this.properties[j].GetRead()
			return lhs.LessThan(rhs)
		} else if idx1 == 42 {
			lhs := this.properties[i].GetReject()
			rhs := this.properties[j].GetReject()
			return lhs.LessThan(rhs)
		} else if idx1 == 43 {
			lhs := this.properties[i].GetRelationship()
			rhs := this.properties[j].GetRelationship()
			return lhs.LessThan(rhs)
		} else if idx1 == 44 {
			lhs := this.properties[i].GetRemove()
			rhs := this.properties[j].GetRemove()
			return lhs.LessThan(rhs)
		} else if idx1 == 45 {
			lhs := this.properties[i].GetService()
			rhs := this.properties[j].GetService()
			return lhs.LessThan(rhs)
		} else if idx1 == 46 {
			lhs := this.properties[i].GetTentativeAccept()
			rhs := this.properties[j].GetTentativeAccept()
			return lhs.LessThan(rhs)
		} else if idx1 == 47 {
			lhs := this.properties[i].GetTentativeReject()
			rhs := this.properties[j].GetTentativeReject()
			return lhs.LessThan(rhs)
		} else if idx1 == 48 {
			lhs := this.properties[i].GetTombstone()
			rhs := this.properties[j].GetTombstone()
			return lhs.LessThan(rhs)
		} else if idx1 == 49 {
			lhs := this.properties[i].GetTravel()
			rhs := this.properties[j].GetTravel()
			return lhs.LessThan(rhs)
		} else if idx1 == 50 {
			lhs := this.properties[i].GetUndo()
			rhs := this.properties[j].GetUndo()
			return lhs.LessThan(rhs)
		} else if idx1 == 51 {
			lhs := this.properties[i].GetUpdate()
			rhs := this.properties[j].GetUpdate()
			return lhs.LessThan(rhs)
		} else if idx1 == 52 {
			lhs := this.properties[i].GetVideo()
			rhs := this.properties[j].GetVideo()
			return lhs.LessThan(rhs)
		} else if idx1 == 53 {
			lhs := this.properties[i].GetView()
			rhs := this.properties[j].GetView()
			return lhs.LessThan(rhs)
		} else if idx1 == -2 {
			lhs := this.properties[i].GetIRI()
			rhs := this.properties[j].GetIRI()
			return lhs.String() < rhs.String()
		}
	}
	return false
}

// LessThan compares two instances of this property with an arbitrary but stable
// comparison. Applications should not use this because it is only meant to
// help alternative implementations to go-fed to be able to normalize
// nonfunctional properties.
func (this AttachmentProperty) LessThan(o vocab.AttachmentPropertyInterface) bool {
	l1 := this.Len()
	l2 := o.Len()
	l := l1
	if l2 < l1 {
		l = l2
	}
	for i := 0; i < l; i++ {
		if this.properties[i].LessThan(o.At(i)) {
			return true
		} else if o.At(i).LessThan(this.properties[i]) {
			return false
		}
	}
	return l1 < l2
}

// Name returns the name of this property: "attachment".
func (this AttachmentProperty) Name() string {
	return "attachment"
}

// PrependAccept prepends a Accept value to the front of a list of the property
// "attachment". Invalidates all iterators.
func (this *AttachmentProperty) PrependAccept(v vocab.AcceptInterface) {
	this.properties = append([]*AttachmentPropertyIterator{{
		AcceptMember: v,
		alias:        this.alias,
		myIdx:        0,
		parent:       this,
	}}, this.properties...)
	for i := 1; i < this.Len(); i++ {
		(this.properties)[i].myIdx = i
	}
}

// PrependActivity prepends a Activity value to the front of a list of the
// property "attachment". Invalidates all iterators.
func (this *AttachmentProperty) PrependActivity(v vocab.ActivityInterface) {
	this.properties = append([]*AttachmentPropertyIterator{{
		ActivityMember: v,
		alias:          this.alias,
		myIdx:          0,
		parent:         this,
	}}, this.properties...)
	for i := 1; i < this.Len(); i++ {
		(this.properties)[i].myIdx = i
	}
}

// PrependAdd prepends a Add value to the front of a list of the property
// "attachment". Invalidates all iterators.
func (this *AttachmentProperty) PrependAdd(v vocab.AddInterface) {
	this.properties = append([]*AttachmentPropertyIterator{{
		AddMember: v,
		alias:     this.alias,
		myIdx:     0,
		parent:    this,
	}}, this.properties...)
	for i := 1; i < this.Len(); i++ {
		(this.properties)[i].myIdx = i
	}
}

// PrependAnnounce prepends a Announce value to the front of a list of the
// property "attachment". Invalidates all iterators.
func (this *AttachmentProperty) PrependAnnounce(v vocab.AnnounceInterface) {
	this.properties = append([]*AttachmentPropertyIterator{{
		AnnounceMember: v,
		alias:          this.alias,
		myIdx:          0,
		parent:         this,
	}}, this.properties...)
	for i := 1; i < this.Len(); i++ {
		(this.properties)[i].myIdx = i
	}
}

// PrependApplication prepends a Application value to the front of a list of the
// property "attachment". Invalidates all iterators.
func (this *AttachmentProperty) PrependApplication(v vocab.ApplicationInterface) {
	this.properties = append([]*AttachmentPropertyIterator{{
		ApplicationMember: v,
		alias:             this.alias,
		myIdx:             0,
		parent:            this,
	}}, this.properties...)
	for i := 1; i < this.Len(); i++ {
		(this.properties)[i].myIdx = i
	}
}

// PrependArrive prepends a Arrive value to the front of a list of the property
// "attachment". Invalidates all iterators.
func (this *AttachmentProperty) PrependArrive(v vocab.ArriveInterface) {
	this.properties = append([]*AttachmentPropertyIterator{{
		ArriveMember: v,
		alias:        this.alias,
		myIdx:        0,
		parent:       this,
	}}, this.properties...)
	for i := 1; i < this.Len(); i++ {
		(this.properties)[i].myIdx = i
	}
}

// PrependArticle prepends a Article value to the front of a list of the property
// "attachment". Invalidates all iterators.
func (this *AttachmentProperty) PrependArticle(v vocab.ArticleInterface) {
	this.properties = append([]*AttachmentPropertyIterator{{
		ArticleMember: v,
		alias:         this.alias,
		myIdx:         0,
		parent:        this,
	}}, this.properties...)
	for i := 1; i < this.Len(); i++ {
		(this.properties)[i].myIdx = i
	}
}

// PrependAudio prepends a Audio value to the front of a list of the property
// "attachment". Invalidates all iterators.
func (this *AttachmentProperty) PrependAudio(v vocab.AudioInterface) {
	this.properties = append([]*AttachmentPropertyIterator{{
		AudioMember: v,
		alias:       this.alias,
		myIdx:       0,
		parent:      this,
	}}, this.properties...)
	for i := 1; i < this.Len(); i++ {
		(this.properties)[i].myIdx = i
	}
}

// PrependBlock prepends a Block value to the front of a list of the property
// "attachment". Invalidates all iterators.
func (this *AttachmentProperty) PrependBlock(v vocab.BlockInterface) {
	this.properties = append([]*AttachmentPropertyIterator{{
		BlockMember: v,
		alias:       this.alias,
		myIdx:       0,
		parent:      this,
	}}, this.properties...)
	for i := 1; i < this.Len(); i++ {
		(this.properties)[i].myIdx = i
	}
}

// PrependCollection prepends a Collection value to the front of a list of the
// property "attachment". Invalidates all iterators.
func (this *AttachmentProperty) PrependCollection(v vocab.CollectionInterface) {
	this.properties = append([]*AttachmentPropertyIterator{{
		CollectionMember: v,
		alias:            this.alias,
		myIdx:            0,
		parent:           this,
	}}, this.properties...)
	for i := 1; i < this.Len(); i++ {
		(this.properties)[i].myIdx = i
	}
}

// PrependCollectionPage prepends a CollectionPage value to the front of a list of
// the property "attachment". Invalidates all iterators.
func (this *AttachmentProperty) PrependCollectionPage(v vocab.CollectionPageInterface) {
	this.properties = append([]*AttachmentPropertyIterator{{
		CollectionPageMember: v,
		alias:                this.alias,
		myIdx:                0,
		parent:               this,
	}}, this.properties...)
	for i := 1; i < this.Len(); i++ {
		(this.properties)[i].myIdx = i
	}
}

// PrependCreate prepends a Create value to the front of a list of the property
// "attachment". Invalidates all iterators.
func (this *AttachmentProperty) PrependCreate(v vocab.CreateInterface) {
	this.properties = append([]*AttachmentPropertyIterator{{
		CreateMember: v,
		alias:        this.alias,
		myIdx:        0,
		parent:       this,
	}}, this.properties...)
	for i := 1; i < this.Len(); i++ {
		(this.properties)[i].myIdx = i
	}
}

// PrependDelete prepends a Delete value to the front of a list of the property
// "attachment". Invalidates all iterators.
func (this *AttachmentProperty) PrependDelete(v vocab.DeleteInterface) {
	this.properties = append([]*AttachmentPropertyIterator{{
		DeleteMember: v,
		alias:        this.alias,
		myIdx:        0,
		parent:       this,
	}}, this.properties...)
	for i := 1; i < this.Len(); i++ {
		(this.properties)[i].myIdx = i
	}
}

// PrependDislike prepends a Dislike value to the front of a list of the property
// "attachment". Invalidates all iterators.
func (this *AttachmentProperty) PrependDislike(v vocab.DislikeInterface) {
	this.properties = append([]*AttachmentPropertyIterator{{
		DislikeMember: v,
		alias:         this.alias,
		myIdx:         0,
		parent:        this,
	}}, this.properties...)
	for i := 1; i < this.Len(); i++ {
		(this.properties)[i].myIdx = i
	}
}

// PrependDocument prepends a Document value to the front of a list of the
// property "attachment". Invalidates all iterators.
func (this *AttachmentProperty) PrependDocument(v vocab.DocumentInterface) {
	this.properties = append([]*AttachmentPropertyIterator{{
		DocumentMember: v,
		alias:          this.alias,
		myIdx:          0,
		parent:         this,
	}}, this.properties...)
	for i := 1; i < this.Len(); i++ {
		(this.properties)[i].myIdx = i
	}
}

// PrependEvent prepends a Event value to the front of a list of the property
// "attachment". Invalidates all iterators.
func (this *AttachmentProperty) PrependEvent(v vocab.EventInterface) {
	this.properties = append([]*AttachmentPropertyIterator{{
		EventMember: v,
		alias:       this.alias,
		myIdx:       0,
		parent:      this,
	}}, this.properties...)
	for i := 1; i < this.Len(); i++ {
		(this.properties)[i].myIdx = i
	}
}

// PrependFlag prepends a Flag value to the front of a list of the property
// "attachment". Invalidates all iterators.
func (this *AttachmentProperty) PrependFlag(v vocab.FlagInterface) {
	this.properties = append([]*AttachmentPropertyIterator{{
		FlagMember: v,
		alias:      this.alias,
		myIdx:      0,
		parent:     this,
	}}, this.properties...)
	for i := 1; i < this.Len(); i++ {
		(this.properties)[i].myIdx = i
	}
}

// PrependFollow prepends a Follow value to the front of a list of the property
// "attachment". Invalidates all iterators.
func (this *AttachmentProperty) PrependFollow(v vocab.FollowInterface) {
	this.properties = append([]*AttachmentPropertyIterator{{
		FollowMember: v,
		alias:        this.alias,
		myIdx:        0,
		parent:       this,
	}}, this.properties...)
	for i := 1; i < this.Len(); i++ {
		(this.properties)[i].myIdx = i
	}
}

// PrependGroup prepends a Group value to the front of a list of the property
// "attachment". Invalidates all iterators.
func (this *AttachmentProperty) PrependGroup(v vocab.GroupInterface) {
	this.properties = append([]*AttachmentPropertyIterator{{
		GroupMember: v,
		alias:       this.alias,
		myIdx:       0,
		parent:      this,
	}}, this.properties...)
	for i := 1; i < this.Len(); i++ {
		(this.properties)[i].myIdx = i
	}
}

// PrependIRI prepends an IRI value to the front of a list of the property
// "attachment".
func (this *AttachmentProperty) PrependIRI(v *url.URL) {
	this.properties = append([]*AttachmentPropertyIterator{{
		alias:  this.alias,
		iri:    v,
		myIdx:  0,
		parent: this,
	}}, this.properties...)
	for i := 1; i < this.Len(); i++ {
		(this.properties)[i].myIdx = i
	}
}

// PrependIgnore prepends a Ignore value to the front of a list of the property
// "attachment". Invalidates all iterators.
func (this *AttachmentProperty) PrependIgnore(v vocab.IgnoreInterface) {
	this.properties = append([]*AttachmentPropertyIterator{{
		IgnoreMember: v,
		alias:        this.alias,
		myIdx:        0,
		parent:       this,
	}}, this.properties...)
	for i := 1; i < this.Len(); i++ {
		(this.properties)[i].myIdx = i
	}
}

// PrependImage prepends a Image value to the front of a list of the property
// "attachment". Invalidates all iterators.
func (this *AttachmentProperty) PrependImage(v vocab.ImageInterface) {
	this.properties = append([]*AttachmentPropertyIterator{{
		ImageMember: v,
		alias:       this.alias,
		myIdx:       0,
		parent:      this,
	}}, this.properties...)
	for i := 1; i < this.Len(); i++ {
		(this.properties)[i].myIdx = i
	}
}

// PrependIntransitiveActivity prepends a IntransitiveActivity value to the front
// of a list of the property "attachment". Invalidates all iterators.
func (this *AttachmentProperty) PrependIntransitiveActivity(v vocab.IntransitiveActivityInterface) {
	this.properties = append([]*AttachmentPropertyIterator{{
		IntransitiveActivityMember: v,
		alias:                      this.alias,
		myIdx:                      0,
		parent:                     this,
	}}, this.properties...)
	for i := 1; i < this.Len(); i++ {
		(this.properties)[i].myIdx = i
	}
}

// PrependInvite prepends a Invite value to the front of a list of the property
// "attachment". Invalidates all iterators.
func (this *AttachmentProperty) PrependInvite(v vocab.InviteInterface) {
	this.properties = append([]*AttachmentPropertyIterator{{
		InviteMember: v,
		alias:        this.alias,
		myIdx:        0,
		parent:       this,
	}}, this.properties...)
	for i := 1; i < this.Len(); i++ {
		(this.properties)[i].myIdx = i
	}
}

// PrependJoin prepends a Join value to the front of a list of the property
// "attachment". Invalidates all iterators.
func (this *AttachmentProperty) PrependJoin(v vocab.JoinInterface) {
	this.properties = append([]*AttachmentPropertyIterator{{
		JoinMember: v,
		alias:      this.alias,
		myIdx:      0,
		parent:     this,
	}}, this.properties...)
	for i := 1; i < this.Len(); i++ {
		(this.properties)[i].myIdx = i
	}
}

// PrependLeave prepends a Leave value to the front of a list of the property
// "attachment". Invalidates all iterators.
func (this *AttachmentProperty) PrependLeave(v vocab.LeaveInterface) {
	this.properties = append([]*AttachmentPropertyIterator{{
		LeaveMember: v,
		alias:       this.alias,
		myIdx:       0,
		parent:      this,
	}}, this.properties...)
	for i := 1; i < this.Len(); i++ {
		(this.properties)[i].myIdx = i
	}
}

// PrependLike prepends a Like value to the front of a list of the property
// "attachment". Invalidates all iterators.
func (this *AttachmentProperty) PrependLike(v vocab.LikeInterface) {
	this.properties = append([]*AttachmentPropertyIterator{{
		LikeMember: v,
		alias:      this.alias,
		myIdx:      0,
		parent:     this,
	}}, this.properties...)
	for i := 1; i < this.Len(); i++ {
		(this.properties)[i].myIdx = i
	}
}

// PrependLink prepends a Link value to the front of a list of the property
// "attachment". Invalidates all iterators.
func (this *AttachmentProperty) PrependLink(v vocab.LinkInterface) {
	this.properties = append([]*AttachmentPropertyIterator{{
		LinkMember: v,
		alias:      this.alias,
		myIdx:      0,
		parent:     this,
	}}, this.properties...)
	for i := 1; i < this.Len(); i++ {
		(this.properties)[i].myIdx = i
	}
}

// PrependListen prepends a Listen value to the front of a list of the property
// "attachment". Invalidates all iterators.
func (this *AttachmentProperty) PrependListen(v vocab.ListenInterface) {
	this.properties = append([]*AttachmentPropertyIterator{{
		ListenMember: v,
		alias:        this.alias,
		myIdx:        0,
		parent:       this,
	}}, this.properties...)
	for i := 1; i < this.Len(); i++ {
		(this.properties)[i].myIdx = i
	}
}

// PrependMention prepends a Mention value to the front of a list of the property
// "attachment". Invalidates all iterators.
func (this *AttachmentProperty) PrependMention(v vocab.MentionInterface) {
	this.properties = append([]*AttachmentPropertyIterator{{
		MentionMember: v,
		alias:         this.alias,
		myIdx:         0,
		parent:        this,
	}}, this.properties...)
	for i := 1; i < this.Len(); i++ {
		(this.properties)[i].myIdx = i
	}
}

// PrependMove prepends a Move value to the front of a list of the property
// "attachment". Invalidates all iterators.
func (this *AttachmentProperty) PrependMove(v vocab.MoveInterface) {
	this.properties = append([]*AttachmentPropertyIterator{{
		MoveMember: v,
		alias:      this.alias,
		myIdx:      0,
		parent:     this,
	}}, this.properties...)
	for i := 1; i < this.Len(); i++ {
		(this.properties)[i].myIdx = i
	}
}

// PrependNote prepends a Note value to the front of a list of the property
// "attachment". Invalidates all iterators.
func (this *AttachmentProperty) PrependNote(v vocab.NoteInterface) {
	this.properties = append([]*AttachmentPropertyIterator{{
		NoteMember: v,
		alias:      this.alias,
		myIdx:      0,
		parent:     this,
	}}, this.properties...)
	for i := 1; i < this.Len(); i++ {
		(this.properties)[i].myIdx = i
	}
}

// PrependObject prepends a Object value to the front of a list of the property
// "attachment". Invalidates all iterators.
func (this *AttachmentProperty) PrependObject(v vocab.ObjectInterface) {
	this.properties = append([]*AttachmentPropertyIterator{{
		ObjectMember: v,
		alias:        this.alias,
		myIdx:        0,
		parent:       this,
	}}, this.properties...)
	for i := 1; i < this.Len(); i++ {
		(this.properties)[i].myIdx = i
	}
}

// PrependOffer prepends a Offer value to the front of a list of the property
// "attachment". Invalidates all iterators.
func (this *AttachmentProperty) PrependOffer(v vocab.OfferInterface) {
	this.properties = append([]*AttachmentPropertyIterator{{
		OfferMember: v,
		alias:       this.alias,
		myIdx:       0,
		parent:      this,
	}}, this.properties...)
	for i := 1; i < this.Len(); i++ {
		(this.properties)[i].myIdx = i
	}
}

// PrependOrderedCollection prepends a OrderedCollection value to the front of a
// list of the property "attachment". Invalidates all iterators.
func (this *AttachmentProperty) PrependOrderedCollection(v vocab.OrderedCollectionInterface) {
	this.properties = append([]*AttachmentPropertyIterator{{
		OrderedCollectionMember: v,
		alias:                   this.alias,
		myIdx:                   0,
		parent:                  this,
	}}, this.properties...)
	for i := 1; i < this.Len(); i++ {
		(this.properties)[i].myIdx = i
	}
}

// PrependOrderedCollectionPage prepends a OrderedCollectionPage value to the
// front of a list of the property "attachment". Invalidates all iterators.
func (this *AttachmentProperty) PrependOrderedCollectionPage(v vocab.OrderedCollectionPageInterface) {
	this.properties = append([]*AttachmentPropertyIterator{{
		OrderedCollectionPageMember: v,
		alias:                       this.alias,
		myIdx:                       0,
		parent:                      this,
	}}, this.properties...)
	for i := 1; i < this.Len(); i++ {
		(this.properties)[i].myIdx = i
	}
}

// PrependOrganization prepends a Organization value to the front of a list of the
// property "attachment". Invalidates all iterators.
func (this *AttachmentProperty) PrependOrganization(v vocab.OrganizationInterface) {
	this.properties = append([]*AttachmentPropertyIterator{{
		OrganizationMember: v,
		alias:              this.alias,
		myIdx:              0,
		parent:             this,
	}}, this.properties...)
	for i := 1; i < this.Len(); i++ {
		(this.properties)[i].myIdx = i
	}
}

// PrependPage prepends a Page value to the front of a list of the property
// "attachment". Invalidates all iterators.
func (this *AttachmentProperty) PrependPage(v vocab.PageInterface) {
	this.properties = append([]*AttachmentPropertyIterator{{
		PageMember: v,
		alias:      this.alias,
		myIdx:      0,
		parent:     this,
	}}, this.properties...)
	for i := 1; i < this.Len(); i++ {
		(this.properties)[i].myIdx = i
	}
}

// PrependPerson prepends a Person value to the front of a list of the property
// "attachment". Invalidates all iterators.
func (this *AttachmentProperty) PrependPerson(v vocab.PersonInterface) {
	this.properties = append([]*AttachmentPropertyIterator{{
		PersonMember: v,
		alias:        this.alias,
		myIdx:        0,
		parent:       this,
	}}, this.properties...)
	for i := 1; i < this.Len(); i++ {
		(this.properties)[i].myIdx = i
	}
}

// PrependPlace prepends a Place value to the front of a list of the property
// "attachment". Invalidates all iterators.
func (this *AttachmentProperty) PrependPlace(v vocab.PlaceInterface) {
	this.properties = append([]*AttachmentPropertyIterator{{
		PlaceMember: v,
		alias:       this.alias,
		myIdx:       0,
		parent:      this,
	}}, this.properties...)
	for i := 1; i < this.Len(); i++ {
		(this.properties)[i].myIdx = i
	}
}

// PrependProfile prepends a Profile value to the front of a list of the property
// "attachment". Invalidates all iterators.
func (this *AttachmentProperty) PrependProfile(v vocab.ProfileInterface) {
	this.properties = append([]*AttachmentPropertyIterator{{
		ProfileMember: v,
		alias:         this.alias,
		myIdx:         0,
		parent:        this,
	}}, this.properties...)
	for i := 1; i < this.Len(); i++ {
		(this.properties)[i].myIdx = i
	}
}

// PrependQuestion prepends a Question value to the front of a list of the
// property "attachment". Invalidates all iterators.
func (this *AttachmentProperty) PrependQuestion(v vocab.QuestionInterface) {
	this.properties = append([]*AttachmentPropertyIterator{{
		QuestionMember: v,
		alias:          this.alias,
		myIdx:          0,
		parent:         this,
	}}, this.properties...)
	for i := 1; i < this.Len(); i++ {
		(this.properties)[i].myIdx = i
	}
}

// PrependRead prepends a Read value to the front of a list of the property
// "attachment". Invalidates all iterators.
func (this *AttachmentProperty) PrependRead(v vocab.ReadInterface) {
	this.properties = append([]*AttachmentPropertyIterator{{
		ReadMember: v,
		alias:      this.alias,
		myIdx:      0,
		parent:     this,
	}}, this.properties...)
	for i := 1; i < this.Len(); i++ {
		(this.properties)[i].myIdx = i
	}
}

// PrependReject prepends a Reject value to the front of a list of the property
// "attachment". Invalidates all iterators.
func (this *AttachmentProperty) PrependReject(v vocab.RejectInterface) {
	this.properties = append([]*AttachmentPropertyIterator{{
		RejectMember: v,
		alias:        this.alias,
		myIdx:        0,
		parent:       this,
	}}, this.properties...)
	for i := 1; i < this.Len(); i++ {
		(this.properties)[i].myIdx = i
	}
}

// PrependRelationship prepends a Relationship value to the front of a list of the
// property "attachment". Invalidates all iterators.
func (this *AttachmentProperty) PrependRelationship(v vocab.RelationshipInterface) {
	this.properties = append([]*AttachmentPropertyIterator{{
		RelationshipMember: v,
		alias:              this.alias,
		myIdx:              0,
		parent:             this,
	}}, this.properties...)
	for i := 1; i < this.Len(); i++ {
		(this.properties)[i].myIdx = i
	}
}

// PrependRemove prepends a Remove value to the front of a list of the property
// "attachment". Invalidates all iterators.
func (this *AttachmentProperty) PrependRemove(v vocab.RemoveInterface) {
	this.properties = append([]*AttachmentPropertyIterator{{
		RemoveMember: v,
		alias:        this.alias,
		myIdx:        0,
		parent:       this,
	}}, this.properties...)
	for i := 1; i < this.Len(); i++ {
		(this.properties)[i].myIdx = i
	}
}

// PrependService prepends a Service value to the front of a list of the property
// "attachment". Invalidates all iterators.
func (this *AttachmentProperty) PrependService(v vocab.ServiceInterface) {
	this.properties = append([]*AttachmentPropertyIterator{{
		ServiceMember: v,
		alias:         this.alias,
		myIdx:         0,
		parent:        this,
	}}, this.properties...)
	for i := 1; i < this.Len(); i++ {
		(this.properties)[i].myIdx = i
	}
}

// PrependTentativeAccept prepends a TentativeAccept value to the front of a list
// of the property "attachment". Invalidates all iterators.
func (this *AttachmentProperty) PrependTentativeAccept(v vocab.TentativeAcceptInterface) {
	this.properties = append([]*AttachmentPropertyIterator{{
		TentativeAcceptMember: v,
		alias:                 this.alias,
		myIdx:                 0,
		parent:                this,
	}}, this.properties...)
	for i := 1; i < this.Len(); i++ {
		(this.properties)[i].myIdx = i
	}
}

// PrependTentativeReject prepends a TentativeReject value to the front of a list
// of the property "attachment". Invalidates all iterators.
func (this *AttachmentProperty) PrependTentativeReject(v vocab.TentativeRejectInterface) {
	this.properties = append([]*AttachmentPropertyIterator{{
		TentativeRejectMember: v,
		alias:                 this.alias,
		myIdx:                 0,
		parent:                this,
	}}, this.properties...)
	for i := 1; i < this.Len(); i++ {
		(this.properties)[i].myIdx = i
	}
}

// PrependTombstone prepends a Tombstone value to the front of a list of the
// property "attachment". Invalidates all iterators.
func (this *AttachmentProperty) PrependTombstone(v vocab.TombstoneInterface) {
	this.properties = append([]*AttachmentPropertyIterator{{
		TombstoneMember: v,
		alias:           this.alias,
		myIdx:           0,
		parent:          this,
	}}, this.properties...)
	for i := 1; i < this.Len(); i++ {
		(this.properties)[i].myIdx = i
	}
}

// PrependTravel prepends a Travel value to the front of a list of the property
// "attachment". Invalidates all iterators.
func (this *AttachmentProperty) PrependTravel(v vocab.TravelInterface) {
	this.properties = append([]*AttachmentPropertyIterator{{
		TravelMember: v,
		alias:        this.alias,
		myIdx:        0,
		parent:       this,
	}}, this.properties...)
	for i := 1; i < this.Len(); i++ {
		(this.properties)[i].myIdx = i
	}
}

// PrependUndo prepends a Undo value to the front of a list of the property
// "attachment". Invalidates all iterators.
func (this *AttachmentProperty) PrependUndo(v vocab.UndoInterface) {
	this.properties = append([]*AttachmentPropertyIterator{{
		UndoMember: v,
		alias:      this.alias,
		myIdx:      0,
		parent:     this,
	}}, this.properties...)
	for i := 1; i < this.Len(); i++ {
		(this.properties)[i].myIdx = i
	}
}

// PrependUpdate prepends a Update value to the front of a list of the property
// "attachment". Invalidates all iterators.
func (this *AttachmentProperty) PrependUpdate(v vocab.UpdateInterface) {
	this.properties = append([]*AttachmentPropertyIterator{{
		UpdateMember: v,
		alias:        this.alias,
		myIdx:        0,
		parent:       this,
	}}, this.properties...)
	for i := 1; i < this.Len(); i++ {
		(this.properties)[i].myIdx = i
	}
}

// PrependVideo prepends a Video value to the front of a list of the property
// "attachment". Invalidates all iterators.
func (this *AttachmentProperty) PrependVideo(v vocab.VideoInterface) {
	this.properties = append([]*AttachmentPropertyIterator{{
		VideoMember: v,
		alias:       this.alias,
		myIdx:       0,
		parent:      this,
	}}, this.properties...)
	for i := 1; i < this.Len(); i++ {
		(this.properties)[i].myIdx = i
	}
}

// PrependView prepends a View value to the front of a list of the property
// "attachment". Invalidates all iterators.
func (this *AttachmentProperty) PrependView(v vocab.ViewInterface) {
	this.properties = append([]*AttachmentPropertyIterator{{
		ViewMember: v,
		alias:      this.alias,
		myIdx:      0,
		parent:     this,
	}}, this.properties...)
	for i := 1; i < this.Len(); i++ {
		(this.properties)[i].myIdx = i
	}
}

// Remove deletes an element at the specified index from a list of the property
// "attachment", regardless of its type. Panics if the index is out of bounds.
// Invalidates all iterators.
func (this *AttachmentProperty) Remove(idx int) {
	(this.properties)[idx].parent = nil
	copy((this.properties)[idx:], (this.properties)[idx+1:])
	(this.properties)[len(this.properties)-1] = &AttachmentPropertyIterator{}
	this.properties = (this.properties)[:len(this.properties)-1]
	for i := idx; i < this.Len(); i++ {
		(this.properties)[i].myIdx = i
	}
}

// Serialize converts this into an interface representation suitable for
// marshalling into a text or binary format. Applications should not need this
// function as most typical use cases serialize types instead of individual
// properties. It is exposed for alternatives to go-fed implementations to use.
func (this AttachmentProperty) Serialize() (interface{}, error) {
	s := make([]interface{}, 0, len(this.properties))
	for _, iterator := range this.properties {
		if b, err := iterator.serialize(); err != nil {
			return s, err
		} else {
			s = append(s, b)
		}
	}
	// Shortcut: if serializing one value, don't return an array -- pretty sure other Fediverse software would choke on a "type" value with array, for example.
	if len(s) == 1 {
		return s[0], nil
	}
	return s, nil
}

// SetAccept sets a Accept value to be at the specified index for the property
// "attachment". Panics if the index is out of bounds. Invalidates all
// iterators.
func (this *AttachmentProperty) SetAccept(idx int, v vocab.AcceptInterface) {
	(this.properties)[idx].parent = nil
	(this.properties)[idx] = &AttachmentPropertyIterator{
		AcceptMember: v,
		alias:        this.alias,
		myIdx:        idx,
		parent:       this,
	}
}

// SetActivity sets a Activity value to be at the specified index for the property
// "attachment". Panics if the index is out of bounds. Invalidates all
// iterators.
func (this *AttachmentProperty) SetActivity(idx int, v vocab.ActivityInterface) {
	(this.properties)[idx].parent = nil
	(this.properties)[idx] = &AttachmentPropertyIterator{
		ActivityMember: v,
		alias:          this.alias,
		myIdx:          idx,
		parent:         this,
	}
}

// SetAdd sets a Add value to be at the specified index for the property
// "attachment". Panics if the index is out of bounds. Invalidates all
// iterators.
func (this *AttachmentProperty) SetAdd(idx int, v vocab.AddInterface) {
	(this.properties)[idx].parent = nil
	(this.properties)[idx] = &AttachmentPropertyIterator{
		AddMember: v,
		alias:     this.alias,
		myIdx:     idx,
		parent:    this,
	}
}

// SetAnnounce sets a Announce value to be at the specified index for the property
// "attachment". Panics if the index is out of bounds. Invalidates all
// iterators.
func (this *AttachmentProperty) SetAnnounce(idx int, v vocab.AnnounceInterface) {
	(this.properties)[idx].parent = nil
	(this.properties)[idx] = &AttachmentPropertyIterator{
		AnnounceMember: v,
		alias:          this.alias,
		myIdx:          idx,
		parent:         this,
	}
}

// SetApplication sets a Application value to be at the specified index for the
// property "attachment". Panics if the index is out of bounds. Invalidates
// all iterators.
func (this *AttachmentProperty) SetApplication(idx int, v vocab.ApplicationInterface) {
	(this.properties)[idx].parent = nil
	(this.properties)[idx] = &AttachmentPropertyIterator{
		ApplicationMember: v,
		alias:             this.alias,
		myIdx:             idx,
		parent:            this,
	}
}

// SetArrive sets a Arrive value to be at the specified index for the property
// "attachment". Panics if the index is out of bounds. Invalidates all
// iterators.
func (this *AttachmentProperty) SetArrive(idx int, v vocab.ArriveInterface) {
	(this.properties)[idx].parent = nil
	(this.properties)[idx] = &AttachmentPropertyIterator{
		ArriveMember: v,
		alias:        this.alias,
		myIdx:        idx,
		parent:       this,
	}
}

// SetArticle sets a Article value to be at the specified index for the property
// "attachment". Panics if the index is out of bounds. Invalidates all
// iterators.
func (this *AttachmentProperty) SetArticle(idx int, v vocab.ArticleInterface) {
	(this.properties)[idx].parent = nil
	(this.properties)[idx] = &AttachmentPropertyIterator{
		ArticleMember: v,
		alias:         this.alias,
		myIdx:         idx,
		parent:        this,
	}
}

// SetAudio sets a Audio value to be at the specified index for the property
// "attachment". Panics if the index is out of bounds. Invalidates all
// iterators.
func (this *AttachmentProperty) SetAudio(idx int, v vocab.AudioInterface) {
	(this.properties)[idx].parent = nil
	(this.properties)[idx] = &AttachmentPropertyIterator{
		AudioMember: v,
		alias:       this.alias,
		myIdx:       idx,
		parent:      this,
	}
}

// SetBlock sets a Block value to be at the specified index for the property
// "attachment". Panics if the index is out of bounds. Invalidates all
// iterators.
func (this *AttachmentProperty) SetBlock(idx int, v vocab.BlockInterface) {
	(this.properties)[idx].parent = nil
	(this.properties)[idx] = &AttachmentPropertyIterator{
		BlockMember: v,
		alias:       this.alias,
		myIdx:       idx,
		parent:      this,
	}
}

// SetCollection sets a Collection value to be at the specified index for the
// property "attachment". Panics if the index is out of bounds. Invalidates
// all iterators.
func (this *AttachmentProperty) SetCollection(idx int, v vocab.CollectionInterface) {
	(this.properties)[idx].parent = nil
	(this.properties)[idx] = &AttachmentPropertyIterator{
		CollectionMember: v,
		alias:            this.alias,
		myIdx:            idx,
		parent:           this,
	}
}

// SetCollectionPage sets a CollectionPage value to be at the specified index for
// the property "attachment". Panics if the index is out of bounds.
// Invalidates all iterators.
func (this *AttachmentProperty) SetCollectionPage(idx int, v vocab.CollectionPageInterface) {
	(this.properties)[idx].parent = nil
	(this.properties)[idx] = &AttachmentPropertyIterator{
		CollectionPageMember: v,
		alias:                this.alias,
		myIdx:                idx,
		parent:               this,
	}
}

// SetCreate sets a Create value to be at the specified index for the property
// "attachment". Panics if the index is out of bounds. Invalidates all
// iterators.
func (this *AttachmentProperty) SetCreate(idx int, v vocab.CreateInterface) {
	(this.properties)[idx].parent = nil
	(this.properties)[idx] = &AttachmentPropertyIterator{
		CreateMember: v,
		alias:        this.alias,
		myIdx:        idx,
		parent:       this,
	}
}

// SetDelete sets a Delete value to be at the specified index for the property
// "attachment". Panics if the index is out of bounds. Invalidates all
// iterators.
func (this *AttachmentProperty) SetDelete(idx int, v vocab.DeleteInterface) {
	(this.properties)[idx].parent = nil
	(this.properties)[idx] = &AttachmentPropertyIterator{
		DeleteMember: v,
		alias:        this.alias,
		myIdx:        idx,
		parent:       this,
	}
}

// SetDislike sets a Dislike value to be at the specified index for the property
// "attachment". Panics if the index is out of bounds. Invalidates all
// iterators.
func (this *AttachmentProperty) SetDislike(idx int, v vocab.DislikeInterface) {
	(this.properties)[idx].parent = nil
	(this.properties)[idx] = &AttachmentPropertyIterator{
		DislikeMember: v,
		alias:         this.alias,
		myIdx:         idx,
		parent:        this,
	}
}

// SetDocument sets a Document value to be at the specified index for the property
// "attachment". Panics if the index is out of bounds. Invalidates all
// iterators.
func (this *AttachmentProperty) SetDocument(idx int, v vocab.DocumentInterface) {
	(this.properties)[idx].parent = nil
	(this.properties)[idx] = &AttachmentPropertyIterator{
		DocumentMember: v,
		alias:          this.alias,
		myIdx:          idx,
		parent:         this,
	}
}

// SetEvent sets a Event value to be at the specified index for the property
// "attachment". Panics if the index is out of bounds. Invalidates all
// iterators.
func (this *AttachmentProperty) SetEvent(idx int, v vocab.EventInterface) {
	(this.properties)[idx].parent = nil
	(this.properties)[idx] = &AttachmentPropertyIterator{
		EventMember: v,
		alias:       this.alias,
		myIdx:       idx,
		parent:      this,
	}
}

// SetFlag sets a Flag value to be at the specified index for the property
// "attachment". Panics if the index is out of bounds. Invalidates all
// iterators.
func (this *AttachmentProperty) SetFlag(idx int, v vocab.FlagInterface) {
	(this.properties)[idx].parent = nil
	(this.properties)[idx] = &AttachmentPropertyIterator{
		FlagMember: v,
		alias:      this.alias,
		myIdx:      idx,
		parent:     this,
	}
}

// SetFollow sets a Follow value to be at the specified index for the property
// "attachment". Panics if the index is out of bounds. Invalidates all
// iterators.
func (this *AttachmentProperty) SetFollow(idx int, v vocab.FollowInterface) {
	(this.properties)[idx].parent = nil
	(this.properties)[idx] = &AttachmentPropertyIterator{
		FollowMember: v,
		alias:        this.alias,
		myIdx:        idx,
		parent:       this,
	}
}

// SetGroup sets a Group value to be at the specified index for the property
// "attachment". Panics if the index is out of bounds. Invalidates all
// iterators.
func (this *AttachmentProperty) SetGroup(idx int, v vocab.GroupInterface) {
	(this.properties)[idx].parent = nil
	(this.properties)[idx] = &AttachmentPropertyIterator{
		GroupMember: v,
		alias:       this.alias,
		myIdx:       idx,
		parent:      this,
	}
}

// SetIRI sets an IRI value to be at the specified index for the property
// "attachment". Panics if the index is out of bounds.
func (this *AttachmentProperty) SetIRI(idx int, v *url.URL) {
	(this.properties)[idx].parent = nil
	(this.properties)[idx] = &AttachmentPropertyIterator{
		alias:  this.alias,
		iri:    v,
		myIdx:  idx,
		parent: this,
	}
}

// SetIgnore sets a Ignore value to be at the specified index for the property
// "attachment". Panics if the index is out of bounds. Invalidates all
// iterators.
func (this *AttachmentProperty) SetIgnore(idx int, v vocab.IgnoreInterface) {
	(this.properties)[idx].parent = nil
	(this.properties)[idx] = &AttachmentPropertyIterator{
		IgnoreMember: v,
		alias:        this.alias,
		myIdx:        idx,
		parent:       this,
	}
}

// SetImage sets a Image value to be at the specified index for the property
// "attachment". Panics if the index is out of bounds. Invalidates all
// iterators.
func (this *AttachmentProperty) SetImage(idx int, v vocab.ImageInterface) {
	(this.properties)[idx].parent = nil
	(this.properties)[idx] = &AttachmentPropertyIterator{
		ImageMember: v,
		alias:       this.alias,
		myIdx:       idx,
		parent:      this,
	}
}

// SetIntransitiveActivity sets a IntransitiveActivity value to be at the
// specified index for the property "attachment". Panics if the index is out
// of bounds. Invalidates all iterators.
func (this *AttachmentProperty) SetIntransitiveActivity(idx int, v vocab.IntransitiveActivityInterface) {
	(this.properties)[idx].parent = nil
	(this.properties)[idx] = &AttachmentPropertyIterator{
		IntransitiveActivityMember: v,
		alias:                      this.alias,
		myIdx:                      idx,
		parent:                     this,
	}
}

// SetInvite sets a Invite value to be at the specified index for the property
// "attachment". Panics if the index is out of bounds. Invalidates all
// iterators.
func (this *AttachmentProperty) SetInvite(idx int, v vocab.InviteInterface) {
	(this.properties)[idx].parent = nil
	(this.properties)[idx] = &AttachmentPropertyIterator{
		InviteMember: v,
		alias:        this.alias,
		myIdx:        idx,
		parent:       this,
	}
}

// SetJoin sets a Join value to be at the specified index for the property
// "attachment". Panics if the index is out of bounds. Invalidates all
// iterators.
func (this *AttachmentProperty) SetJoin(idx int, v vocab.JoinInterface) {
	(this.properties)[idx].parent = nil
	(this.properties)[idx] = &AttachmentPropertyIterator{
		JoinMember: v,
		alias:      this.alias,
		myIdx:      idx,
		parent:     this,
	}
}

// SetLeave sets a Leave value to be at the specified index for the property
// "attachment". Panics if the index is out of bounds. Invalidates all
// iterators.
func (this *AttachmentProperty) SetLeave(idx int, v vocab.LeaveInterface) {
	(this.properties)[idx].parent = nil
	(this.properties)[idx] = &AttachmentPropertyIterator{
		LeaveMember: v,
		alias:       this.alias,
		myIdx:       idx,
		parent:      this,
	}
}

// SetLike sets a Like value to be at the specified index for the property
// "attachment". Panics if the index is out of bounds. Invalidates all
// iterators.
func (this *AttachmentProperty) SetLike(idx int, v vocab.LikeInterface) {
	(this.properties)[idx].parent = nil
	(this.properties)[idx] = &AttachmentPropertyIterator{
		LikeMember: v,
		alias:      this.alias,
		myIdx:      idx,
		parent:     this,
	}
}

// SetLink sets a Link value to be at the specified index for the property
// "attachment". Panics if the index is out of bounds. Invalidates all
// iterators.
func (this *AttachmentProperty) SetLink(idx int, v vocab.LinkInterface) {
	(this.properties)[idx].parent = nil
	(this.properties)[idx] = &AttachmentPropertyIterator{
		LinkMember: v,
		alias:      this.alias,
		myIdx:      idx,
		parent:     this,
	}
}

// SetListen sets a Listen value to be at the specified index for the property
// "attachment". Panics if the index is out of bounds. Invalidates all
// iterators.
func (this *AttachmentProperty) SetListen(idx int, v vocab.ListenInterface) {
	(this.properties)[idx].parent = nil
	(this.properties)[idx] = &AttachmentPropertyIterator{
		ListenMember: v,
		alias:        this.alias,
		myIdx:        idx,
		parent:       this,
	}
}

// SetMention sets a Mention value to be at the specified index for the property
// "attachment". Panics if the index is out of bounds. Invalidates all
// iterators.
func (this *AttachmentProperty) SetMention(idx int, v vocab.MentionInterface) {
	(this.properties)[idx].parent = nil
	(this.properties)[idx] = &AttachmentPropertyIterator{
		MentionMember: v,
		alias:         this.alias,
		myIdx:         idx,
		parent:        this,
	}
}

// SetMove sets a Move value to be at the specified index for the property
// "attachment". Panics if the index is out of bounds. Invalidates all
// iterators.
func (this *AttachmentProperty) SetMove(idx int, v vocab.MoveInterface) {
	(this.properties)[idx].parent = nil
	(this.properties)[idx] = &AttachmentPropertyIterator{
		MoveMember: v,
		alias:      this.alias,
		myIdx:      idx,
		parent:     this,
	}
}

// SetNote sets a Note value to be at the specified index for the property
// "attachment". Panics if the index is out of bounds. Invalidates all
// iterators.
func (this *AttachmentProperty) SetNote(idx int, v vocab.NoteInterface) {
	(this.properties)[idx].parent = nil
	(this.properties)[idx] = &AttachmentPropertyIterator{
		NoteMember: v,
		alias:      this.alias,
		myIdx:      idx,
		parent:     this,
	}
}

// SetObject sets a Object value to be at the specified index for the property
// "attachment". Panics if the index is out of bounds. Invalidates all
// iterators.
func (this *AttachmentProperty) SetObject(idx int, v vocab.ObjectInterface) {
	(this.properties)[idx].parent = nil
	(this.properties)[idx] = &AttachmentPropertyIterator{
		ObjectMember: v,
		alias:        this.alias,
		myIdx:        idx,
		parent:       this,
	}
}

// SetOffer sets a Offer value to be at the specified index for the property
// "attachment". Panics if the index is out of bounds. Invalidates all
// iterators.
func (this *AttachmentProperty) SetOffer(idx int, v vocab.OfferInterface) {
	(this.properties)[idx].parent = nil
	(this.properties)[idx] = &AttachmentPropertyIterator{
		OfferMember: v,
		alias:       this.alias,
		myIdx:       idx,
		parent:      this,
	}
}

// SetOrderedCollection sets a OrderedCollection value to be at the specified
// index for the property "attachment". Panics if the index is out of bounds.
// Invalidates all iterators.
func (this *AttachmentProperty) SetOrderedCollection(idx int, v vocab.OrderedCollectionInterface) {
	(this.properties)[idx].parent = nil
	(this.properties)[idx] = &AttachmentPropertyIterator{
		OrderedCollectionMember: v,
		alias:                   this.alias,
		myIdx:                   idx,
		parent:                  this,
	}
}

// SetOrderedCollectionPage sets a OrderedCollectionPage value to be at the
// specified index for the property "attachment". Panics if the index is out
// of bounds. Invalidates all iterators.
func (this *AttachmentProperty) SetOrderedCollectionPage(idx int, v vocab.OrderedCollectionPageInterface) {
	(this.properties)[idx].parent = nil
	(this.properties)[idx] = &AttachmentPropertyIterator{
		OrderedCollectionPageMember: v,
		alias:                       this.alias,
		myIdx:                       idx,
		parent:                      this,
	}
}

// SetOrganization sets a Organization value to be at the specified index for the
// property "attachment". Panics if the index is out of bounds. Invalidates
// all iterators.
func (this *AttachmentProperty) SetOrganization(idx int, v vocab.OrganizationInterface) {
	(this.properties)[idx].parent = nil
	(this.properties)[idx] = &AttachmentPropertyIterator{
		OrganizationMember: v,
		alias:              this.alias,
		myIdx:              idx,
		parent:             this,
	}
}

// SetPage sets a Page value to be at the specified index for the property
// "attachment". Panics if the index is out of bounds. Invalidates all
// iterators.
func (this *AttachmentProperty) SetPage(idx int, v vocab.PageInterface) {
	(this.properties)[idx].parent = nil
	(this.properties)[idx] = &AttachmentPropertyIterator{
		PageMember: v,
		alias:      this.alias,
		myIdx:      idx,
		parent:     this,
	}
}

// SetPerson sets a Person value to be at the specified index for the property
// "attachment". Panics if the index is out of bounds. Invalidates all
// iterators.
func (this *AttachmentProperty) SetPerson(idx int, v vocab.PersonInterface) {
	(this.properties)[idx].parent = nil
	(this.properties)[idx] = &AttachmentPropertyIterator{
		PersonMember: v,
		alias:        this.alias,
		myIdx:        idx,
		parent:       this,
	}
}

// SetPlace sets a Place value to be at the specified index for the property
// "attachment". Panics if the index is out of bounds. Invalidates all
// iterators.
func (this *AttachmentProperty) SetPlace(idx int, v vocab.PlaceInterface) {
	(this.properties)[idx].parent = nil
	(this.properties)[idx] = &AttachmentPropertyIterator{
		PlaceMember: v,
		alias:       this.alias,
		myIdx:       idx,
		parent:      this,
	}
}

// SetProfile sets a Profile value to be at the specified index for the property
// "attachment". Panics if the index is out of bounds. Invalidates all
// iterators.
func (this *AttachmentProperty) SetProfile(idx int, v vocab.ProfileInterface) {
	(this.properties)[idx].parent = nil
	(this.properties)[idx] = &AttachmentPropertyIterator{
		ProfileMember: v,
		alias:         this.alias,
		myIdx:         idx,
		parent:        this,
	}
}

// SetQuestion sets a Question value to be at the specified index for the property
// "attachment". Panics if the index is out of bounds. Invalidates all
// iterators.
func (this *AttachmentProperty) SetQuestion(idx int, v vocab.QuestionInterface) {
	(this.properties)[idx].parent = nil
	(this.properties)[idx] = &AttachmentPropertyIterator{
		QuestionMember: v,
		alias:          this.alias,
		myIdx:          idx,
		parent:         this,
	}
}

// SetRead sets a Read value to be at the specified index for the property
// "attachment". Panics if the index is out of bounds. Invalidates all
// iterators.
func (this *AttachmentProperty) SetRead(idx int, v vocab.ReadInterface) {
	(this.properties)[idx].parent = nil
	(this.properties)[idx] = &AttachmentPropertyIterator{
		ReadMember: v,
		alias:      this.alias,
		myIdx:      idx,
		parent:     this,
	}
}

// SetReject sets a Reject value to be at the specified index for the property
// "attachment". Panics if the index is out of bounds. Invalidates all
// iterators.
func (this *AttachmentProperty) SetReject(idx int, v vocab.RejectInterface) {
	(this.properties)[idx].parent = nil
	(this.properties)[idx] = &AttachmentPropertyIterator{
		RejectMember: v,
		alias:        this.alias,
		myIdx:        idx,
		parent:       this,
	}
}

// SetRelationship sets a Relationship value to be at the specified index for the
// property "attachment". Panics if the index is out of bounds. Invalidates
// all iterators.
func (this *AttachmentProperty) SetRelationship(idx int, v vocab.RelationshipInterface) {
	(this.properties)[idx].parent = nil
	(this.properties)[idx] = &AttachmentPropertyIterator{
		RelationshipMember: v,
		alias:              this.alias,
		myIdx:              idx,
		parent:             this,
	}
}

// SetRemove sets a Remove value to be at the specified index for the property
// "attachment". Panics if the index is out of bounds. Invalidates all
// iterators.
func (this *AttachmentProperty) SetRemove(idx int, v vocab.RemoveInterface) {
	(this.properties)[idx].parent = nil
	(this.properties)[idx] = &AttachmentPropertyIterator{
		RemoveMember: v,
		alias:        this.alias,
		myIdx:        idx,
		parent:       this,
	}
}

// SetService sets a Service value to be at the specified index for the property
// "attachment". Panics if the index is out of bounds. Invalidates all
// iterators.
func (this *AttachmentProperty) SetService(idx int, v vocab.ServiceInterface) {
	(this.properties)[idx].parent = nil
	(this.properties)[idx] = &AttachmentPropertyIterator{
		ServiceMember: v,
		alias:         this.alias,
		myIdx:         idx,
		parent:        this,
	}
}

// SetTentativeAccept sets a TentativeAccept value to be at the specified index
// for the property "attachment". Panics if the index is out of bounds.
// Invalidates all iterators.
func (this *AttachmentProperty) SetTentativeAccept(idx int, v vocab.TentativeAcceptInterface) {
	(this.properties)[idx].parent = nil
	(this.properties)[idx] = &AttachmentPropertyIterator{
		TentativeAcceptMember: v,
		alias:                 this.alias,
		myIdx:                 idx,
		parent:                this,
	}
}

// SetTentativeReject sets a TentativeReject value to be at the specified index
// for the property "attachment". Panics if the index is out of bounds.
// Invalidates all iterators.
func (this *AttachmentProperty) SetTentativeReject(idx int, v vocab.TentativeRejectInterface) {
	(this.properties)[idx].parent = nil
	(this.properties)[idx] = &AttachmentPropertyIterator{
		TentativeRejectMember: v,
		alias:                 this.alias,
		myIdx:                 idx,
		parent:                this,
	}
}

// SetTombstone sets a Tombstone value to be at the specified index for the
// property "attachment". Panics if the index is out of bounds. Invalidates
// all iterators.
func (this *AttachmentProperty) SetTombstone(idx int, v vocab.TombstoneInterface) {
	(this.properties)[idx].parent = nil
	(this.properties)[idx] = &AttachmentPropertyIterator{
		TombstoneMember: v,
		alias:           this.alias,
		myIdx:           idx,
		parent:          this,
	}
}

// SetTravel sets a Travel value to be at the specified index for the property
// "attachment". Panics if the index is out of bounds. Invalidates all
// iterators.
func (this *AttachmentProperty) SetTravel(idx int, v vocab.TravelInterface) {
	(this.properties)[idx].parent = nil
	(this.properties)[idx] = &AttachmentPropertyIterator{
		TravelMember: v,
		alias:        this.alias,
		myIdx:        idx,
		parent:       this,
	}
}

// SetUndo sets a Undo value to be at the specified index for the property
// "attachment". Panics if the index is out of bounds. Invalidates all
// iterators.
func (this *AttachmentProperty) SetUndo(idx int, v vocab.UndoInterface) {
	(this.properties)[idx].parent = nil
	(this.properties)[idx] = &AttachmentPropertyIterator{
		UndoMember: v,
		alias:      this.alias,
		myIdx:      idx,
		parent:     this,
	}
}

// SetUpdate sets a Update value to be at the specified index for the property
// "attachment". Panics if the index is out of bounds. Invalidates all
// iterators.
func (this *AttachmentProperty) SetUpdate(idx int, v vocab.UpdateInterface) {
	(this.properties)[idx].parent = nil
	(this.properties)[idx] = &AttachmentPropertyIterator{
		UpdateMember: v,
		alias:        this.alias,
		myIdx:        idx,
		parent:       this,
	}
}

// SetVideo sets a Video value to be at the specified index for the property
// "attachment". Panics if the index is out of bounds. Invalidates all
// iterators.
func (this *AttachmentProperty) SetVideo(idx int, v vocab.VideoInterface) {
	(this.properties)[idx].parent = nil
	(this.properties)[idx] = &AttachmentPropertyIterator{
		VideoMember: v,
		alias:       this.alias,
		myIdx:       idx,
		parent:      this,
	}
}

// SetView sets a View value to be at the specified index for the property
// "attachment". Panics if the index is out of bounds. Invalidates all
// iterators.
func (this *AttachmentProperty) SetView(idx int, v vocab.ViewInterface) {
	(this.properties)[idx].parent = nil
	(this.properties)[idx] = &AttachmentPropertyIterator{
		ViewMember: v,
		alias:      this.alias,
		myIdx:      idx,
		parent:     this,
	}
}

// Swap swaps the location of values at two indices for the "attachment" property.
func (this AttachmentProperty) Swap(i, j int) {
	this.properties[i], this.properties[j] = this.properties[j], this.properties[i]
}