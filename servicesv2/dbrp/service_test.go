package dbrp_test

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	influxdb "github.com/influxdata/influxdb/servicesv2"
	"github.com/influxdata/influxdb/servicesv2/bolt"
	"github.com/influxdata/influxdb/servicesv2/dbrp"
	"github.com/influxdata/influxdb/servicesv2/kv"
	"github.com/influxdata/influxdb/servicesv2/mock"
	itesting "github.com/influxdata/influxdb/servicesv2/testing"
	"go.uber.org/zap/zaptest"
)

var (
	bucket        = []byte("dbrpv1")
	indexBucket   = []byte("dbrpbyorganddbindexv1")
	defaultBucket = []byte("dbrpdefaultv1")
)

func NewTestBoltStore(t *testing.T) (kv.Store, func(), error) {
	t.Helper()

	f, err := ioutil.TempFile("", "influxdata-bolt-")
	if err != nil {
		return nil, nil, errors.New("unable to open temporary boltdb file")
	}
	f.Close()

	//ctx := context.Background() //ROHAN
	logger := zaptest.NewLogger(t)
	path := f.Name()
	s := bolt.NewKVStore(logger, path)
	if err := s.Open(context.Background()); err != nil {
		return nil, nil, err
	}

	// if err := all.Up(ctx, logger, s); err != nil {
	// 	return nil, nil, err
	// }

	buckets := [][]byte{
		bucket,
		indexBucket,
		defaultBucket,
	}

	for _, b := range buckets {
		err = s.CreateBucket(context.Background(), b)
		if err != nil {
			t.Fatalf("Cannot create bucket: %v", err)
		}
	}

	close := func() {
		s.Close()
		os.Remove(path)
	}

	return s, close, nil
}

func initDBRPMappingService(f itesting.DBRPMappingFieldsV2, t *testing.T) (influxdb.DBRPMappingServiceV2, func()) {
	s, closeStore, err := NewTestBoltStore(t)
	if err != nil {
		t.Fatalf("failed to create new bolt kv store: %v", err)
	}

	if f.BucketSvc == nil {
		f.BucketSvc = &mock.BucketService{
			FindBucketByIDFn: func(ctx context.Context, id influxdb.ID) (*influxdb.Bucket, error) {
				// always find a bucket.
				return &influxdb.Bucket{
					ID:   id,
					Name: fmt.Sprintf("bucket-%v", id),
				}, nil
			},
		}
	}

	svc := dbrp.NewService(context.Background(), f.BucketSvc, s) ///issue

	if err := f.Populate(context.Background(), svc); err != nil {
		t.Fatalf("I ERROED: %v", err)
	}
	return svc, func() {
		if err := itesting.CleanupDBRPMappingsV2(context.Background(), svc); err != nil {
			t.Error(err)
		}
		closeStore()
	}
}

func TestBoltDBRPMappingServiceV2(t *testing.T) {
	t.Parallel()
	itesting.DBRPMappingServiceV2(initDBRPMappingService, t)
}
