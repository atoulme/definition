/**
 * Copyright 2019 Whiteblock Inc. All rights reserved.
 * Use of this source code is governed by a BSD-style
 * license that can be found in the LICENSE file.
 */

package entity

import (
	"errors"

	"github.com/whiteblock/definition/config"

	"github.com/sirupsen/logrus"
)

type ResourceBuckets interface {
	Add(segments []Segment) error
	Remove(segments []Segment) error
	Resources() []Bucket
}

var (
	ErrSizeLimitExceeded = errors.New("size limit exceeded")
	ErrSegmentTooLarge   = errors.New("segment size too large")
	ErrSegmentNotFound   = errors.New("segment not found")
)

type resourceBuckets struct {
	conf    config.Bucket
	buckets []*Bucket
	log     logrus.Ext1FieldLogger
}

func NewResourceBuckets(conf config.Bucket, log logrus.Ext1FieldLogger) ResourceBuckets {
	return &resourceBuckets{conf: conf, log: log}
}

func (rb *resourceBuckets) add(segment Segment) error {
	for i := range rb.buckets {
		if rb.buckets[i].tryAdd(segment) {
			return nil
		}
	}
	if int64(len(rb.buckets)) == rb.conf.MaxBuckets {
		return ErrSizeLimitExceeded
	}
	bucket := NewBucket(&rb.conf, rb.log)
	if !bucket.tryAdd(segment) {
		return ErrSegmentTooLarge
	}
	rb.buckets = append(rb.buckets, bucket)
	return nil
}

func (rb *resourceBuckets) Add(segments []Segment) error {
	for _, segment := range segments {
		err := rb.add(segment)
		if err != nil {
			rb.log.WithFields(logrus.Fields{
				"segment": segment,
				"error":   err}).Warn("failed to add segment")
			return err
		}
	}
	return nil
}

func (rb *resourceBuckets) remove(segment Segment) error {
	for i := range rb.buckets {
		if rb.buckets[i].tryRemove(segment) {
			return nil
		}
	}
	return ErrSegmentNotFound
}

func (rb *resourceBuckets) Remove(segments []Segment) error {
	for _, segment := range segments {
		err := rb.remove(segment)
		if err != nil {
			rb.log.WithFields(logrus.Fields{
				"segment": segment,
				"error":   err}).Warn("failed to remove segment")
			return err
		}
	}
	return nil
}

func (rb *resourceBuckets) Resources() []Bucket {
	out := make([]Bucket, len(rb.buckets))
	for i := range out {
		out[i] = rb.buckets[i].Clone()
	}
	return out
}
