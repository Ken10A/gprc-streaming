import sys

import grpc

from proto import archiver_pb2
from proto import archiver_pb2_grpc as rpc


def _read(path):
    with open(path, 'rb') as stream:
        return archiver_pb2.ZipRequest(file_name=path, contents=stream.read())


def archive(file_paths):
    channel = grpc.insecure_channel('localhost:8080')
    stub = rpc.ArchiverStub(channel)

    files = iter([_read(path) for path in file_paths])
    future = stub.Zip.future(files)

    # we'll just let errors go to the console
    response = future.result()

    with open('compressed.zip', 'wb') as stream:
        stream.write(response.zipped_contents)


if __name__ == '__main__':
    archive(sys.argv[1:])
