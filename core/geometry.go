/*
 * Ported from three.js by @rydrman
 */

package core

import "github.com/rydrman/three.go/math3"

var count int = 0

func GeometryIdCount() int {
	count++
	return count
}

type Geometry struct {
	ID            int
	UUID          string
	Name          string
	Type          string
	Vertices      []*math3.Vector3
	Colors        []*math3.Color
	Faces         []*Face3
	FaceVertexUVs []*math3.Vector2

	MorphTargets []int
	MorphNormals []int

	SkinWeights []int
	SkinIndices []int

	LineDistances []float32

	BoundingBox *math3.Box3
}

func NewGeometry() *Geometry {

	return &Geometry{
		ID:   GeometryIdCount(),
		UUID: math3.GenerateUUID(),
		Name: "",
		Type: "Geometry",
	}

	/*this.vertices = []
	  this.colors = []
	  this.faces = []
	  this.faceVertexUvs = [[]]

	  this.morphTargets = []
	  this.morphNormals = []

	  this.skinWeights = []
	  this.skinIndices = []

	  this.lineDistances = []

	  this.boundingBox = nil
	  this.boundingSphere = nil

	  // update flags

	  this.elementsNeedUpdate = false
	  this.verticesNeedUpdate = false
	  this.uvsNeedUpdate = false
	  this.normalsNeedUpdate = false
	  this.colorsNeedUpdate = false
	  this.lineDistancesNeedUpdate = false
	  this.groupsNeedUpdate = false*/

}

/*Object.assign(Geometry.prototype, EventDispatcher.prototype, {

isGeometry: true,

applyMatrix: function (matrix) {

    var normalMatrix = new Matrix3().getNormalMatrix(matrix)

    for (var i = 0, il = this.vertices.length i < il i ++) {

        var vertex = this.vertices[i]
        vertex.applyMatrix4(matrix)

    }

    for (var i = 0, il = this.faces.length i < il i ++) {

        var face = this.faces[i]
        face.normal.applyMatrix3(normalMatrix).normalize()

        for (var j = 0, jl = face.vertexNormals.length j < jl j ++) {

            face.vertexNormals[j].applyMatrix3(normalMatrix).normalize()

        }

    }

    if (this.boundingBox != nil) {

        this.computeBoundingBox()

    }

    if (this.boundingSphere != nil) {

        this.computeBoundingSphere()

    }

    this.verticesNeedUpdate = true
    this.normalsNeedUpdate = true

    return this

},*/

/*rotateX: function () {

    // rotate geometry around world x-axis

    var m1

    return function rotateX(angle) {

        if (m1 === nil) m1 = new Matrix4()

        m1.makeRotationX(angle)

        this.applyMatrix(m1)

        return this

    }

}(),*/

/*rotateY: function () {

    // rotate geometry around world y-axis

    var m1

    return function rotateY(angle) {

        if (m1 === nil) m1 = new Matrix4()

        m1.makeRotationY(angle)

        this.applyMatrix(m1)

        return this

    }

}(),*/

/*rotateZ: function () {

    // rotate geometry around world z-axis

    var m1

    return function rotateZ(angle) {

        if (m1 === nil) m1 = new Matrix4()

        m1.makeRotationZ(angle)

        this.applyMatrix(m1)

        return this

    }

}(),
*/
/*translate: function () {

    // translate geometry

    var m1

    return function translate(x, y, z) {

        if (m1 === nil) m1 = new Matrix4()

        m1.makeTranslation(x, y, z)

        this.applyMatrix(m1)

        return this

    }

}(),*/

/*scale: function () {

    // scale geometry

    var m1

    return function scale(x, y, z) {

        if (m1 === nil) m1 = new Matrix4()

        m1.makeScale(x, y, z)

        this.applyMatrix(m1)

        return this

    }

}(),*/

