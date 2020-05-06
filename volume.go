/*
 *
 * Copyright © 2020 Dell Inc. or its subsidiaries. All Rights Reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package gopowerstore

import (
	"context"
	"github.com/dell/gopowerstore/api"
	"fmt"
)

const (
	volumeURL = "volume"
)

func getVolumeDefaultQueryParams(c Client) api.QueryParamsEncoder {
	vol := Volume{}
	return c.APIClient().QueryParamsWithFields(&vol)
}

// GetVolume query and return specific volume by id
func (c *ClientIMPL) GetVolume(ctx context.Context, id string) (resp Volume, err error) {
	_, err = c.APIClient().Query(
		ctx,
		RequestConfig{
			Method:      "GET",
			Endpoint:    volumeURL,
			ID:          id,
			QueryParams: getVolumeDefaultQueryParams(c)},
		&resp)
	return resp, WrapErr(err)
}

// GetVolumeByName query and return specific volume by name
func (c *ClientIMPL) GetVolumeByName(ctx context.Context, name string) (resp Volume, err error) {
	var volList []Volume
	qp := getVolumeDefaultQueryParams(c)
	qp.RawArg("name", fmt.Sprintf("eq.%s", name))
	_, err = c.APIClient().Query(
		ctx,
		RequestConfig{
			Method:      "GET",
			Endpoint:    volumeURL,
			QueryParams: qp},
		&volList)
	err = WrapErr(err)
	if err != nil {
		return resp, err
	}
	if len(volList) != 1 {
		return resp, NewVolumeIsNotExistError()
	}
	return volList[0], err
}

// GetVolumes returns a list of volumes
func (c *ClientIMPL) GetVolumes(ctx context.Context) ([]Volume, error) {
	var result []Volume
	err := c.readPaginatedData(func(offset int) (api.RespMeta, error) {
		var page []Volume
		qp := getVolumeDefaultQueryParams(c)
		qp.RawArg("type", fmt.Sprintf("not.eq.%s", VolumeTypeEnumSnapshot))
		qp.Order("name")
		qp.Offset(offset).Limit(paginationDefaultPageSize)
		meta, err := c.APIClient().Query(
			ctx,
			RequestConfig{
				Method:      "GET",
				Endpoint:    volumeURL,
				QueryParams: qp},
			&page)
		err = WrapErr(err)
		if err == nil {
			result = append(result, page...)
		}
		return meta, err
	})
	return result, err
}

// GetSnapshot query and return specific snapshot by it's id
func (c *ClientIMPL) GetSnapshot(ctx context.Context, snapID string) (resVol Volume, err error) {
	qp := getVolumeDefaultQueryParams(c)
	qp.RawArg("type", fmt.Sprintf("eq.%s", VolumeTypeEnumSnapshot))
	_, err = c.APIClient().Query(
		ctx,
		RequestConfig{
			Method:      "GET",
			Endpoint:    volumeURL,
			ID:          snapID,
			QueryParams: qp},
		&resVol)
	return resVol, WrapErr(err)
}

// GetSnapshots returns all snapshots
func (c *ClientIMPL) GetSnapshots(ctx context.Context) ([]Volume, error) {
	var result []Volume
	err := c.readPaginatedData(func(offset int) (api.RespMeta, error) {
		var page []Volume
		qp := getVolumeDefaultQueryParams(c)
		qp.RawArg("type", fmt.Sprintf("eq.%s", VolumeTypeEnumSnapshot))
		qp.Order("name")
		qp.Offset(offset).Limit(paginationDefaultPageSize)
		meta, err := c.APIClient().Query(
			ctx,
			RequestConfig{
				Method:      "GET",
				Endpoint:    volumeURL,
				QueryParams: qp},
			&page)
		err = WrapErr(err)
		if err == nil {
			result = append(result, page...)
		}
		return meta, err
	})
	return result, err
}

// GetSnapshotsByVolumeID returns a list of snapshots for specific volume
func (c *ClientIMPL) GetSnapshotsByVolumeID(ctx context.Context, volID string) ([]Volume, error) {
	var result []Volume
	err := c.readPaginatedData(func(offset int) (api.RespMeta, error) {
		var page []Volume
		qp := getVolumeDefaultQueryParams(c)
		qp.RawArg("protection_data->>source_id", fmt.Sprintf("eq.%s", volID))
		qp.RawArg("type", fmt.Sprintf("eq.%s", VolumeTypeEnumSnapshot))
		qp.Order("name")
		qp.Offset(offset).Limit(paginationDefaultPageSize)
		meta, err := c.APIClient().Query(
			ctx,
			RequestConfig{
				Method:      "GET",
				Endpoint:    volumeURL,
				QueryParams: qp},
			&page)
		err = WrapErr(err)
		if err == nil {
			result = append(result, page...)
		}
		return meta, err
	})
	return result, err
}

// CreateVolume creates new volume
func (c *ClientIMPL) CreateVolume(ctx context.Context,
	createParams *VolumeCreate) (resp CreateResponse, err error) {
	_, err = c.APIClient().Query(
		ctx,
		RequestConfig{
			Method:   "POST",
			Endpoint: volumeURL,
			Body:     createParams},
		&resp)
	return resp, WrapErr(err)
}

// CreateVolumeFromSnapshot creates a new volume by cloning a snapshot
func (c *ClientIMPL) CreateVolumeFromSnapshot(ctx context.Context,
	createParams *VolumeClone, snapID string) (resp CreateResponse, err error) {
	_, err = c.APIClient().Query(
		ctx,
		RequestConfig{
			Method:   "POST",
			Endpoint: volumeURL,
			ID:       snapID,
			Action:   "clone",
			Body:     createParams,
		},
		&resp)
	return resp, WrapErr(err)
}

// CreateSnapshot creates a new snapshot
func (c *ClientIMPL) CreateSnapshot(ctx context.Context,
	createSnapParams *SnapshotCreate, id string) (resp CreateResponse, err error) {
	_, err = c.APIClient().Query(
		ctx,
		RequestConfig{
			Method:   "POST",
			Endpoint: volumeURL,
			ID:       id,
			Action:   "snapshot",
			Body:     createSnapParams},
		&resp)
	return resp, WrapErr(err)
}

// DeleteVolume deletes existing volume
func (c *ClientIMPL) DeleteVolume(ctx context.Context,
	deleteParams *VolumeDelete, id string) (resp EmptyResponse, err error) {
	_, err = c.APIClient().Query(
		ctx,
		RequestConfig{
			Method:   "DELETE",
			Endpoint: volumeURL,
			ID:       id,
			Body:     deleteParams},
		&resp)
	return resp, WrapErr(err)
}

// DeleteSnapshot is an alias for delete volume, because snapshots are essentially -- volumes
func (c *ClientIMPL) DeleteSnapshot(ctx context.Context,
	deleteParams *VolumeDelete, id string) (resp EmptyResponse, err error) {
	return c.DeleteVolume(ctx, deleteParams, id)
}