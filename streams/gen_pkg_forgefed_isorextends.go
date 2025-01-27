// Code generated by astool. DO NOT EDIT.

package streams

import (
	typebranch "github.com/poast-social/activity/streams/impl/forgefed/type_branch"
	typecommit "github.com/poast-social/activity/streams/impl/forgefed/type_commit"
	typepush "github.com/poast-social/activity/streams/impl/forgefed/type_push"
	typerepository "github.com/poast-social/activity/streams/impl/forgefed/type_repository"
	typeticket "github.com/poast-social/activity/streams/impl/forgefed/type_ticket"
	typeticketdependency "github.com/poast-social/activity/streams/impl/forgefed/type_ticketdependency"
	vocab "github.com/poast-social/activity/streams/vocab"
)

// IsOrExtendsForgeFedBranch returns true if the other provided type is the Branch
// type or extends from the Branch type.
func IsOrExtendsForgeFedBranch(other vocab.Type) bool {
	return typebranch.IsOrExtendsBranch(other)
}

// IsOrExtendsForgeFedCommit returns true if the other provided type is the Commit
// type or extends from the Commit type.
func IsOrExtendsForgeFedCommit(other vocab.Type) bool {
	return typecommit.IsOrExtendsCommit(other)
}

// IsOrExtendsForgeFedPush returns true if the other provided type is the Push
// type or extends from the Push type.
func IsOrExtendsForgeFedPush(other vocab.Type) bool {
	return typepush.IsOrExtendsPush(other)
}

// IsOrExtendsForgeFedRepository returns true if the other provided type is the
// Repository type or extends from the Repository type.
func IsOrExtendsForgeFedRepository(other vocab.Type) bool {
	return typerepository.IsOrExtendsRepository(other)
}

// IsOrExtendsForgeFedTicket returns true if the other provided type is the Ticket
// type or extends from the Ticket type.
func IsOrExtendsForgeFedTicket(other vocab.Type) bool {
	return typeticket.IsOrExtendsTicket(other)
}

// IsOrExtendsForgeFedTicketDependency returns true if the other provided type is
// the TicketDependency type or extends from the TicketDependency type.
func IsOrExtendsForgeFedTicketDependency(other vocab.Type) bool {
	return typeticketdependency.IsOrExtendsTicketDependency(other)
}