/*lookAt: function () {

    var obj

    return function lookAt(vector) {

        if (obj === nil) obj = new Object3D()

        obj.lookAt(vector)

        obj.updateMatrix()

        this.applyMatrix(obj.matrix)

    }

}(),*/

func (geo *Geometry) FromBufferGeometry(srcGeo *BufferGeometry) *Geometry {

	var indices [][]float32
	if srcGeo.getIndex != nil {
		indices = srcGeo.Index.Array
	}

	var attributes = srcGeo.Attributes
	var positions = attributes["Position"].Array

	var (
		normals []*math3.Vector3
		colors  []*math3.Vector3
		uvs     []*math3.Vector2
		uvs2    []*math3.Vector2
	)

	if val, ok := attributes["normal"]; ok {
		normals = val.Array
	}
	if val, ok := attributes["color"]; ok {
		colors = val.Array
	}
	if val, ok := attributes["uv"]; ok {
		uvs = val.Array
	}
	if val, ok := attributes["uv2"]; ok {
		uvs2 = val.Array
	}

	if uvs2 != nil {

		geo.FaceVertexUvs[1] = make([]*math3.Vector2)

	}

	var (
		tempNormals []*math3.Vector2
		tempUVs     []*math3.Vector2
		tempUVs2    []*maths.Vector2
	)

	for i, j := 0, 0; i < len(positions); i, j = i+3, j+2 {

		geo.Vertices = append(geo.Vertices, math3.NewVector3(positions[i], positions[i+1], positions[i+2]))

		if normals != nil {

			tempNormals = append(tempNormals, NewVector3(normals[i], normals[i+1], normals[i+2]))

		}

		if colors != nil {

			scope.colors = append(colors, NewColor(colors[i], colors[i+1], colors[i+2]))

		}

		if uvs != nil {

			tempUVs = append(tempUVs, NewVector2(uvs[j], uvs[j+1]))

		}

		if uvs2 != nil {

			tempUVs2 = append(tempUVs2, NewVector2(uvs2[j], uvs2[j+1]))

		}

	}

	addFace := func(a, b, c int, materialIndex int) {

		var vertexNormals []*math3.Vector3
		if normals != nil {
			vertexNormals = append(vertexNormals, tempNormals[a].clone(), tempNormals[b].clone(), tempNormals[c].clone())
		}

		var vertexColors []*math3.Color
		if colors != nil {
			vertexColors = append(vertexColors, geo.Colors[a].clone(), geo.Colors[b].clone(), geo.Colors[c].clone())
		}

		face := NewFace3(a, b, c, vertexNormals, vertexColors, materialIndex)

		geo.Faces = append(geo.Faces, face)

		if uvs != nil {

			geo.faceVertexUvs[0] = append(geo.FaceVertexUVs[0], []*math3.Vector2{tempUVs[a].clone(), tempUVs[b].clone(), tempUVs[c].clone()})

		}

		if uvs2 != nil {

			geo.faceVertexUvs[1] = append(geo.FaceVertexUVs[1], []*math3.Vector2{tempUVs2[a].clone(), tempUVs2[b].clone(), tempUVs2[c].clone()})

		}

	}

	groups := srcGeo.groups

	if len(groups) > 0 {

		for i := 0; i < len(groups); i++ {

			group := groups[i]

			start := group.Start
			count := group.Count

			for j, jl := start, start+count; j < jl; j += 3 {

				if indices != nil {

					addFace(indices[j], indices[j+1], indices[j+2], group.materialIndex)

				} else {

					addFace(j, j+1, j+2, group.materialIndex)

				}

			}

		}

	} else {

		if indices != nil {

			for i := 0; i < len(indices); i += 3 {

				addFace(indices[i], indices[i+1], indices[i+2], -1)

			}

		} else {

			for i := 0; i < len(positions)/3; i += 3 {

				addFace(i, i+1, i+2, -1)

			}

		}

	}

	geo.ComputeFaceNormals()

	if srcGeo.BoundingBox != nil {

		geo.BoundingBox = srcGeo.BoundingBox.clone()

	}

	if srcGeo.BoundingSphere != nil {

		this.BoundingSphere = srcGeo.BoundingSphere.clone()

	}

	return geo

}

