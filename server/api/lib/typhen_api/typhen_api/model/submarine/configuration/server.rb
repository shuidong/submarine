# This file was generated by typhen-api

module TyphenApi::Model::Submarine::Configuration
  class Server
    include Virtus.model(:strict => true)

    attribute :api_server_base_uri, String, :required => true
    attribute :battle_server_base_uri, String, :required => true
    attribute :database, TyphenApi::Model::Submarine::Configuration::ServerDatabaseObject, :required => true
  end
end