/*center: function () {

    this.computeBoundingBox()

    var offset = this.boundingBox.getCenter().negate()

    this.translate(offset.x, offset.y, offset.z)

    return offset

}/*

/*normalize: function () {

    this.computeBoundingSphere()

    var center = this.boundingSphere.center
    var radius = this.boundingSphere.radius

    var s = radius === 0 ? 1 : 1.0 / radius

    var matrix = new Matrix4()
    matrix.set(
        s, 0, 0, - s * center.x,
        0, s, 0, - s * center.y,
        0, 0, s, - s * center.z,
        0, 0, 0, 1
   )

    this.applyMatrix(matrix)

    return this

},*/

/*computeFaceNormals: function () {

    var cb = new Vector3(), ab = new Vector3()

    for (var f = 0, fl = this.faces.length f < fl f ++) {

        var face = this.faces[f]

        var vA = this.vertices[face.a]
        var vB = this.vertices[face.b]
        var vC = this.vertices[face.c]

        cb.subVectors(vC, vB)
        ab.subVectors(vA, vB)
        cb.cross(ab)

        cb.normalize()

        face.normal.copy(cb)

    }

},*/

/*computeVertexNormals: function (areaWeighted) {

    if (areaWeighted === nil) areaWeighted = true

    var v, vl, f, fl, face, vertices

    vertices = new Array(this.vertices.length)

    for (v = 0, vl = this.vertices.length v < vl v ++) {

        vertices[v] = new Vector3()

    }

    if (areaWeighted) {

        // vertex normals weighted by triangle areas
        // http://www.iquilezles.org/www/articles/normals/normals.htm

        var vA, vB, vC
        var cb = new Vector3(), ab = new Vector3()

        for (f = 0, fl = this.faces.length f < fl f ++) {

            face = this.faces[f]

            vA = this.vertices[face.a]
            vB = this.vertices[face.b]
            vC = this.vertices[face.c]

            cb.subVectors(vC, vB)
            ab.subVectors(vA, vB)
            cb.cross(ab)

            vertices[face.a].add(cb)
            vertices[face.b].add(cb)
            vertices[face.c].add(cb)

        }

    } else {

        this.computeFaceNormals()

        for (f = 0, fl = this.faces.length f < fl f ++) {

            face = this.faces[f]

            vertices[face.a].add(face.normal)
            vertices[face.b].add(face.normal)
            vertices[face.c].add(face.normal)

        }

    }

    for (v = 0, vl = this.vertices.length v < vl v ++) {

        vertices[v].normalize()

    }

    for (f = 0, fl = this.faces.length f < fl f ++) {

        face = this.faces[f]

        var vertexNormals = face.vertexNormals

        if (vertexNormals.length === 3) {

            vertexNormals[0].copy(vertices[face.a])
            vertexNormals[1].copy(vertices[face.b])
            vertexNormals[2].copy(vertices[face.c])

        } else {

            vertexNormals[0] = vertices[face.a].clone()
            vertexNormals[1] = vertices[face.b].clone()
            vertexNormals[2] = vertices[face.c].clone()

        }

    }

    if (this.faces.length > 0) {

        this.normalsNeedUpdate = true

    }

},

computeFlatVertexNormals: function () {

    var f, fl, face

    this.computeFaceNormals()

    for (f = 0, fl = this.faces.length f < fl f ++) {

        face = this.faces[f]

        var vertexNormals = face.vertexNormals

        if (vertexNormals.length === 3) {

            vertexNormals[0].copy(face.normal)
            vertexNormals[1].copy(face.normal)
            vertexNormals[2].copy(face.normal)

        } else {

            vertexNormals[0] = face.normal.clone()
            vertexNormals[1] = face.normal.clone()
            vertexNormals[2] = face.normal.clone()

        }

    }

    if (this.faces.length > 0) {

        this.normalsNeedUpdate = true

    }

},*/

/*computeMorphNormals: function () {

    var i, il, f, fl, face

    // save original normals
    // - create temp variables on first access
    //   otherwise just copy (for faster repeated calls)

    for (f = 0, fl = this.faces.length f < fl f ++) {

        face = this.faces[f]

        if (! face.__originalFaceNormal) {

            face.__originalFaceNormal = face.normal.clone()

        } else {

            face.__originalFaceNormal.copy(face.normal)

        }

        if (! face.__originalVertexNormals) face.__originalVertexNormals = []

        for (i = 0, il = face.vertexNormals.length i < il i ++) {

            if (! face.__originalVertexNormals[i]) {

                face.__originalVertexNormals[i] = face.vertexNormals[i].clone()

            } else {

                face.__originalVertexNormals[i].copy(face.vertexNormals[i])

            }

        }

    }

    // use temp geometry to compute face and vertex normals for each morph

    var tmpGeo = new Geometry()
    tmpGeo.faces = this.faces

    for (i = 0, il = this.morphTargets.length i < il i ++) {

        // create on first access

        if (! this.morphNormals[i]) {

            this.morphNormals[i] = {}
            this.morphNormals[i].faceNormals = []
            this.morphNormals[i].vertexNormals = []

            var dstNormalsFace = this.morphNormals[i].faceNormals
            var dstNormalsVertex = this.morphNormals[i].vertexNormals

            var faceNormal, vertexNormals

            for (f = 0, fl = this.faces.length f < fl f ++) {

                faceNormal = new Vector3()
                vertexNormals = { a: new Vector3(), b: new Vector3(), c: new Vector3() }

                dstNormalsFace.push(faceNormal)
                dstNormalsVertex.push(vertexNormals)

            }

        }

        var morphNormals = this.morphNormals[i]

        // set vertices to morph target

        tmpGeo.vertices = this.morphTargets[i].vertices

        // compute morph normals

        tmpGeo.computeFaceNormals()
        tmpGeo.computeVertexNormals()

        // store morph normals

        var faceNormal, vertexNormals

        for (f = 0, fl = this.faces.length f < fl f ++) {

            face = this.faces[f]

            faceNormal = morphNormals.faceNormals[f]
            vertexNormals = morphNormals.vertexNormals[f]

            faceNormal.copy(face.normal)

            vertexNormals.a.copy(face.vertexNormals[0])
            vertexNormals.b.copy(face.vertexNormals[1])
            vertexNormals.c.copy(face.vertexNormals[2])

        }

    }

    // restore original normals

    for (f = 0, fl = this.faces.length f < fl f ++) {

        face = this.faces[f]

        face.normal = face.__originalFaceNormal
        face.vertexNormals = face.__originalVertexNormals

    }

},

computeLineDistances: function () {

    var d = 0
    var vertices = this.vertices

    for (var i = 0, il = vertices.length i < il i ++) {

        if (i > 0) {

            d += vertices[i].distanceTo(vertices[i - 1])

        }

        this.lineDistances[i] = d

    }

},
*/
/*computeBoundingBox: function () {

    if (this.boundingBox === nil) {

        this.boundingBox = new Box3()

    }

    this.boundingBox.setFromPoints(this.vertices)

},

computeBoundingSphere: function () {

    if (this.boundingSphere === nil) {

        this.boundingSphere = new Sphere()

    }

    this.boundingSphere.setFromPoints(this.vertices)

},*/

/*merge: function (geometry, matrix, materialIndexOffset) {

    if ((geometry && geometry.isGeometry) === false) {

        console.error('THREE.Geometry.merge(): geometry not an instance of THREE.Geometry.', geometry)
        return

    }

    var normalMatrix,
        vertexOffset = this.vertices.length,
        vertices1 = this.vertices,
        vertices2 = geometry.vertices,
        faces1 = this.faces,
        faces2 = geometry.faces,
        uvs1 = this.faceVertexUvs[0],
        uvs2 = geometry.faceVertexUvs[0],
        colors1 = this.colors,
        colors2 = geometry.colors

    if (materialIndexOffset === nil) materialIndexOffset = 0

    if (matrix != nil) {

        normalMatrix = new Matrix3().getNormalMatrix(matrix)

    }

    // vertices

    for (var i = 0, il = vertices2.length i < il i ++) {

        var vertex = vertices2[i]

        var vertexCopy = vertex.clone()

        if (matrix != nil) vertexCopy.applyMatrix4(matrix)

        vertices1.push(vertexCopy)

    }

    // colors

    for (var i = 0, il = colors2.length i < il i ++) {

        colors1.push(colors2[i].clone())

    }

    // faces

    for (i = 0, il = faces2.length i < il i ++) {

        var face = faces2[i], faceCopy, normal, color,
            faceVertexNormals = face.vertexNormals,
            faceVertexColors = face.vertexColors

        faceCopy = new Face3(face.a + vertexOffset, face.b + vertexOffset, face.c + vertexOffset)
        faceCopy.normal.copy(face.normal)

        if (normalMatrix != nil) {

            faceCopy.normal.applyMatrix3(normalMatrix).normalize()

        }

        for (var j = 0, jl = faceVertexNormals.length j < jl j ++) {

            normal = faceVertexNormals[j].clone()

            if (normalMatrix != nil) {

                normal.applyMatrix3(normalMatrix).normalize()

            }

            faceCopy.vertexNormals.push(normal)

        }

        faceCopy.color.copy(face.color)

        for (var j = 0, jl = faceVertexColors.length j < jl j ++) {

            color = faceVertexColors[j]
            faceCopy.vertexColors.push(color.clone())

        }

        faceCopy.materialIndex = face.materialIndex + materialIndexOffset

        faces1.push(faceCopy)

    }

    // uvs

    for (i = 0, il = uvs2.length i < il i ++) {

        var uv = uvs2[i], uvCopy = []

        if (uv === nil) {

            continue

        }

        for (var j = 0, jl = uv.length j < jl j ++) {

            uvCopy.push(uv[j].clone())

        }

        uvs1.push(uvCopy)

    }

},*/

/*mergeMesh: function (mesh) {

    if ((mesh && mesh.isMesh) === false) {

        console.error('THREE.Geometry.mergeMesh(): mesh not an instance of THREE.Mesh.', mesh)
        return

    }

    mesh.matrixAutoUpdate && mesh.updateMatrix()

    this.merge(mesh.geometry, mesh.matrix)

},*/

/*
 * Checks for duplicate vertices with hashmap.
 * Duplicated vertices are removed
 * and faces' vertices are updated.
 */

/*mergeVertices: function () {

    var verticesMap = {} // Hashmap for looking up vertices by position coordinates (and making sure they are unique)
    var unique = [], changes = []

    var v, key
    var precisionPoints = 4 // number of decimal points, e.g. 4 for epsilon of 0.0001
    var precision = Math.pow(10, precisionPoints)
    var i, il, face
    var indices, j, jl

    for (i = 0, il = this.vertices.length i < il i ++) {

        v = this.vertices[i]
        key = Math.round(v.x * precision) + '_' + Math.round(v.y * precision) + '_' + Math.round(v.z * precision)

        if (verticesMap[key] === nil) {

            verticesMap[key] = i
            unique.push(this.vertices[i])
            changes[i] = unique.length - 1

        } else {

            //console.log('Duplicate vertex found. ', i, ' could be using ', verticesMap[key])
            changes[i] = changes[verticesMap[key]]

        }

    }


    // if faces are completely degenerate after merging vertices, we
    // have to remove them from the geometry.
    var faceIndicesToRemove = []

    for (i = 0, il = this.faces.length i < il i ++) {

        face = this.faces[i]

        face.a = changes[face.a]
        face.b = changes[face.b]
        face.c = changes[face.c]

        indices = [face.a, face.b, face.c]

        // if any duplicate vertices are found in a Face3
        // we have to remove the face as nothing can be saved
        for (var n = 0 n < 3 n ++) {

            if (indices[n] === indices[(n + 1) % 3]) {

                faceIndicesToRemove.push(i)
                break

            }

        }

    }

    for (i = faceIndicesToRemove.length - 1 i >= 0 i --) {

        var idx = faceIndicesToRemove[i]

        this.faces.splice(idx, 1)

        for (j = 0, jl = this.faceVertexUvs.length j < jl j ++) {

            this.faceVertexUvs[j].splice(idx, 1)

        }

    }

    // Use unique set of vertices

    var diff = this.vertices.length - unique.length
    this.vertices = unique
    return diff

},*/

/*sortFacesByMaterialIndex: function () {

    var faces = this.faces
    var length = faces.length

    // tag faces

    for (var i = 0 i < length i ++) {

        faces[i]._id = i

    }

    // sort faces

    function materialIndexSort(a, b) {

        return a.materialIndex - b.materialIndex

    }

    faces.sort(materialIndexSort)

    // sort uvs

    var uvs1 = this.faceVertexUvs[0]
    var uvs2 = this.faceVertexUvs[1]

    var newUvs1, newUvs2

    if (uvs1 && uvs1.length === length) newUvs1 = []
    if (uvs2 && uvs2.length === length) newUvs2 = []

    for (var i = 0 i < length i ++) {

        var id = faces[i]._id

        if (newUvs1) newUvs1.push(uvs1[id])
        if (newUvs2) newUvs2.push(uvs2[id])

    }

    if (newUvs1) this.faceVertexUvs[0] = newUvs1
    if (newUvs2) this.faceVertexUvs[1] = newUvs2

},

toJSON: function () {

    var data = {
        metadata: {
            version: 4.5,
            type: 'Geometry',
            generator: 'Geometry.toJSON'
        }
    }

    // standard Geometry serialization

    data.uuid = this.uuid
    data.type = this.type
    if (this.name != '') data.name = this.name

    if (this.parameters != nil) {

        var parameters = this.parameters

        for (var key in parameters) {

            if (parameters[key] != nil) data[key] = parameters[key]

        }

        return data

    }

    var vertices = []

    for (var i = 0 i < this.vertices.length i ++) {

        var vertex = this.vertices[i]
        vertices.push(vertex.x, vertex.y, vertex.z)

    }

    var faces = []
    var normals = []
    var normalsHash = {}
    var colors = []
    var colorsHash = {}
    var uvs = []
    var uvsHash = {}

    for (var i = 0 i < this.faces.length i ++) {

        var face = this.faces[i]

        var hasMaterial = true
        var hasFaceUv = false // deprecated
        var hasFaceVertexUv = this.faceVertexUvs[0][i] != nil
        var hasFaceNormal = face.normal.length() > 0
        var hasFaceVertexNormal = face.vertexNormals.length > 0
        var hasFaceColor = face.color.r != 1 || face.color.g != 1 || face.color.b != 1
        var hasFaceVertexColor = face.vertexColors.length > 0

        var faceType = 0

        faceType = setBit(faceType, 0, 0) // isQuad
        faceType = setBit(faceType, 1, hasMaterial)
        faceType = setBit(faceType, 2, hasFaceUv)
        faceType = setBit(faceType, 3, hasFaceVertexUv)
        faceType = setBit(faceType, 4, hasFaceNormal)
        faceType = setBit(faceType, 5, hasFaceVertexNormal)
        faceType = setBit(faceType, 6, hasFaceColor)
        faceType = setBit(faceType, 7, hasFaceVertexColor)

        faces.push(faceType)
        faces.push(face.a, face.b, face.c)
        faces.push(face.materialIndex)

        if (hasFaceVertexUv) {

            var faceVertexUvs = this.faceVertexUvs[0][i]

            faces.push(
                getUvIndex(faceVertexUvs[0]),
                getUvIndex(faceVertexUvs[1]),
                getUvIndex(faceVertexUvs[2])
           )

        }

        if (hasFaceNormal) {

            faces.push(getNormalIndex(face.normal))

        }

        if (hasFaceVertexNormal) {

            var vertexNormals = face.vertexNormals

            faces.push(
                getNormalIndex(vertexNormals[0]),
                getNormalIndex(vertexNormals[1]),
                getNormalIndex(vertexNormals[2])
           )

        }

        if (hasFaceColor) {

            faces.push(getColorIndex(face.color))

        }

        if (hasFaceVertexColor) {

            var vertexColors = face.vertexColors

            faces.push(
                getColorIndex(vertexColors[0]),
                getColorIndex(vertexColors[1]),
                getColorIndex(vertexColors[2])
           )

        }

    }

    function setBit(value, position, enabled) {

        return enabled ? value | (1 << position) : value & (~ (1 << position))

    }

    function getNormalIndex(normal) {

        var hash = normal.x.toString() + normal.y.toString() + normal.z.toString()

        if (normalsHash[hash] != nil) {

            return normalsHash[hash]

        }

        normalsHash[hash] = normals.length / 3
        normals.push(normal.x, normal.y, normal.z)

        return normalsHash[hash]

    }

    function getColorIndex(color) {

        var hash = color.r.toString() + color.g.toString() + color.b.toString()

        if (colorsHash[hash] != nil) {

            return colorsHash[hash]

        }

        colorsHash[hash] = colors.length
        colors.push(color.getHex())

        return colorsHash[hash]

    }

    function getUvIndex(uv) {

        var hash = uv.x.toString() + uv.y.toString()

        if (uvsHash[hash] != nil) {

            return uvsHash[hash]

        }

        uvsHash[hash] = uvs.length / 2
        uvs.push(uv.x, uv.y)

        return uvsHash[hash]

    }

    data.data = {}

    data.data.vertices = vertices
    data.data.normals = normals
    if (colors.length > 0) data.data.colors = colors
    if (uvs.length > 0) data.data.uvs = [uvs] // temporal backward compatibility
    data.data.faces = faces

    return data

},
*/
//clone: function () {

/*
 // Handle primitives

 var parameters = this.parameters

 if (parameters != nil) {

 var values = []

 for (var key in parameters) {

 values.push(parameters[key])

 }

 var geometry = Object.create(this.constructor.prototype)
 this.constructor.apply(geometry, values)
 return geometry

 }

 return new this.constructor().copy(this)


return new Geometry().copy(this)*/

//},

/*copy: function (source) {

    var i, il, j, jl, k, kl

    // reset

    this.vertices = []
    this.colors = []
    this.faces = []
    this.faceVertexUvs = [[]]
    this.morphTargets = []
    this.morphNormals = []
    this.skinWeights = []
    this.skinIndices = []
    this.lineDistances = []
    this.boundingBox = nil
    this.boundingSphere = nil

    // name

    this.name = source.name

    // vertices

    var vertices = source.vertices

    for (i = 0, il = vertices.length i < il i ++) {

        this.vertices.push(vertices[i].clone())

    }

    // colors

    var colors = source.colors

    for (i = 0, il = colors.length i < il i ++) {

        this.colors.push(colors[i].clone())

    }

    // faces

    var faces = source.faces

    for (i = 0, il = faces.length i < il i ++) {

        this.faces.push(faces[i].clone())

    }

    // face vertex uvs

    for (i = 0, il = source.faceVertexUvs.length i < il i ++) {

        var faceVertexUvs = source.faceVertexUvs[i]

        if (this.faceVertexUvs[i] === nil) {

            this.faceVertexUvs[i] = []

        }

        for (j = 0, jl = faceVertexUvs.length j < jl j ++) {

            var uvs = faceVertexUvs[j], uvsCopy = []

            for (k = 0, kl = uvs.length k < kl k ++) {

                var uv = uvs[k]

                uvsCopy.push(uv.clone())

            }

            this.faceVertexUvs[i].push(uvsCopy)

        }

    }

    // morph targets

    var morphTargets = source.morphTargets

    for (i = 0, il = morphTargets.length i < il i ++) {

        var morphTarget = {}
        morphTarget.name = morphTargets[i].name

        // vertices

        if (morphTargets[i].vertices != nil) {

            morphTarget.vertices = []

            for (j = 0, jl = morphTargets[i].vertices.length j < jl j ++) {

                morphTarget.vertices.push(morphTargets[i].vertices[j].clone())

            }

        }

        // normals

        if (morphTargets[i].normals != nil) {

            morphTarget.normals = []

            for (j = 0, jl = morphTargets[i].normals.length j < jl j ++) {

                morphTarget.normals.push(morphTargets[i].normals[j].clone())

            }

        }

        this.morphTargets.push(morphTarget)

    }

    // morph normals

    var morphNormals = source.morphNormals

    for (i = 0, il = morphNormals.length i < il i ++) {

        var morphNormal = {}

        // vertex normals

        if (morphNormals[i].vertexNormals != nil) {

            morphNormal.vertexNormals = []

            for (j = 0, jl = morphNormals[i].vertexNormals.length j < jl j ++) {

                var srcVertexNormal = morphNormals[i].vertexNormals[j]
                var destVertexNormal = {}

                destVertexNormal.a = srcVertexNormal.a.clone()
                destVertexNormal.b = srcVertexNormal.b.clone()
                destVertexNormal.c = srcVertexNormal.c.clone()

                morphNormal.vertexNormals.push(destVertexNormal)

            }

        }

        // face normals

        if (morphNormals[i].faceNormals != nil) {

            morphNormal.faceNormals = []

            for (j = 0, jl = morphNormals[i].faceNormals.length j < jl j ++) {

                morphNormal.faceNormals.push(morphNormals[i].faceNormals[j].clone())

            }

        }

        this.morphNormals.push(morphNormal)

    }

    // skin weights

    var skinWeights = source.skinWeights

    for (i = 0, il = skinWeights.length i < il i ++) {

        this.skinWeights.push(skinWeights[i].clone())

    }

    // skin indices

    var skinIndices = source.skinIndices

    for (i = 0, il = skinIndices.length i < il i ++) {

        this.skinIndices.push(skinIndices[i].clone())

    }

    // line distances

    var lineDistances = source.lineDistances

    for (i = 0, il = lineDistances.length i < il i ++) {

        this.lineDistances.push(lineDistances[i])

    }

    // bounding box

    var boundingBox = source.boundingBox

    if (boundingBox != nil) {

        this.boundingBox = boundingBox.clone()

    }

    // bounding sphere

    var boundingSphere = source.boundingSphere

    if (boundingSphere != nil) {

        this.boundingSphere = boundingSphere.clone()

    }

    // update flags

    this.elementsNeedUpdate = source.elementsNeedUpdate
    this.verticesNeedUpdate = source.verticesNeedUpdate
    this.uvsNeedUpdate = source.uvsNeedUpdate
    this.normalsNeedUpdate = source.normalsNeedUpdate
    this.colorsNeedUpdate = source.colorsNeedUpdate
    this.lineDistancesNeedUpdate = source.lineDistancesNeedUpdate
    this.groupsNeedUpdate = source.groupsNeedUpdate

    return this

},
*/
/*dispose: function () {

        this.dispatchEvent({ type: 'dispose' })

    }

})


export { GeometryIdCount, Geometry }
*/
